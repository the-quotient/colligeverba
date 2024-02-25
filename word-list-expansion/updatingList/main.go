package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/the-quotient/vestigiaverbi/engine"
)

/*
* This function updates the latin.txt file with the forms
* from the files in the corpus directory.
 */
func main() {

	// Old file with word forms to be updated
	filePath := "latin.txt"

	// Open this file and store the words in a string slice
	words, err := ReadWordsFromFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// the directory of the corpus
	root := "./corpus/"

	// Walk through the directory and read words from each file
	err2 := filepath.WalkDir(root,
		func(path string, d fs.DirEntry, err2 error) error {

			if err2 != nil {
				return err2
			}

			fmt.Println(path)

			// Open a corpus file and store the lines in a string slice
			newWords, err3 := ReadWordsFromFile(path)
			if err3 != nil {
				fmt.Println("Error reading file:", err3)
			}

			wordList := list.New()

			// seperate the words in each line and add them to the list
			for i := range newWords {
				seperatedLine := strings.Split(newWords[i], " ")
				for j := range seperatedLine {
					wordList.PushBack(seperatedLine[j])
				}
			}

			newWordList := engine.TransformToArray(wordList)

			//Remove punctuation, all kind of brackets and numbers
			cleanerRegEx := regexp.MustCompile(`[^\w\s]|\d+|[\[\]{}()<>]`)
			for i := range newWordList {
				newWordList[i] =
					cleanerRegEx.ReplaceAllString(newWordList[i], "")
			}

			//Insert the new words into the old word list
			//if they are not already in there
			for w := range newWordList {
				words = InsertWord(words, newWordList[w])
			}

			fmt.Println(len(words))

			return nil
		})

	if err2 != nil {
		fmt.Println("Error walking through directory:", err)
	}

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

/*
* Function to insert a word into the sorted slice at the correct position
 */
func InsertWord(slice []string, word string) []string {
	i := sort.SearchStrings(slice, word)
	if i < len(slice) && slice[i] == word {
		// Word already exists
		return slice
	}
	// Insert the word at index i
	slice = append(slice[:i], append([]string{word}, slice[i:]...)...)
	return slice
}
