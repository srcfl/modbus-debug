package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"
)

//go:embed web
var webFS embed.FS

var version = "dev"

func main() {
	port := flag.Int("port", 8765, "HTTP server port")
	noBrowser := flag.Bool("no-browser", false, "Don't auto-open browser")
	flag.Parse()

	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("GET /api/profiles", handleProfiles)
	mux.HandleFunc("GET /api/network/interfaces", handleNetworkInterfaces)
	mux.HandleFunc("POST /api/network/scan", handleNetworkScan)
	mux.HandleFunc("GET /api/network/ports", handleScanPorts)
	mux.HandleFunc("POST /api/diagnose/tcp", handleDiagnoseTCP)
	mux.HandleFunc("POST /api/diagnose/detect", handleDiagnoseDetect)
	mux.HandleFunc("POST /api/diagnose/detect-one", handleDiagnoseDetectOne)
	mux.HandleFunc("POST /api/diagnose/detect-batch", handleDiagnoseDetectBatch)
	mux.HandleFunc("POST /api/diagnose/read", handleDiagnoseRead)
	mux.HandleFunc("GET /api/diagnose/report", handleDiagnoseReport)
	mux.HandleFunc("POST /api/diagnose/report", handleUpdateReport)
	mux.HandleFunc("POST /api/raw/read", handleRawRead)

	// Serve embedded web UI
	webContent, err := fs.Sub(webFS, "web")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/", noCacheHandler(http.FileServer(http.FS(webContent))))

	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}

	url := fmt.Sprintf("http://%s", addr)
	fmt.Printf("Modbus Debug Tool %s\n", version)
	fmt.Printf("Listening on %s\n", url)

	if !*noBrowser {
		openBrowser(url)
	}

	if err := http.Serve(listener, mux); err != nil {
		log.Fatal(err)
	}
}

func noCacheHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
		h.ServeHTTP(w, r)
	})
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		return
	}
	cmd.Start()
}
