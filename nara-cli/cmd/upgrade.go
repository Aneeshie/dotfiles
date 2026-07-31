package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunUpgrade() error {
	fmt.Println("\033[1;35m🚀 [nara upgrade]\033[0m Starting full macOS system update...")
	fmt.Println()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	dotfilesDir := filepath.Join(homeDir, ".dotfiles")

	fmt.Println("\033[1;36m[1/4]\033[0m Updating Nix Flake lockfile...")
	flakeCmd := exec.Command("nix", "flake", "update", "--flake", dotfilesDir)
	flakeCmd.Stdout = os.Stdout
	flakeCmd.Stderr = os.Stderr
	if err := flakeCmd.Run(); err != nil {
		fmt.Printf("\033[33mWarning:\033[0m nix flake update encountered an issue: %v\n", err)
	}
	fmt.Println()

	fmt.Println("\033[1;36m[2/4]\033[0m Updating Homebrew packages and casks...")
	brewCmd := exec.Command("brew", "update")
	brewCmd.Stdout = os.Stdout
	brewCmd.Stderr = os.Stderr
	_ = brewCmd.Run()

	brewUpgradeCmd := exec.Command("brew", "upgrade")
	brewUpgradeCmd.Stdout = os.Stdout
	brewUpgradeCmd.Stderr = os.Stderr
	_ = brewUpgradeCmd.Run()
	fmt.Println()

	fmt.Println("\033[1;36m[3/4]\033[0m Rebuilding macOS system configuration...")
	if err := RunRebuild(); err != nil {
		return fmt.Errorf("rebuild phase failed: %w", err)
	}
	fmt.Println()

	fmt.Println("\033[1;36m[4/4]\033[0m Optimising Nix storage...")
	_ = RunClean([]string{})

	fmt.Println()
	fmt.Println("\033[1;32m[nara]\033[0m Full macOS system & package upgrade completed")
	return nil
}
