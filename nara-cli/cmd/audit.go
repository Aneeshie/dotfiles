package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunAudit() error {
	fmt.Println("\033[1;35m⚡ [nara]\033[0m Running macOS Performance & Battery Audit...")
	fmt.Println()

	fmt.Println("\033[1;36m1. Battery & Power Status:\033[0m")
	battOut, err := exec.Command("pmset", "-g", "batt").Output()
	if err == nil {
		lines := strings.Split(strings.TrimSpace(string(battOut)), "\n")
		for _, line := range lines {
			fmt.Printf("   %s\n", line)
		}
	} else {
		fmt.Println("   Unable to fetch battery status.")
	}
	fmt.Println()

	fmt.Println("\033[1;36m2. Top Memory Consuming Processes:\033[0m")
	memOut, err := exec.Command("sh", "-c", "ps -eo pmem,pcpu,comm -m | head -n 6").Output()
	if err == nil {
		lines := strings.Split(strings.TrimSpace(string(memOut)), "\n")
		for i, line := range lines {
			if i == 0 {
				fmt.Printf("   \033[1m%-8s %-8s %s\033[0m\n", "MEM %", "CPU %", "PROCESS")
			} else {
				fmt.Printf("   %s\n", line)
			}
		}
	}
	fmt.Println()

	fmt.Println("\033[1;36m3. Top CPU Consuming Processes:\033[0m")
	cpuOut, err := exec.Command("sh", "-c", "ps -eo pcpu,pmem,comm -r | head -n 6").Output()
	if err == nil {
		lines := strings.Split(strings.TrimSpace(string(cpuOut)), "\n")
		for i, line := range lines {
			if i == 0 {
				fmt.Printf("   \033[1m%-8s %-8s %s\033[0m\n", "CPU %", "MEM %", "PROCESS")
			} else {
				fmt.Printf("   %s\n", line)
			}
		}
	}
	fmt.Println()

	fmt.Println("\033[1;36m4. Background Homebrew Services:\033[0m")
	brewOut, err := exec.Command("brew", "services", "list").Output()
	if err == nil {
		lines := strings.Split(strings.TrimSpace(string(brewOut)), "\n")
		for _, line := range lines {
			fmt.Printf("   %s\n", line)
		}
	} else {
		fmt.Println("   No Homebrew services active.")
	}

	fmt.Println()
	fmt.Println("\033[1;32m[nara]\033[0m Audit scan completed!")
	return nil
}
