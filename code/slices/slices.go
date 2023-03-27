package main

import (
	"fmt"
	"reflect"
)

var s []int // s is a slice of ints

func main() {
	s1 := intSlice(s)         // len is nil safe
	fmt.Printf("s: %v\n", s1) // s: []

	fmt.Println(reflect.TypeOf(string(2)))
}

func intSlice(s []int) []int {
	if s == nil {
		fmt.Println("len is nil") // len is nil safe
	} else {
		fmt.Println("len", len(s)) // print len if len is not nil
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}

	s3 := s2[0:4] // s3 is a slice of s2

	return s3
}
