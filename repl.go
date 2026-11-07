package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	rawWords := strings.Split(text, " ")
	cleanWords := []string{}
	for _, word := range rawWords {
		if (word != " ") && (word != "") {
			cleanWords = append(cleanWords, strings.ToLower(word))
		}
	}
	return cleanWords
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "there is a weird mistake here")
		}
		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}
		command, exists := cfg.commandRegistry[userInput[0]]
		if !exists {
			fmt.Println("Your command does not exist. Try something else.")
			continue
		}
		if err := command.callback(cfg); err != nil {
			fmt.Printf("Try something different, %v", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	commandRegistry map[string]cliCommand
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
