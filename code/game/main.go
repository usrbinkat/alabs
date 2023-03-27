package main

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	//var i1 Loc
	//fmt.Printf("i1: %#v\n", i1)

	//i2 := Loc{1, 2}
	//fmt.Printf("i1: %#v\n", i2)

	i3, error := newLoc(200, 300)
	checkErr(1, error, "newLoc failed")

	i3.Move(10, 10)
	checkErr(0, error, "Move failed")

	log.Printf("i3: %#v", i3)

	p1 := Player{
		Name: "Bob",
		Loc: Loc{
			X: 10,
			Y: 20,
		},
	}
	log.Printf(fmt.Sprintf("p1: %#v, xy: %#v", p1.Name, p1.Loc))

	ms := []mover{
		i3,
		&p1,
	}

	moveAll(ms, 0, 0)
	for _, m := range ms {
		log.Printf("m: %#v", m)
	}

	log.Printf("Jade: %d", Jade)
}

func (k Key) string() string {
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

func checkErr(lvl int, err error, msg string) {
	if err != nil {
		if lvl == 1 {
			log.Fatal().Err(err).Msg(msg)
		} else if lvl == 2 {
			log.Error().Err(err).Msg(msg)
		} else {
			log.Warn().Err(err).Msg(msg)
		}
	}
}

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

func newLoc(x, y int) (*Loc, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds: %d/%d", x, y, maxX, maxY)
	}

	i := Loc{
		X: x,
		Y: y,
	}

	return &i, nil
}

func (i *Loc) Move(x, y int) {
	i.X += x
	i.Y += y
}

const (
	maxX = 1000
	maxY = 600
)

type Key int

const (
	Jade Key = iota + 1
	Copper
	Crystal
)

type mover interface {
	Move(int, int)
}

// Loc is an item in the game
type Loc struct {
	X int
	Y int
}

type Player struct {
	Name string
	Loc
}
