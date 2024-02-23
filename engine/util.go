package engine

import (
	"container/list"
	"fmt"
)

func ReplaceChar(s string, oldChar byte, newChar byte) string {
	strBytes := []byte(s)

	for i := range strBytes {
		if strBytes[i] == oldChar {
			strBytes[i] = newChar
		}
	}

	return string(strBytes)
}

func StringToList(s string) *list.List {
	l := list.New()

	for _, char := range s {
		l.PushBack(char)
	}

	return l
}

func TransformToArray(l *list.List) []string {
	// Initialize a slice with capacity equal to the length of the list
	slice := make([]string, 0, l.Len())

	// Iterate over the list and append each element to the slice
	for e := l.Front(); e != nil; e = e.Next() {
		// Assuming the list contains string values, assert the type to string
		if str, ok := e.Value.(string); ok {
			slice = append(slice, str)
		} else {
			// Handle the case where the type assertion fails
			fmt.Println("Warning: Non-string value encountered in the list, skipping.")
		}
	}

	return slice
}
