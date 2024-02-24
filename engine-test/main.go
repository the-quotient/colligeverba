package main

import (
	"fmt"

	"github.com/the-quotient/vestigiaverbi/engine"
)

func main() {

	// this is given by the web container to the engine
	pattern1 := "?{}ccl??si?{a,b}"
	pattern2 := "?{}c?{}?{o,j,l}"
	pattern3 := "?{d,j}??o"
	pattern4 := "d?{l,k,j,i}c?{a,b,h,o}"
	pattern5 := "?{}?{}c?{l,m,n}esi?{}?{}"
	pattern6 := "?{d,j}??.o"
	pattern7 := ""

	fmt.Println("Test 1: pattern = " + pattern1)
	words1 := engine.Search(pattern1)
	for i := range words1 {
		fmt.Println(words1[i])
	}
	fmt.Println()

	fmt.Println("Test 2: pattern = " + pattern2)
	words2 := engine.Search(pattern2)
	for i := range words2 {
		fmt.Println(words2[i])
	}
	fmt.Println()

	fmt.Println("Test 3: pattern = " + pattern3)
	words3 := engine.Search(pattern3)
	for i := range words3 {
		fmt.Println(words3[i])
	}
	fmt.Println()

	fmt.Println("Test 4: pattern = " + pattern4)
	words4 := engine.Search(pattern4)
	for i := range words4 {
		fmt.Println(words4[i])
	}
	fmt.Println()

	fmt.Println("Test 5: pattern = " + pattern5)
	words5 := engine.Search(pattern5)
	for i := range words5 {
		fmt.Println(words5[i])
	}
	fmt.Println()

	fmt.Println("Test 6: pattern = " + pattern6)
	words6 := engine.Search(pattern6)
	for i := range words6 {
		fmt.Println(words6[i])
	}
	fmt.Println()

	fmt.Println("Test 7: pattern = " + pattern7)
	words7 := engine.Search(pattern7)
	for i := range words7 {
		fmt.Println(words7[i])
	}
	fmt.Println()
}
