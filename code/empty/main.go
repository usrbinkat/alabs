package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var v any
var debugMode bool = true

func main() {
	// set time format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// log to stdout
	log.Logger = log.Output(os.Stderr)

	v = 7
	valueType := typeInfo(v)
	logIt(zerolog.InfoLevel, valueType, "v is", v)

	v = "hello"
	valueType = typeInfo(v)
	logIt(zerolog.DebugLevel, valueType, "Debug: v is", v)
}

func typeInfo(value any) string {
	switch value.(type) {
	case int:
		return "int"
	case string:
		return "str"
	default:
		return "unk"
	}
}

// Level based logging
func logIt(level zerolog.Level, t string, msg string, value any) {

	if level == zerolog.DebugLevel && !debugMode {
		return
	}

	m := fmt.Sprintf("%s %s", msg, t)

	info := zerolog.Dict().
		Str("type", t).
		Interface("log", m)

	event := log.WithLevel(level).Dict("msg", info)
	event.Msg("")
}
