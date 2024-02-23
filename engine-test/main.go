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

	fmt.Println("Test 1: pattern = " + pattern1)
	word1 := engine.Search(pattern1)

	for e := word1.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println()

	fmt.Println("Test 2: pattern = " + pattern2)
	word2 := engine.Search(pattern2)

	for e := word2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println()

	fmt.Println("Test 3: pattern = " + pattern3)
	word3 := engine.Search(pattern3)
	for e := word3.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println()

	fmt.Println("Test 4: pattern = " + pattern4)
	word4 := engine.Search(pattern4)

	for e := word4.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println()

	fmt.Println("Test 5: pattern = " + pattern5)
	word5 := engine.Search(pattern5)

	for e := word5.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println()
}
