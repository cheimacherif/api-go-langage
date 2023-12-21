package main

import (
	"fmt"
)

dictionary := Dictionary{}

func main() {
	
	// Add some words and definitions
	dictionary.Add("golang", "The Go Programming Language")
	dictionary.Add("docker", "An open platform for developing, shipping, and running applications")
	dictionary.Add("kubernetes", "An open-source container orchestration system for automating computer application deployment, scaling, and management")

	// Display the definition of a specific word
	word := "golang"
	fmt.Printf("The definition of '%s' is: %s\n", word, dictionary.Get(word))

	// Remove a word from the dictionary
	wordToRemove := "docker"
	dictionary.Remove(wordToRemove)
	fmt.Printf("'%s' has been removed from the dictionary.\n", wordToRemove)

	// Display the list of words and their definitions
	list := dictionary.List()
	for word, definition := range list {
		fmt.Printf("Word: '%s', Definition: '%s'\n", word, definition)
	}
}