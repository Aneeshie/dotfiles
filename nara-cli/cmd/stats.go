package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunStats() error {
	homeDir, _ := os.UserHomeDir()
	dotfilesDir := filepath.Join(homeDir, ".dotfiles")

	osVersionOut, _ := exec.Command("sw_vers", "-productVersion").Output()
	osVersion := strings.TrimSpace(string(osVersionOut))

	duOut, _ := exec.Command("du", "-sh", "/nix/store").Output()
	storeSize := "Unknown"
	if len(duOut) > 0 {
		fields := strings.Fields(string(duOut))
		if len(fields) > 0 {
			storeSize = fields[0]
		}
	}

	gitStatusOut, _ := exec.Command("git", "-C", dotfilesDir, "status", "--porcelain").Output()
	isClean := len(strings.TrimSpace(string(gitStatusOut))) == 0

	gitBranchOut, _ := exec.Command("git", "-C", dotfilesDir, "branch", "--show-current").Output()
	gitBranch := strings.TrimSpace(string(gitBranchOut))
	if gitBranch == "" {
		gitBranch = "main"
	}

	fmt.Println("\033[1;35m")
	fmt.Println("  _  _   __   ____   __   ")
	fmt.Println(" ( \\( ) / _\\ (  _ \\ / _\\  ")
	fmt.Println("  )  ( /    \\ )   //    \\ ")
	fmt.Println(" (_)\\_)\\_/\\_/(__\\_)\\_/\\_/ ")
	fmt.Println("\033[0m")

	fmt.Println("\033[1;36m┌──────────────────────────────────────────────────────────┐\033[0m")
	fmt.Println("\033[1;36m│\033[0m \033[1mSYSTEM & NIX DASHBOARD\033[0m                                   \033[1;36m│\033[0m")
	fmt.Println("\033[1;36m├──────────────────────────────────────────────────────────┤\033[0m")
	fmt.Printf("\033[1;36m│\033[0m  💻  \033[1mOS Version\033[0m    : macOS %-29s \033[1;36m│\033[0m\n", osVersion)
	fmt.Printf("\033[1;36m│\033[0m  📦  \033[1mNix Store\033[0m     : %-35s \033[1;36m│\033[0m\n", storeSize)
	if isClean {
		fmt.Printf("\033[1;36m│\033[0m  🌿  \033[1mDotfiles Git\033[0m  : \033[32mClean\033[0m (branch: %-18s) \033[1;36m│\033[0m\n", gitBranch)
	} else {
		fmt.Printf("\033[1;36m│\033[0m  🌿  \033[1mDotfiles Git\033[0m  : \033[33mDirty\033[0m (branch: %-18s) \033[1;36m│\033[0m\n", gitBranch)
	}
	fmt.Println("\033[1;36m└──────────────────────────────────────────────────────────┘\033[0m")
	fmt.Println()

	return nil
}
