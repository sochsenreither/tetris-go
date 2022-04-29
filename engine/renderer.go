package engine

import (
	"image/color"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func (e *Engine) render() {
	// Draw canvas
	e.renderer.SetDrawColor(45, 45, 45, 255)
	e.renderer.FillRect(&sdl.Rect{
		X: xOffset - 6,
		Y: yOffset - 6,
		W: COLS*scale + 12,
		H: ROWS*scale + 12,
	})
	for y, row := range e.game.Canvas() {
		for x, block := range row {
			if block != nil {
				e.renderBlock(block.Col(), block.Row(), block.Color())
			} else {
				e.renderCanvas(x, y)
			}
		}
	}

	if preview {
		e.renderNextPiece()
	}
	e.renderStats()

	if e.pause {
		e.renderPause()
	}

	if e.game.GameOVer() {
		e.renderGameOver()
	}

	// Draw stats
	e.renderer.Present()
}

func (e *Engine) renderNextPiece() {
	p := e.game.NextPiece()
	for x := 0; x < 4; x++ {
		for y := 0; y < 2; y++ {
			for _, block := range p.Blocks() {
				if block.Col()-4 == x && block.Row() == y {
					e.renderBlock(block.Col() + COLS/2 + 2, block.Row(), block.Color())
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

func (e *Engine) renderBlock(x, y int, color color.Color) {
	r,g,b,a := color.RGBA()
	e.renderer.SetDrawColor(uint8(r), uint8(g), uint8(b), uint8(a))
	rect := &sdl.Rect{
		X: int32(x)*scale + xOffset,
		Y: int32(y)*scale + yOffset,
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
	e.renderText(strconv.Itoa(int(e.game.Level())), x-3*scale, y+scale+scale/2, e.font1)

	e.renderText("Score:", x-3*scale, y+h, e.font1)
	e.renderText(strconv.Itoa(int(e.game.Score())), x-3*scale, y+scale+scale/2+h, e.font1)
}
