package main

import (
	"flag"

	"github.com/sochsenreither/tetris-go/engine"
)

func main() {
	ai := flag.Bool("ai", false, "Run tetris with the ai")
	flag.Parse()

	e, err := engine.NewEngine()
	if err != nil {
		panic(err)
	}
	if *ai {
		e.RunAI()
	} else {
		e.Run()
	}
}