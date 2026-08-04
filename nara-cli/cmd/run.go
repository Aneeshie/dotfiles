package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunExec(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("package name is required. Usage: nara run <package> [args...]")
	}

	pkgName := args[0]
	extraArgs := args[1:]

	nixTarget := pkgName
	// If it's not a local path and doesn't already contain a flake reference (#), default to nixpkgs
	if pkgName != "." && pkgName[0] != '.' && pkgName[0] != '/' && !strings.Contains(pkgName, "#") {
		nixTarget = fmt.Sprintf("nixpkgs#%s", pkgName)
		fmt.Printf("\033[1;35m⚡ [nara]\033[0m Running ephemeral package \033[1;36m%s\033[0m...\n", nixTarget)
	} else {
		fmt.Printf("\033[1;35m⚡ [nara]\033[0m Running ephemeral package \033[1;36m%s\033[0m...\n", nixTarget)
	}

	cmdArgs := []string{"run", nixTarget}

	if len(extraArgs) > 0 {
		if extraArgs[0] != "--" {
			cmdArgs = append(cmdArgs, "--")
		}
		cmdArgs = append(cmdArgs, extraArgs...)
	}

	cmd := exec.Command("nix", cmdArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
