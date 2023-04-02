package main

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var i any

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	i = 7
	log.Printf("i is %v", i)
	logTypeInfo(i)

	i = "hello"
	log.Printf("i is %v", i)
	logTypeInfo(i)
}

// Logs the value and its type information.
func logTypeInfo(value any) {
	log.Printf("Value info: %v", TypeInfo(value))
}

// represents a value and its type information as a string.
func TypeInfo(value any) string {

	/*
		Third iteration
	*/
	switch v := value.(type) {
	case int:
		return fmt.Sprintf("an int: %d", v)
	case string:
		return fmt.Sprintf("a string: %s", v)
	default:
		return fmt.Sprintf("unknown type: %v", value)
	}

	/*
		Second iteration
	*/
	//  switch i.(type) {
	//  case int:
	//  	return fmt.Sprintf("an int: %d", i.(int))
	//  case string:
	//  	return fmt.Sprintf("a string: %s", i.(string))
	//  default:
	//  	return fmt.Sprintf("unknown type: %v", i)
	//  }

	/*
		First iteration
	*/
	//	if n, ok := i.(int); ok {
	//		return fmt.Sprintf("an int: %d", n)
	//	} else if s, ok := i.(string); ok {
	//
	//		return fmt.Sprintf("a string: %s", s)
	//	} else {
	//
	//		return fmt.Sprintf("unknown: %v", i)
	//	}
}
