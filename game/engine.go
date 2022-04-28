package game

import (
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// TODO: higher level -> faster
// TODO: When game over press start for new game
// TODO: press space to drop
// TODO: persistent high scores

const (
	scale   = 30
	WIDTH   = 1024
	HEIGHT  = 900
	ROWS    = 22
	COLS    = 10
	xOffset = WIDTH/2 - COLS*scale/2
	yOffset = HEIGHT/2 - ROWS*scale/2
	preview = true
)

type Engine struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	game     *game
	font     *ttf.Font
	font1    *ttf.Font
	pause    bool
	running  bool
}

func NewEngine() (*Engine, error) {
	game := NewGame()
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
		if !e.pause && !e.game.gameover {
			var tick bool
			counter += 1
			if counter == interval {
				tick = true
				counter = 0
			}
			e.game.step(direction, tick)
		}

		e.render()
		sdl.Delay(16)
	}
}

func (e *Engine) render() {
	// Draw canvas
	e.renderer.SetDrawColor(45, 45, 45, 255)
	e.renderer.FillRect(&sdl.Rect{
		X: xOffset - 6,
		Y: yOffset - 6,
		W: COLS*scale + 12,
		H: ROWS*scale + 12,
	})
	for y, row := range e.game.board.canvas {
		for x, block := range row {
			if block != nil {
				e.renderBlock(block)
			} else {
				e.renderCanvas(x, y)
			}
		}
	}

	if preview {
		e.renderNextPiece(e.game.nextPiece)
	}
	e.renderStats()

	if e.pause {
		e.renderPause()
	}

	if e.game.gameover {
		e.renderGameOver()
	}

	// Draw stats
	e.renderer.Present()
}

func (e *Engine) renderNextPiece(p *piece) {
	for x := 0; x < 4; x++ {
		for y := 0; y < 2; y++ {
			for _, b := range p.blocks {
				if b.col-4 == x && b.row == y {
					block := b.clone()
					block.col += COLS/2 + 2
					e.renderBlock(block)
				}
			}
		}
	}
}

func (e *Engine) renderPause() {
	e.renderTextBox("Pause")
}

func (e *Engine) renderGameOver() {
	e.renderTextBox("Game Over")
}

func (e *Engine) renderTextBox(text string) {
	e.renderer.SetDrawColor(135, 135, 135, 0)
	w := int32(COLS * scale)
	h := int32(4 * scale)
	x := int32(WIDTH/2 - COLS*scale/2)
	y := int32(HEIGHT/2 - h - h/2)
	e.renderer.FillRect(&sdl.Rect{
		X: x - 6,
		Y: y - 6,
		W: w + 12,
		H: h + 12,
	})
	e.renderer.SetDrawColor(35, 35, 35, 0)
	e.renderer.FillRect(&sdl.Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	})
	e.renderText(text, x+w/2, y, e.font)
}

func (e *Engine) renderText(text string, x, y int32, font *ttf.Font) error {
	stext, err := font.RenderUTF8Blended(text, sdl.Color{
		R: 135,
		G: 135,
		B: 135,
		A: 0,
	})
	if err != nil {
		return err
	}
	defer stext.Free()
	texture, err := e.renderer.CreateTextureFromSurface(stext)
	if err != nil {
		return err
	}
	defer texture.Destroy()
	e.renderer.Copy(texture, nil, &sdl.Rect{
		X: x - stext.W/2,
		Y: y + stext.H/2 - scale/10,
		W: stext.W,
		H: stext.H,
	})
	return nil
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
	e.renderer.SetDrawColor(35, 35, 35, 150)
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
	e.renderer.SetDrawColor(35, 35, 35, 150)
	e.renderer.DrawRect(rect)
}

func (e *Engine) renderStats() {
	// w := int32(COLS * scale)
	h := int32(4 * scale)
	x := int32(WIDTH/2 - COLS*scale/2)
	y := int32(HEIGHT/2 - 3*h)
	e.renderText("Level:", x-3*scale, y, e.font1)
	e.renderText(strconv.Itoa(int(e.game.level)), x-3*scale, y+scale+scale/2, e.font1)

	e.renderText("Score:", x-3*scale, y+h, e.font1)
	e.renderText(strconv.Itoa(int(e.game.score)), x-3*scale, y+scale+scale/2+h, e.font1)
}
