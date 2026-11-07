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

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "there is a weird mistake here")
		}
		userInput := scanner.Text()
		fmt.Printf("Your command was %s\n", cleanInput(userInput)[0])
	}
}
