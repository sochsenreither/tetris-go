package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	scale   = 25
	xOffset = 150
	yOffset = 0
)

type Engine struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	game     *Game
	pause    bool
	running  bool
}

func NewEngine() (*Engine, error) {
	game := NewGame()
	window, err := sdl.CreateWindow("tetris", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 680, 560, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}

	engine := &Engine{
		window:   window,
		renderer: renderer,
		game:     game,
		pause:    false,
		running:  true,
	}

	return engine, nil
}

func (e *Engine) Run() {
	defer e.window.Destroy()
	defer e.renderer.Destroy()

	counter := 0
	// TODO: this is wonky, make it so that the ticks come in relation with fps
	interval := 50

	for e.running {
		e.renderer.SetDrawColor(0, 0, 0, 0)
		e.renderer.Clear()
		// Get direction
		direction := ""
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				e.running = false
			case *sdl.KeyboardEvent:
				if t.Type != sdl.KEYDOWN {
					break
				}
				switch t.Keysym.Sym {
				case sdl.K_ESCAPE:
					e.running = false
				case sdl.K_UP:
					direction = "UP"
				case sdl.K_DOWN:
					direction = "DOWN"
				case sdl.K_RIGHT:
					direction = "RIGHT"
				case sdl.K_LEFT:
					direction = "LEFT"
				case sdl.K_p:
					e.pause = !e.pause
				}
			}
		}

		// Pass direction to game
		if !e.pause {
			var tick bool
			counter += 1
			if counter == interval {
				tick = true
				counter = 0
			}
			err := e.game.step(direction, tick)
			if err != nil {
				fmt.Println("Game over!")
				return
			}
		}

		e.render()
		sdl.Delay(16)
	}
}

func (e *Engine) render() {
	// Draw canvas in the middle, next block and score besides
	for y, row := range e.game.board.canvas {
		for x, block := range row {
			if block != nil {
				e.renderBlock(block)
			} else {
				e.renderCanvas(x, y)
			}
		}
	}

	e.renderer.Present()
}

func (e *Engine) renderCanvas(x, y int) {
	e.renderer.SetDrawColor(75, 75, 75, 255)
	rect := &sdl.Rect{
		X: int32(x*scale + xOffset),
		Y: int32(y*scale + yOffset),
		W: scale,
		H: scale,
	}
	e.renderer.FillRect(rect)
	e.renderer.SetDrawColor(45, 45, 45, 150)
	e.renderer.DrawRect(rect)

}

func (e *Engine) renderBlock(b *block) {
	e.renderer.SetDrawColor(b.getColor())
	rect := &sdl.Rect{
		X: int32(b.col)*scale + xOffset,
		Y: int32(b.row)*scale + yOffset,
		W: scale,
		H: scale,
	}
	e.renderer.FillRect(rect)
	e.renderer.SetDrawColor(45, 45, 45, 150)
	e.renderer.DrawRect(rect)
}

func (e *Engine) renderStats() {

}
