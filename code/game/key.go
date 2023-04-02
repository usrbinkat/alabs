package main

type Key int

const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey // internal (not exported)
)

func (k Key) String() string {
	switch k {
	case Jade:
		return "Jade"
	case Copper:
		return "Copper"
	case Crystal:
		return "Crystal"
	default:
		return "Unknown"
	}
}
