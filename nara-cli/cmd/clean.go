package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func RunClean(args []string) error {
	removeAll := false
	for _, arg := range args {
		if arg == "--all" || arg == "-a" {
			removeAll = true
		}
	}

	fmt.Println("\033[1;35m[nara]\033[0m Starting Nix store cleanup...")
	fmt.Println()

	fmt.Println("\033[1;36m[1/2]\033[0m Deleting old Nix generations...")
	gcArgs := []string{"--delete-older-than", "7d"}
	if removeAll {
		fmt.Println("  (Purging ALL old generations)")
		gcArgs = []string{"-d"}
	}

	gcCmd := exec.Command("nix-collect-garbage", gcArgs...)
	gcCmd.Stdout = os.Stdout
	gcCmd.Stderr = os.Stderr
	if err := gcCmd.Run(); err != nil {
		fmt.Printf("\033[33mWarning:\033[0m garbage collection encountered an error: %v\n", err)
	}

	fmt.Println()
	fmt.Println("\033[1;36m[2/2]\033[0m Optimising Nix store (hard-linking duplicate files)...")
	optCmd := exec.Command("nix", "store", "optimise")
	optCmd.Stdout = os.Stdout
	optCmd.Stderr = os.Stderr
	if err := optCmd.Run(); err != nil {
		fmt.Printf("\033[33mWarning:\033[0m store optimisation encountered an error: %v\n", err)
	}

	fmt.Println()
	fmt.Println("\033[1;32m[nara]\033[0m Cleanup completed successfully!")
	return nil
}
