package engine

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
)

func Search(pattern string) []string {

	regex := TransformToRegEx("^" + pattern + "$")
	matches := SearchInFile(regex)

	return TransformToArray(matches)
}

func SearchInFile(regex *regexp.Regexp) *list.List {

	matches := list.New()

	file, err := os.Open("../latin.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
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
	}

	return matches
}
