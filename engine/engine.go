package engine

import (
	"math/rand"
	"time"

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
	window   *sdl.Window
	renderer *sdl.Renderer
	game     *game.Game
	font     *ttf.Font
	font1    *ttf.Font
	pause    bool
	running  bool
}

func NewEngine() (*Engine, error) {
	rand.Seed(time.Now().Unix())
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
		window:   window,
		renderer: renderer,
		game:     game,
		font:     font,
		font1:    font1,
		pause:    false,
		running:  true,
	}

	return engine, nil
}

func (e *Engine) Run() {
	defer e.window.Destroy()
	defer e.renderer.Destroy()
	defer ttf.Quit()

	counter := 0
	interval := 50

	for e.running {
		e.renderer.SetDrawColor(35, 35, 35, 0)
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
		if !e.pause && !e.game.GameOVer() {
			var tick bool
			counter += 1
			if counter == interval {
				tick = true
				counter = 0
			}
			e.game.Step(direction, tick)
		}

		e.render()
		sdl.Delay(16)
	}
}
