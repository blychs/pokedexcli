package main

import ("strings")

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
