package scanner

import (
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

// Interface represents a network interface with its subnets.
type Interface struct {
	Name    string   `json:"name"`
	Subnets []string `json:"subnets"` // CIDR notation, e.g. "192.168.1.0/24"
}

// Host represents a discovered host with an open port.
type Host struct {
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	Duration string `json:"duration"` // connection time
}

// DefaultPorts are common Modbus TCP ports to scan.
var DefaultPorts = []PortInfo{
	{Port: 502, Description: "Modbus TCP (standard)"},
	{Port: 1502, Description: "Kostal/Alt Modbus TCP"},
	{Port: 8899, Description: "Deye/Solis/Sofar Wi-Fi dongle"},
	{Port: 6607, Description: "Huawei SmartDongle"},
}

// PortInfo describes a port with its purpose.
type PortInfo struct {
	Port        int    `json:"port"`
	Description string `json:"description"`
}

// UniqueDefaultPorts returns deduplicated default ports.
func UniqueDefaultPorts() []PortInfo {
	seen := make(map[int]bool)
	var result []PortInfo
	for _, p := range DefaultPorts {
		if !seen[p.Port] {
			seen[p.Port] = true
			result = append(result, p)
		}
	}
	return result
}

// ScanProgress reports scan status via a callback.
type ScanProgress struct {
	Port       int    `json:"port"`
	Status     string `json:"status"` // "scanning", "done"
	HostsFound int    `json:"hosts_found"`
	Total      int    `json:"total_ips"`
	Scanned    int    `json:"scanned"`
}

// GetInterfaces returns all non-loopback network interfaces with IPv4 subnets.
func GetInterfaces() ([]Interface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var result []Interface
	for _, iface := range ifaces {
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		var subnets []string
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok || ipnet.IP.To4() == nil {
				continue
			}
			// Calculate network address
			network := ipnet.IP.Mask(ipnet.Mask)
			ones, _ := ipnet.Mask.Size()
			subnets = append(subnets, fmt.Sprintf("%s/%d", network.String(), ones))
		}
		if len(subnets) > 0 {
			result = append(result, Interface{Name: iface.Name, Subnets: subnets})
		}
	}
	return result, nil
}

// ScanSubnet scans a CIDR subnet for hosts with the given port open.
// Uses concurrent TCP dial with the specified timeout.
func ScanSubnet(cidr string, port int, timeout time.Duration) ([]Host, error) {
	hosts, err := enumerateHosts(cidr)
	if err != nil {
		return nil, err
	}
	return scanHosts(hosts, port, timeout), nil
}

// ScanSubnetMultiPort scans a subnet for multiple ports. Calls onProgress after each port completes.
func ScanSubnetMultiPort(cidr string, ports []int, timeout time.Duration, onProgress func(ScanProgress)) ([]Host, error) {
	hosts, err := enumerateHosts(cidr)
	if err != nil {
		return nil, err
	}

	var allResults []Host
	for _, port := range ports {
		if onProgress != nil {
			onProgress(ScanProgress{Port: port, Status: "scanning", Total: len(hosts)})
		}

		results := scanHosts(hosts, port, timeout)
		allResults = append(allResults, results...)

		if onProgress != nil {
			onProgress(ScanProgress{Port: port, Status: "done", HostsFound: len(results), Total: len(hosts), Scanned: len(hosts)})
		}
	}

	// Deduplicate by IP:Port
	seen := make(map[string]bool)
	var deduped []Host
	for _, h := range allResults {
		key := fmt.Sprintf("%s:%d", h.IP, h.Port)
		if !seen[key] {
			seen[key] = true
			deduped = append(deduped, h)
		}
	}

	sort.Slice(deduped, func(i, j int) bool {
		if deduped[i].IP == deduped[j].IP {
			return deduped[i].Port < deduped[j].Port
		}
		return deduped[i].IP < deduped[j].IP
	})

	return deduped, nil
}

func enumerateHosts(cidr string) ([]net.IP, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR: %w", err)
	}

	var hosts []net.IP
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incrementIP(ip) {
		hostIP := make(net.IP, len(ip))
		copy(hostIP, ip)
		hosts = append(hosts, hostIP)
	}

	// Skip network and broadcast addresses for /24 and larger
	if len(hosts) > 2 {
		hosts = hosts[1 : len(hosts)-1]
	}
	return hosts, nil
}

func scanHosts(hosts []net.IP, port int, timeout time.Duration) []Host {
	maxConcurrent := 256
	sem := make(chan struct{}, maxConcurrent)

	var mu sync.Mutex
	var results []Host
	var wg sync.WaitGroup

	for _, hostIP := range hosts {
		wg.Add(1)
		sem <- struct{}{}
		go func(ip string) {
			defer wg.Done()
			defer func() { <-sem }()

			addr := fmt.Sprintf("%s:%d", ip, port)
			start := time.Now()
			conn, err := net.DialTimeout("tcp", addr, timeout)
			duration := time.Since(start)
			if err != nil {
				return
			}
			conn.Close()

			mu.Lock()
			results = append(results, Host{
				IP:       ip,
				Port:     port,
				Duration: fmt.Sprintf("%dms", duration.Milliseconds()),
			})
			mu.Unlock()
		}(hostIP.String())
	}

	wg.Wait()
	return results
}

func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
