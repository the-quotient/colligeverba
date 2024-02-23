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

func insertAtPosition(l *list.List, value interface{}, position int) *list.Element {
	var i int
	for e := l.Front(); e != nil; e = e.Next() {
		if i == position {
			return l.InsertBefore(value, e)
		}
		i++
	}
	return l.PushBack(value)
}
