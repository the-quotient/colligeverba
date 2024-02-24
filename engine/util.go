package engine

import (
	"container/list"
	"fmt"
)

func StringToList(s string) *list.List {
	l := list.New()

	for _, char := range s {
		l.PushBack(char)
	}

	return l
}

func TransformToArray(l *list.List) []string {

	slice := make([]string, 0, l.Len())

	for e := l.Front(); e != nil; e = e.Next() {
		if str, ok := e.Value.(string); ok {
			slice = append(slice, str)
		} else {
			fmt.Println("Warning: Non-string value encountered in the list, skipping.")
		}
	}

	return slice
}
