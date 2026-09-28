// Package pokedexcli: functions for the repl
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		if err := scanner.Err(); err != nil {
			os.Exit(1)
		}

		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		if err := commandRegistry(words[0]); err != nil {
			fmt.Errorf("Somthing is wrong: %w", err)
		}
	}
}

func cleanInput(text string) []string {
	words := strings.ToLower(text)
	return strings.Fields(words)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandRegistry(text string) error {
	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}
	registry := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	calledCommand, ok := registry[strings.ToLower(text)]
	if !ok {
		fmt.Println("Unknown command")
		return nil
	}
	return calledCommand.callback()
}
