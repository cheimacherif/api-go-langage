package dictionary

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

type Dictionary struct {
	filename string
	lock     sync.Mutex
	addCh    chan string
	removeCh chan string
}

func NewDictionary(filename string) *Dictionary {
	dict := &Dictionary{filename: filename}
	dict.addCh = make(chan string)
	dict.removeCh = make(chan string)

	go dict.processOperations()

	return dict
}

func (d *Dictionary) processOperations() {
	for {
		select {
		case word := <-d.addCh:
			d.Add(word)
		case word := <-d.removeCh:
			d.Remove(word)
		}
	}
}

func (d *Dictionary) Add(word string) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	// Check if the word is already present
	file, err := os.Open(d.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "=")
		if parts[0] == word {
			return errors.New("word already exists")
		}
	}

	// Append the new word to the file
	file, err = os.OpenFile(d.filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s=definition\n", word))
	return err
}

func (d *Dictionary) Remove(word string) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	file, err := os.Open(d.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "=")
		if parts[0] != word {
			lines = append(lines, scanner.Text())
		}
	}

	file, err = os.Create(d.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := w.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return w.Flush()
}

func (d *Dictionary) List() (map[string]string, error) {
	file, err := os.Open(d.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dict := map[string]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "=")
		dict[parts[0]] = parts[1]
	}

	return dict, nil
}