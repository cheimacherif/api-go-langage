package main

import (
	dictionnary "exercice2/dictionnary"
	"fmt"
)

func main() {
	// Instantiate a new Dictionary
	dict := dictionnary.NewDictionary("dict.txt")

	// Add some words and definitions
	dict.Add("golang", "The Go Programming Language")
	dict.Add("docker", "An open platform for developing, shipping, and running applications")
	dict.Add("kubernetes", "An open-source container orchestration system for automating computer application deployment, scaling, and management")

	// Display the definition of a specific word
	word := "golang"
	definition, err := dict.Get(word)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("The definition of '%s' is: %s\n", word, definition)
	}

	// Remove a word from the dictionary
	wordToRemove := "docker"
	err = dict.Remove(wordToRemove)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("has been removed from the dictionary.\n")
	}}








		