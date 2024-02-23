package engine

import (
	"container/list"
	"log"
	"regexp"
)

func SearchForWord(pattern string, words []string) *list.List {

	regex := TransformToRegEx(pattern)
	matches := list.New()

	for _, word := range words {

		found, err := regexp.MatchString(regex, word)

		if err != nil {
			log.Fatal(err)
		}

		if found {
			matches.PushBack(word)
		}
	}

	return matches
}
