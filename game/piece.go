package game

import (
	"image/color"
	"math/rand"
)

var (
	mint        = color.RGBA{72, 207, 173, 255}
	bluejeans   = color.RGBA{93, 156, 236, 255}
	bittersweet = color.RGBA{252, 110, 81, 255}
	sunflower   = color.RGBA{255, 206, 84, 255}
	grass       = color.RGBA{160, 212, 104, 255}
	lavender    = color.RGBA{172, 146, 236, 255}
	ruby        = color.RGBA{216, 51, 74, 255}
)

// The pivot point is always at index 0.
type piece struct {
	t      string
	blocks []*block
}

type block struct {
	row      int
	col      int
	color    color.Color
	inactive bool
}

func (block *block) getColor() (uint8, uint8, uint8, uint8) {
	r, g, b, a := block.color.RGBA()
	return uint8(r), uint8(g), uint8(b), uint8(a)
}

func (b *block) clone() *block {
	return &block{
		row:      b.row,
		col:      b.col,
		color:    b.color,
		inactive: b.inactive,
	}
}

func randomPiece() *piece {
	pieces := []func() *piece{
		iPiece,
		jPiece,
		lPiece,
		oPiece,
		sPiece,
		tPiece,
		zPiece,
	}
	i := rand.Intn(len(pieces))
	p := pieces[i]()
	// Spawns the block in the middle of the canvas
	for _, b := range p.blocks {
		b.col += 4
	}
	return p
}

func iPiece() *piece {
	return &piece{
		t: "I",
		blocks: []*block{
			{row: 1, col: 2, color: mint, inactive: false},
			{row: 1, col: 1, color: mint, inactive: false},
			{row: 1, col: 0, color: mint, inactive: false},
			{row: 1, col: 3, color: mint, inactive: false},
		},
	}
}

func lPiece() *piece {
	return &piece{
		t: "L",
		blocks: []*block{
			{row: 1, col: 1, color: bittersweet, inactive: false},
			{row: 1, col: 0, color: bittersweet, inactive: false},
			{row: 1, col: 2, color: bittersweet, inactive: false},
			{row: 0, col: 2, color: bittersweet, inactive: false},
		},
	}
}

func jPiece() *piece {
	return &piece{
		t: "J",
		blocks: []*block{
			{row: 1, col: 1, color: bluejeans, inactive: false},
			{row: 1, col: 0, color: bluejeans, inactive: false},
			{row: 0, col: 0, color: bluejeans, inactive: false},
			{row: 1, col: 2, color: bluejeans, inactive: false},
		},
	}
}

func oPiece() *piece {
	return &piece{
		t: "O",
		blocks: []*block{
			{row: 0, col: 0, color: sunflower, inactive: false},
			{row: 1, col: 1, color: sunflower, inactive: false},
			{row: 0, col: 1, color: sunflower, inactive: false},
			{row: 1, col: 0, color: sunflower, inactive: false},
		},
	}
}

func sPiece() *piece {
	return &piece{
		t: "S",
		blocks: []*block{
			{row: 1, col: 1, color: grass, inactive: false},
			{row: 1, col: 0, color: grass, inactive: false},
			{row: 0, col: 1, color: grass, inactive: false},
			{row: 0, col: 2, color: grass, inactive: false},
		},
	}
}

func tPiece() *piece {
	return &piece{
		t: "T",
		blocks: []*block{
			{row: 1, col: 1, color: lavender, inactive: false},
			{row: 1, col: 0, color: lavender, inactive: false},
			{row: 0, col: 1, color: lavender, inactive: false},
			{row: 1, col: 2, color: lavender, inactive: false},
		},
	}
}

func zPiece() *piece {
	return &piece{
		t: "Z",
		blocks: []*block{
			{row: 1, col: 1, color: ruby, inactive: false},
			{row: 0, col: 0, color: ruby, inactive: false},
			{row: 0, col: 1, color: ruby, inactive: false},
			{row: 1, col: 2, color: ruby, inactive: false},
		},
	}
}

func (p *piece) moveDown() *piece {
	return p.move(1, 0)
}

func (p *piece) moveLeft() *piece {
	return p.move(0, -1)
}

func (p *piece) moveRight() *piece {
	return p.move(0, 1)
}

func (p *piece) move(x, y int) *piece {
	blocks := make([]*block, 4)
	for i := range p.blocks {
		blocks[i] = &block{
			row:      p.blocks[i].row + x,
			col:      p.blocks[i].col + y,
			color:    p.blocks[i].color,
			inactive: p.blocks[i].inactive,
		}
	}
	return &piece{
		t:      p.t,
		blocks: blocks,
	}
}

func (p *piece) rotate() *piece {
	if p.t == "O" {
		blocks := make([]*block, 4)
		for i := range blocks {
			blocks[i] = p.blocks[i].clone()
		}
		return &piece{
			t:      p.t,
			blocks: blocks,
		}
	}

	blocks := make([]*block, 4)
	pivot := p.blocks[0].clone()
	blocks[0] = pivot

	for i := range blocks {
		if i == 0 {
			continue
		}
		dRow := pivot.row - p.blocks[i].row
		dCol := pivot.col - p.blocks[i].col
		blocks[i] = &block{
			row:      pivot.row + (dCol * -1),
			col:      pivot.col + dRow,
			color:    p.blocks[0].color,
			inactive: p.blocks[0].inactive,
		}
	}
	return &piece{
		t:      p.t,
		blocks: blocks,
	}
}

func (p *piece) setInactive() {
	for _, b := range p.blocks {
		b.inactive = true
	}
}
