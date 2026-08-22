package main

import (
	"fmt"
	"os"
)

const version = "0.1.0"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	switch args[0] {
	case "version":
		fmt.Println("devvsite", version)

	case "help", "--help", "-h":
		printHelp()

	case "db":
		fmt.Println("Database")

	default:
		fmt.Printf("Unknown command: %s\n", args[0])
		fmt.Println("Run 'devvsite help' for usage.")
	}
}

func printHelp() {
	fmt.Println("devvsite - PHP development environment CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  devvsite <command>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  help       Show this help message")
	fmt.Println("  version    Show devvsite version")
	fmt.Println("  db         Database command")
}
