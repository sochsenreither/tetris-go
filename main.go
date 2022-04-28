package main

import "github.com/sochsenreither/tetris-go/game"

func main() {
	e, err := game.NewEngine()
	if err != nil {
		panic(err)
	}
	e.Run()
}