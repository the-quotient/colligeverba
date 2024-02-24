package engine

import (
	"bufio"
	"container/list"
	"embed"
	"fmt"
	"regexp"
)

var (
	//go:embed assets/*
	content embed.FS
)

func Search(pattern string) []string {

	if !InputValidation(pattern) {
		errorMessage := []string{"Invalid input"}
		return errorMessage[:]
	}

	regex := TransformToRegEx("^" + pattern + "$")
	matches := SearchInFile(regex)

	return TransformToArray(matches)
}

func SearchInFile(regex *regexp.Regexp) *list.List {

	matches := list.New()
	errorMessage := list.New()

	//file, err := os.Open("../latin.txt")
	file, err := content.Open("assets/latin.txt")
	if err != nil {
		fmt.Println("Error:", err)
		errorMessage.PushBack("Error: Problem opening the file")
		return errorMessage
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if regex.MatchString(line) {
			matches.PushBack(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		errorMessage.PushBack("Error: Problem reading the file")
		return errorMessage
	}

	if matches.Len() == 0 {
		errorMessage.PushBack("No matches found")
		return errorMessage
	}

	return matches
}

func InputValidation(pattern string) bool {

	if len(pattern) == 0 {
		return false
	}

	regex := regexp.MustCompile("^([a-zA-Z]+|(\\?{})+|(\\?{([a-zA-Z],)+|[a-zA-Z]+})+|(\\?\\?)+)+$")

	if !regex.MatchString(pattern) {
		return false
	}

	return true
}
