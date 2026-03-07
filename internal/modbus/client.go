package modbus

import (
	"fmt"
	"net"
	"sync"
	"time"

	mb "github.com/goburrow/modbus"
)

// Client wraps a Modbus TCP connection with thread-safe read access.
type Client struct {
	handler *mb.TCPClientHandler
	client  mb.Client
	mu      sync.Mutex
}

// NewTCPClient creates a Modbus TCP client connected to host:port with the given slave ID.
func NewTCPClient(host string, port int, slaveID byte) (*Client, error) {
	return NewTCPClientWithTimeout(host, port, slaveID, 5*time.Second)
}

// NewTCPClientWithTimeout creates a Modbus TCP client with a custom timeout.
func NewTCPClientWithTimeout(host string, port int, slaveID byte, timeout time.Duration) (*Client, error) {
	handler := mb.NewTCPClientHandler(net.JoinHostPort(host, fmt.Sprintf("%d", port)))
	handler.Timeout = timeout
	handler.SlaveId = slaveID

	if err := handler.Connect(); err != nil {
		return nil, fmt.Errorf("modbus tcp connect %s:%d: %w", host, port, err)
	}

	return &Client{
		handler: handler,
		client:  mb.NewClient(handler),
	}, nil
}

// ReadInputRegisters reads count input registers starting at addr.
func (c *Client) ReadInputRegisters(addr, count uint16) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.client.ReadInputRegisters(addr, count)
}

// ReadHoldingRegisters reads count holding registers starting at addr.
func (c *Client) ReadHoldingRegisters(addr, count uint16) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.client.ReadHoldingRegisters(addr, count)
}

// ReadRegisters reads registers using the appropriate function code.
func (c *Client) ReadRegisters(addr, count uint16, holding bool) ([]byte, error) {
	if holding {
		return c.ReadHoldingRegisters(addr, count)
	}
	return c.ReadInputRegisters(addr, count)
}

// Close closes the underlying Modbus connection.
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.handler.Close()
}

// SetSlaveID changes the slave ID for subsequent requests.
func (c *Client) SetSlaveID(id byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.handler.SlaveId = id
}
