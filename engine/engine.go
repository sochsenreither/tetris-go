package engine

import (
	"github.com/sochsenreither/tetris-go/ai"
	"github.com/sochsenreither/tetris-go/game"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	scale   = 30
	WIDTH   = 1024
	HEIGHT  = 900
	xOffset = WIDTH/2 - game.COLS*scale/2
	yOffset = HEIGHT/2 - game.ROWS*scale/2
	preview = true
)

type Engine struct {
	ai        *ai.TetrisPlayer
	window    *sdl.Window
	renderer  *sdl.Renderer
	game      *game.Game
	font      *ttf.Font
	fontSmall *ttf.Font
	pause     bool
	running   bool
}

func NewEngine() (*Engine, error) {
	game := game.NewGame()
	window, err := sdl.CreateWindow("tetris", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WIDTH, HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}

	if err = ttf.Init(); err != nil {
		panic(err)
	}

	font, err := ttf.OpenFont("font/SourceSansPro-Regular.otf", 48)
	if err != nil {
		panic(err)
	}

	font1, err := ttf.OpenFont("font/SourceSansPro-Regular.otf", 32)
	if err != nil {
		panic(err)
	}

	engine := &Engine{
		ai:        &ai.TetrisPlayer{},
		window:    window,
		renderer:  renderer,
		game:      game,
		font:      font,
		fontSmall: font1,
		pause:     false,
		running:   true,
	}

	return engine, nil
}

func (e *Engine) Run() {
	defer e.window.Destroy()
	defer e.renderer.Destroy()
	defer ttf.Quit()

	counter := 0

	for e.running {
		interval := 50 - (int(e.game.Level) * 3)
		if interval < 5 {
			interval = 5
		}
		e.renderer.SetDrawColor(35, 35, 35, 0)
		e.renderer.Clear()
		// Get direction
		// direction := e.getPlayerMove()
		direction := e.getPlayerMove()
		// direction := e.ai.NextMove(e.game.Board, e.game.ActivePiece)

		// Pass direction to game
		if !e.pause && !e.game.Gameover {
			var tick bool
			counter += 1
			if counter >= interval {
				tick = true
				counter = 0
			}
			e.game.Step(direction, tick)
		}

		e.render()
		sdl.Delay(16)
	}
}

func (e *Engine) getPlayerMove() string {
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
				return "UP"
			case sdl.K_DOWN:
				return "DOWN"
			case sdl.K_RIGHT:
				return "RIGHT"
			case sdl.K_LEFT:
				return "LEFT"
			case sdl.K_SPACE:
				return "SPACE"
			case sdl.K_p:
				e.pause = !e.pause
			}
		}
	}
	return ""
}
