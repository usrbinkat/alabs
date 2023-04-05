package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// enable debug mode
var debugMode bool = true

// main function
func main() {
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Set log to stdout
	log.Logger = log.Output(os.Stdout)

	// slice of integers
	n := []int{0, 2, 3, 4, 5}

	// for each integer in the slice divide 10 by the integer
	for _, i := range n {
		// call the divide function
		r, error := divide(10, i)

		// log result
		if error != "" {
			logIt(zerolog.ErrorLevel, "int", error, r)
		} else {
			logIt(zerolog.InfoLevel, "int", "divide result is", r)
		}

	}
}

// divide two integers and return the result
func divide(a, b int) (result int, err string) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("divide by zero").Error()
		}
	}()

	result = a / b
	return result, err
}

// logIt logs a message to stdout
func logIt(level zerolog.Level, t string, msg string, value interface{}) {

	if level == zerolog.DebugLevel && !debugMode {
		return
	}

	m := fmt.Sprintf("%s %v", msg, value)

	info := zerolog.Dict().
		Str("type", t).
		Interface("log", m)

	event := log.WithLevel(level).Dict("msg", info)
	event.Msg("")
}
