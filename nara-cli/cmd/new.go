package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunNew(projectName string) error {
	if projectName == "" {
		return fmt.Errorf("project name is required. Usage: nara new <project-name>")
	}

	targetDir := projectName
	if projectName != "." {
		if err := os.MkdirAll(targetDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	homeDir, _ := os.UserHomeDir()
	templateDir := filepath.Join(homeDir, "github", "Aneeshie", "dotfiles", "templates", "default")

	//copy flake.nix
	fmt.Printf("\033[1;36m[1/4]\033[0m Copying flake.nix template...\n")
	if err := copyFile(filepath.Join(templateDir, "flake.nix"), filepath.Join(targetDir, "flake.nix")); err != nil {
		return fmt.Errorf("failed to copy flake.nix: %w", err)
	}

	//copy .envrc
	fmt.Printf("\033[1;36m[2/4]\033[0m Copying .envrc...\n")
	_ = copyFile(filepath.Join(templateDir, ".envrc"), filepath.Join(targetDir, ".envrc"))

	// Git Init & Add
	fmt.Printf("\033[1;36m[3/4]\033[0m Initializing Git repository...\n")
	_ = exec.Command("git", "-C", targetDir, "init").Run()
	_ = exec.Command("git", "-C", targetDir, "add", ".").Run()

	//  Direnv Allow
	fmt.Printf("\033[1;36m[4/4]\033[0m Enabling direnv...\n")
	_ = exec.Command("direnv", "allow", targetDir).Run()

	fmt.Printf("\033[1;32mProject '%s' created successfully!\033[0m\n", projectName)
	return nil
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, data ,0644)
}
