package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "stats"
	l := len(s) * 2
	banner(s, l)

	fmt.Println("len: ", len(s))

	for i, r := range s {
		c := string(r)
		fmt.Println(i, r, c)
	}

	b := s[0]
	fmt.Printf("b: %c of type %T\n", b, b)

	p := isPalindrome(s)
	fmt.Printf("'%s' is palindrome: %v\n", s, p)
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func banner(text string, width int) {
	// padding := (width - len(text)) / 2
	padding := (width - utf8.RuneCountInString(text)) / 2

	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
