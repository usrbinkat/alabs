# Game Items and Player Movement in Go

This is a simple xy coordinates game demonstration of items and a player with their respective movement capabilities using the Go using structs, interfaces, methods, and error handling.

## Features

1. A game of `X` x `Y` coordinates.
2. A player with name, location, and keys they possess.
3. Player and item movement by `X` and `Y` integers.
4. The player can find keys and add them to their list.
5. Logging of game items, player, and movements.

## Getting Started

### Prerequisites

- Go 1.16 or later.

### Installation

1. Clone the repository:

```bash
git clone https://github.com/usrbinkat/alabs.git
```

2. Change to the project directory:

```bash
cd alabs/code/game
```

### Running the Program

In the project directory, run the following command:

```bash
go run main.go key.go
```

The output will show the game items, player, and movement:

```json
{"level":"debug","time":"123456","message":"i3: &main.GameItem{X:210, Y:310}"}
{"level":"debug","time":"123456","message":"p1: "Bob", xy: main.GameItem{X:10, Y:20}"}
{"level":"debug","time":"123456","message":"m: &main.GameItem{X:210, Y:310}"}
{"level":"debug","time":"123456","message":"m: &main.Player{Name:"Bob", Keys:[]main.Key(nil), GameItem:main.GameItem{X:10, Y:20}}"}
{"level":"debug","time":"123456","message":"Jade: 1"}
{"level":"debug","time":"123456","message":"p1 keys: [1]"}

```
