package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func RunDoctor() error {
	fmt.Println("\033[1;36m[nara] Running environment health checks...\033[0m")
	fmt.Println()
	fmt.Println("\033[1m1. Checking Core Tools:\033[0m")
	checkBinaries()
	fmt.Println()
	fmt.Println("\033[1m2. Checking Repository Permissions:\033[0m")
	homeDir, _ := os.UserHomeDir()
	checkGitPermissions(filepath.Join(homeDir, ".dotfiles"))
	fmt.Println()
	return nil
}
func checkBinaries() {
	binaries := []string{"nix","direnv", "git"}
	for _, bin := range binaries {
		_, err := exec.LookPath(bin)
		if err != nil {
			fmt.Printf(" [\033[31mFAIL\033[0m] %s is not installed or not in PATH\n", bin)
		}else {
			fmt.Printf("[\033[32mPASS\033[0m] %s is available\n", bin)
		}

	}
}

func checkGitPermissions(dotfilesDir string) {
	gitDir := filepath.Join(dotfilesDir, ".git")

	info ,err := os.Stat(gitDir)
	if err != nil {
		fmt.Printf("[\033[33mWARN\033[0m] Could not stat %s: %v\n", gitDir, err)
		return
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return
	}

	currentUID := uint32(os.Getuid())
	if stat.Uid != currentUID {
		fmt.Printf("[\033[31mFAIL\033[0m] %s is owned by UID %d (expected %d). Run: sudo chown -R $USER .git\n", gitDir, stat.Uid, currentUID)
	} else {
		fmt.Printf("[\033[32mPASS\033[0m] %s ownership is correct\n", gitDir)
	}
}
