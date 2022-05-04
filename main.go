package main

import "github.com/sochsenreither/tetris-go/engine"

// TODO: AI
// TODO: When game over press start for new game
// TODO: persistent high scores

func main() {
	e, err := engine.NewEngine()
	if err != nil {
		panic(err)
	}
	e.Run()
}