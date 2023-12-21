package dictionary

import (
	"bufio"
	"os"
	"strings"
)

type Dictionary struct {
	filename string
}

func NewDictionary(filename string) *Dictionary {
	return &Dictionary{filename: filename}
}

func (d *Dictionary) Add(word, definition string) error {
	file, err := os.OpenFile(d.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(word + "=" + definition + "\n")
	return err
}

func (d *Dictionary) Get(word string) (string, error) {
	file, err := os.Open(d.filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "=")
		if parts[0] == word {
			return parts[1], nil
		}
	}

	return "", nil
}

func (d *Dictionary) Remove(word string) error {
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






