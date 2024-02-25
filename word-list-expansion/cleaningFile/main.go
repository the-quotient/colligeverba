package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	filePath := "../latinNew.txt"

	// Open this file and store the words in a string slice
	words, err := ReadWordsFromFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Replace all tabs with empty string and convert to lowercase
	for word := range words {
		words[word] = strings.ToLower(strings.ReplaceAll(words[word], "\t", ""))
	}

	//Remove all lines that contain an underscore character
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "_") {
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	// Remove duplicates and sort the words
	words = removeDuplicatesAndSort(words)

	// Write updated words back to file
	if err := WriteWordsToFile(words, filePath); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

}

func ReadWordsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func WriteWordsToFile(words []string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, word := range words {
		_, err := writer.WriteString(word + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func removeDuplicatesAndSort(input []string) []string {
	// Create a map to store unique strings
	uniqueMap := make(map[string]bool)

	// Iterate over the slice of strings and insert each string into the map
	for _, str := range input {
		uniqueMap[str] = true
	}

	// Extract the unique strings from the map
	uniqueSlice := make([]string, 0, len(uniqueMap))
	for str := range uniqueMap {
		uniqueSlice = append(uniqueSlice, str)
	}

	// Sort the new slice alphabetically
	sort.Strings(uniqueSlice)

	return uniqueSlice
}
