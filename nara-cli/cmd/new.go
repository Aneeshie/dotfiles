package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunNew(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("project name is required. Usage: nara new <project-name> [template] (templates: default, rust, go, python, java)")
	}

	projectName := args[0]
	templateName := "default"

	for i, arg := range args {
		if i == 0 {
			continue
		}
		if arg == "-t" || arg == "--template" {
			if i+1 < len(args) {
				templateName = args[i+1]
			}
		} else if !strings.HasPrefix(arg, "-") {
			templateName = arg
		}
	}

	targetDir := projectName
	if projectName != "." {
		if err := os.MkdirAll(targetDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	homeDir, _ := os.UserHomeDir()
	templateDir := filepath.Join(homeDir, ".dotfiles", "templates", templateName)
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		templateDir = filepath.Join(homeDir, "github", "Aneeshie", "dotfiles", "templates", templateName)
		if _, err := os.Stat(templateDir); os.IsNotExist(err) {
			return fmt.Errorf("template '%s' not found. Available templates: default, rust, go, python, java", templateName)
		}
	}

	// copy flake.nix
	fmt.Printf("\033[1;36m[1/4]\033[0m Copying \033[1m%s\033[0m flake.nix template...\n", templateName)
	if err := copyFile(filepath.Join(templateDir, "flake.nix"), filepath.Join(targetDir, "flake.nix")); err != nil {
		return fmt.Errorf("failed to copy flake.nix: %w", err)
	}

	// copy .envrc
	fmt.Printf("\033[1;36m[2/4]\033[0m Copying .envrc...\n")
	_ = copyFile(filepath.Join(templateDir, ".envrc"), filepath.Join(targetDir, ".envrc"))

	// Git Init & Add
	fmt.Printf("\033[1;36m[3/4]\033[0m Initializing Git repository...\n")
	_ = exec.Command("git", "-C", targetDir, "init").Run()
	_ = exec.Command("git", "-C", targetDir, "add", ".").Run()

	// Direnv Allow
	fmt.Printf("\033[1;36m[4/4]\033[0m Enabling direnv...\n")
	_ = exec.Command("direnv", "allow", targetDir).Run()

	fmt.Printf("\033[1;32m✨ Project '%s' (%s template) created successfully!\033[0m\n", projectName, templateName)
	return nil
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, data ,0644)
}
