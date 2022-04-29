package main

import "github.com/sochsenreither/tetris-go/engine"

// TODO: AI

func main() {
	e, err := engine.NewEngine()
	if err != nil {
		panic(err)
	}
	e.Run()
}