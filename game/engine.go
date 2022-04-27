package game

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type Engine struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	game     *Game
}

func NewEngine(game *Game) (*Engine, error) {
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
	}

	return engine, nil
}

func (e *Engine) Run() {
	defer e.window.Destroy()
	defer e.renderer.Destroy()

	running := true
	pause := false

	for running {
		e.renderer.SetDrawColor(0,0,0,0)
		e.renderer.Clear()
		// Get direction
		direction := ""
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if t.Type != sdl.KEYDOWN {
					break
				}
				switch t.Keysym.Sym {
				case sdl.K_ESCAPE:
					running = false
				case sdl.K_UP:
					direction = "UP"
				case sdl.K_DOWN:
					direction = "DOWN"
				case sdl.K_RIGHT:
					direction = "RIGHT"
				case sdl.K_LEFT:
					direction = "LEFT"
				case sdl.K_p:
					pause = !pause
				}
			}
		}

		// Pass direction to game
		if !pause {
			e.game.step(direction)
		}

		e.render()
	}
}

func (e *Engine) render() {
	drawCanvas(e.game.board.canvas)
	// Draw canvas in the middle, next block and score besides
	for _, row := range e.game.board.canvas {
		for _, block := range row {
			if block != nil {
				e.renderBlock(block)
			}
		}
	}

	e.renderer.Present()
}

func (e *Engine) renderBlock(b *block) {
	e.renderer.SetDrawColor(255, 255, 255, 120)
	rect := &sdl.Rect{
		X: int32(b.col)*10 + 200,
		Y: int32(b.row)*10 + 200,
		W: 10,
		H: 10,
	}
	e.renderer.FillRect(rect)
}

func (e *Engine) renderStats() {

}

func drawCanvas(canvas [][]*block) {
	for _, r := range canvas {
		log.Println(r)
	}
}