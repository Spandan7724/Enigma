package wordlist

import (
    "bufio"
    "fmt"
    "os"
)

// LoadWordlist loads a wordlist from a given file path and returns a slice of words.
func LoadWordlist(filePath string) ([]string, error) {
    var words []string

    file, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("could not open wordlist file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        words = append(words, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("error reading wordlist file: %v", err)
    }

    return words, nil
}

// GenerateVariations generates common variations of a given word.
func GenerateVariations(word string) []string {
    return []string{word}
}
