package main

import (
	"fmt"

	"github.com/the-quotient/vestigiaverbi/engine"
)

func main() {

	// this is given by the web container to the engine
	pattern1 := "?{}cclesia"
	pattern2 := "?{}c?{}?{}"
	pattern3 := "?{d,j}??o"
	pattern4 := "d?{l,k,j,i}c?{a,b,h,o}"
	pattern5 := "?{}?{}c?{l,m,n}esi?{}?{}"

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

	fmt.Println("Test 1: pattern = " + pattern1)
	word11 := engine.SearchForWord(pattern1, latinBasicForms[:])
	word12 := engine.SearchForWord(pattern1, latinInflectedForms[:])

	for e := word11.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	for e := word12.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println()

	fmt.Println("Test 1: pattern = " + pattern2)
	word21 := engine.SearchForWord(pattern2, latinBasicForms[:])
	word22 := engine.SearchForWord(pattern2, latinInflectedForms[:])

	for e := word21.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	for e := word22.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println()

	fmt.Println("Test 1: pattern = " + pattern3)
	word31 := engine.SearchForWord(pattern3, latinBasicForms[:])
	word32 := engine.SearchForWord(pattern3, latinInflectedForms[:])

	for e := word31.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	for e := word32.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println()

	fmt.Println("Test 1: pattern = " + pattern4)
	word41 := engine.SearchForWord(pattern4, latinBasicForms[:])
	word42 := engine.SearchForWord(pattern4, latinInflectedForms[:])

	for e := word41.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	for e := word42.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println()

	fmt.Println("Test 1: pattern = " + pattern5)
	word51 := engine.SearchForWord(pattern5, latinBasicForms[:])
	word52 := engine.SearchForWord(pattern5, latinInflectedForms[:])

	for e := word51.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	for e := word52.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println()
}
