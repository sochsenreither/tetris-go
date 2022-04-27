package main

import "github.com/sochsenreither/tetris-go/game"

func main() {
	g := game.NewGame()
	engine, err := game.NewEngine(g)
	if err != nil {
		panic(err)
	}

	engine.Run()
}