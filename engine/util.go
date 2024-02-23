package engine

import (
	"container/list"
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
