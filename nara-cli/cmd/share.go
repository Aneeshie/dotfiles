package cmd

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

// getLocalIP finds the active Wi-Fi / LAN IPv4 address reaching the network
func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "127.0.0.1", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

func getFreePort() int {
	for port := 8080; port < 8095; port++ {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			_ = ln.Close()
			return port
		}
	}
	return 8080
}

func formatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func RunShare(args []string) error {
	target := "."
	if len(args) > 0 {
		target = args[0]
	}

	absPath, err := filepath.Abs(target)
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}

	info, err := os.Stat(absPath)
	if err != nil {
		return fmt.Errorf("file/folder not found: %w", err)
	}

	localIP, err := getLocalIP()
	if err != nil {
		return fmt.Errorf("failed to detect local IP: %w", err)
	}

	port := getFreePort()
	var shareURL string

	if info.IsDir() {
		shareURL = fmt.Sprintf("http://%s:%d/", localIP, port)
	} else {
		shareURL = fmt.Sprintf("http://%s:%d/%s", localIP, port, filepath.Base(absPath))
	}

	fmt.Println("\033[1;35m📱 [nara share]\033[0m Starting chunked Wi-Fi file server...")
	fmt.Printf("\033[1mSharing:\033[0m %s (%s)\n", absPath, formatBytes(info.Size()))
	fmt.Printf("\033[1mURL:\033[0m     \033[1;36m%s\033[0m\n\n", shareURL)

	fmt.Println("\033[1mScan with your Phone / Tablet camera:\033[0m")
	qrCmd := exec.Command("nix", "run", "nixpkgs#qrencode", "--", "-t", "UTF8i", shareURL)
	qrCmd.Stdout = os.Stdout
	_ = qrCmd.Run()

	fmt.Println()
	fmt.Println("\033[2mPress Ctrl+C to stop sharing\033[0m")
	fmt.Println("──────────────────────────────────────────")

	if info.IsDir() {
		http.Handle("/", http.FileServer(http.Dir(absPath)))
	} else {
		dir := filepath.Dir(absPath)
		fileName := filepath.Base(absPath)
		http.HandleFunc("/"+fileName, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Accept-Ranges", "bytes")
			w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
			fmt.Printf("\033[32m[+] Device connected!\033[0m Streaming chunked file %s (%s)...\n", fileName, formatBytes(info.Size()))
			http.ServeFile(w, r, absPath)
		})
		http.Handle("/", http.FileServer(http.Dir(dir)))
	}

	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
