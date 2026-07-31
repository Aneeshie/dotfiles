package main

import (
	"fmt"
	"os"

	"nara-cli/cmd"
)

const version = "0.1.0"

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	subcommand := os.Args[1]

	switch subcommand {
	case "rebuild", "sync":
		if err := cmd.RunRebuild(); err != nil {
			fmt.Printf("\033[31mError:\033[0m %v\n", err)
			os.Exit(1)
		}

	case "doctor":
		if err := cmd.RunDoctor(); err != nil {
			fmt.Printf("\033[31mError:\033[0m %v\n", err)
			os.Exit(1)
		}

	case "new":
	  projectName := ""
    if len(os.Args) >= 3 {
        projectName = os.Args[2]
    }
    if err := cmd.RunNew(projectName); err != nil {
        fmt.Printf("\033[31mError:\033[0m %v\n", err)
        os.Exit(1)
    }

	case "version", "-v", "--version":
		fmt.Printf("nara-cli version %s\n", version)

	case "help", "-h", "--help":
		printHelp()

	default:
		fmt.Printf("\033[31mError:\033[0m unknown command '%s'\n\n", subcommand)
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("\033[1;35mnara-cli\033[0m - Personal Development & System CLI")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("  nara <command>")
	fmt.Println()
	fmt.Println("AVAILABLE COMMANDS:")
	fmt.Println("  rebuild, sync  Fix permissions and rebuild Nix-Darwin configuration")
	fmt.Println("  version        Show CLI version")
	fmt.Println("  help           Show this help menu")
}
