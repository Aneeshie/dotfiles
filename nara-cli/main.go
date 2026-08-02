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
		if err := cmd.RunNew(os.Args[2:]); err != nil {
			fmt.Printf("\033[31mError:\033[0m %v\n", err)
			os.Exit(1)
		}

	case "clean":
		if err := cmd.RunClean(os.Args[2:]); err != nil {
			fmt.Printf("\033[31mError:\033[0m %v\n", err)
			os.Exit(1)
		}

	case "run":
		if err := cmd.RunExec(os.Args[2:]); err != nil {
			fmt.Printf("\033[31mError:\033[0m %v\n", err)
			os.Exit(1)
		}

	case "stats":
		if err := cmd.RunStats(); err != nil {
			fmt.Printf("\033[31mError:\033[0m %v\n", err)
			os.Exit(1)
		}

	case "share":
		if err := cmd.RunShare(os.Args[2:]); err != nil {
			fmt.Printf("\033[31mError:\033[0m %v\n", err)
			os.Exit(1)
		}

	case "upgrade":
		if err := cmd.RunUpgrade(); err != nil {
			fmt.Printf("\033[31mError:\033[0m %v\n", err)
			os.Exit(1)
		}

		case "audit":
		if err := cmd.RunAudit(); err != nil {
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
	fmt.Println("  upgrade        Full system update (nix flake update, brew upgrade, rebuild, clean)")
	fmt.Println("  doctor         Check system health and environment status")
	fmt.Println("  new            Scaffold a new project using Nix templates (default, rust, go, python, java)")
	fmt.Println("  clean          Purge old Nix generations and optimize storage")
	fmt.Println("  run            Run any nixpkgs package ephemerally")
	fmt.Println("  stats          Display system & Nix dashboard")
	fmt.Println("  audit          Scan battery status, CPU/RAM hogs, and Homebrew services")
	fmt.Println("  share          Share a file or folder over local Wi-Fi with a QR code")
	fmt.Println("  version        Show CLI version")
	fmt.Println("  help           Show this help menu")
}
