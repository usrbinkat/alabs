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

	mIntSlice := []int{1, 2, 3, 4, 5}
	mInt := maxInts(mIntSlice)
	valueType = typeInfo(mInt)
	logIt(zerolog.InfoLevel, valueType, "maxInts: mInt is", mInt)

	mFloat64Slice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	mFloat64 := maxFloat64s(mFloat64Slice)
	valueType = typeInfo(mFloat64)
	logIt(zerolog.InfoLevel, valueType, "maxFloat64s: mFloat64 is", mFloat64)

	mMax := []float64{1, 2.2, 3, 4.4, 5, 5.9}
	m := max(mMax)
	valueType = typeInfo(m)
	logIt(zerolog.InfoLevel, valueType, "max: m is", m)
}

func typeInfo(value any) string {
	switch value.(type) {
	case int:
		return "int"
	case string:
		return "str"
	case float64:
		return "float64"
	default:
		return "unk"
	}
}

// Level based logging
func logIt(level zerolog.Level, t string, msg string, value any) {

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

type Number interface {
	int | float64
}

func max[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func maxInts(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func maxFloat64s(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}
