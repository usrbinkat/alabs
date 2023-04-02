package main

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	i3, error := newGameItem(200, 300)
	checkErr(1, error, "newGameItem failed")

	i3.Move(10, 10)
	checkErr(0, error, "Move failed")

	log.Printf("i3: %#v", i3)

	p1 := Player{
		Name: "Bob",
		GameItem: GameItem{
			X: 10,
			Y: 20,
		},
	}
	log.Printf(fmt.Sprintf("p1: %#v, xy: %#v", p1.Name, p1.GameItem))

	ms := []mover{
		i3,
		&p1,
	}

	moveAll(ms, 0, 0)
	for _, m := range ms {
		log.Printf("m: %#v", m)
	}

	log.Printf("Jade: %d", Jade)

	if err := p1.FoundKey(Jade); err != nil {
		log.Error().Err(err).Msg("Failed to add key")
	}
	log.Printf("p1 keys: %d", p1.Keys)
}

func checkErr(lvl int, err error, msg string) {
	if err == nil {
		return
	}

	switch lvl {
	case 1:
		log.Fatal().Err(err).Msg(msg)
	case 2:
		log.Error().Err(err).Msg(msg)
	default:
		log.Warn().Err(err).Msg(msg)
	}
}

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

func newGameItem(x, y int) (*GameItem, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds: %d/%d", x, y, maxX, maxY)
	}

	i := GameItem{
		X: x,
		Y: y,
	}

	return &i, nil
}

func (i *GameItem) Move(x, y int) {
	i.X += x
	i.Y += y
}

const (
	maxX = 1000
	maxY = 600
)

func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("invalid key: %#v", k)
	}

	if !containsKey(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}
	return nil
}

func containsKey(keys []Key, k Key) bool {
	for _, key := range keys {
		if key == k {
			return true
		}
	}
	return false
}

type mover interface {
	Move(int, int)
}

// GameItem represents an item in the game with X and Y coordinates.
type GameItem struct {
	X int
	Y int
}

type Player struct {
	Name string
	Keys []Key
	GameItem
}
