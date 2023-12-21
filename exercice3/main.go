package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	
)

func main() {
	// Instantiate a new Dictionary
	dict := dictionary.NewDictionary("dict.txt")

	// Set up routes
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		word := r.FormValue("word")
		if word == "" {
			http.Error(w, "Word not provided", http.StatusBadRequest)
			return
		}

		err := dict.Add(word)
		if err != nil {
			http.Error(w, "Failed to add word", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		word := r.FormValue("word")
		if word == "" {
			http.Error(w, "Word not provided", http.StatusBadRequest)
			return
		}

		definition, err := dict.Get(word)
		if err != nil {
			http.Error(w, "Failed to get word", http.StatusInternalServerError)
			return
		}

		w.Write([]byte(definition))
	})

	http.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		word := r.FormValue("word")
		if word == "" {
			http.Error(w, "Word not provided", http.StatusBadRequest)
			return
		}

		err := dict.Remove(word)
		if err != nil {
			http.Error(w, "Failed to remove word", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		dictList, err := dict.List()
		if err != nil {
			http.Error(w, "Failed to list words", http.StatusInternalServerError)
			return
		}

		words := make([]string, 0, len(dictList))
		for word := range dictList {
			words = append(words, word)
		}
		sort.Strings(words)

		w.Write([]byte(strings.Join(words, "\n")))
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
}