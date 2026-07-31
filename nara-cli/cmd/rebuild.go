package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunRebuild() error {
	fmt.Println("\033[1;35m[nara]\033[0m Starting system rebuild workflow...")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}

	dotfilesDir := filepath.Join(homeDir, ".dotfiles")

	fmt.Println("\033[1;36m[1/2]\033[0m Staging git changes in ~/.dotfiles...")
	gitCmd := exec.Command("git", "-C", dotfilesDir, "add", ".")
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	if err := gitCmd.Run(); err != nil {
		fmt.Printf("\033[33mWarning:\033[0m git add returned an error (you may need to fix .git permissions): %v\n", err)
	}

	fmt.Println("\033[1;36m[2/2]\033[0m Running darwin-rebuild switch...")
	flakePath := fmt.Sprintf("%s#mac", dotfilesDir)

	rebuildCmd := exec.Command("sudo", "darwin-rebuild", "switch", "--flake", flakePath)
	rebuildCmd.Stdin = os.Stdin
	rebuildCmd.Stdout = os.Stdout
	rebuildCmd.Stderr = os.Stderr

	if err := rebuildCmd.Run(); err != nil {
		return fmt.Errorf("rebuild failed: %w", err)
	}

	fmt.Println("\033[1;32m[nara]\033[0m System rebuild completed successfully!")
	return nil
}
