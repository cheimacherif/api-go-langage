package main

import (
	"fmt"
	"sort"
)

type Dictionary map[string]string

func (d Dictionary) Get(word string) string {
	return d[word]
}

func (d Dictionary) Remove(word string) {
	delete(d, word)
}

func (d Dictionary) List() map[string]string {
	keys := make([]string, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sortedMap := make(map[string]string)
	for _, k := range keys {
		sortedMap[k] = d[k]
	}

	return sortedMap
}

func main() {
	// Instantiate a new Dictionary
	dictionary := Dictionary{}

	// Add some words and definitions
	dictionary["golang"] = "The Go Programming Language"
	dictionary["docker"] = "An open platform for developing, shipping, and running applications"
	dictionary["kubernetes"] = "An open-source container orchestration system for automating computer application deployment, scaling, and management"

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