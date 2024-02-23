package main

import (
	"fmt"

	"github.com/the-quotient/vestigiaverbi/engine"
)

func main() {

	// this is given by the web container to the engine
	pattern := "??c?esi??"

	latinBasicForms := [...]string{
		"dicere",
		"ecclesia"}

	latinInflectedForms := [...]string{
		"dico",
		"dicis",
		"dicit",
		"dicimus",
		"dicitis",
		"dicunt",
		"ecclesia",
		"ecclesiae",
		"eccleisae",
		"ecclesiam",
		"ecclesia",
		"ecclesiae",
		"ecclesiarum",
		"ecclesiis",
		"ecclesias",
		"ecclesiis"}

	word1 := engine.SearchForWord(pattern, latinBasicForms[:])
	word2 := engine.SearchForWord(pattern, latinInflectedForms[:])

	for e := word1.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	for e := word2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
