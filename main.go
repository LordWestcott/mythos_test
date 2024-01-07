package main

import "fmt"

// myth: This is a test on one line

// myth: This is a test
// but on two lines

// myth: This is a test
// on three lines
// with something directly beneath it
var x int

/* myth:
	This is a multiline comment test using block on one line
*/

/* myth:
	This is a multiline comment test using block on one line, with something directly beneath it
*/
var y int

/*
	myth: This is a multiline comment block test, on one line, however myth isn't on the first line, and keyword is used twice.
*/

/* myth:
	This is a multiline comment block.
	It has been split out onto 3 lines.
	Check this out
*/

// Myth: uppercase first letter of keyword

// MYTH: uppercase all letters of keyword




func main() {
	fmt.Println("Hello World")
}