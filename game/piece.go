package game

import (
	"image/color"
	"math/rand"
)

type piece struct {
	t string
	blocks []*block
}

type block struct {
	row   int
	col   int
	color color.Color
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
	return pieces[i]()
}

func iPiece() *piece {
	return &piece{
		t: "I",
		blocks: []*block{
			{row: 1, col: 0, color: color.White},
			{row: 1, col: 1, color: color.White},
			{row: 1, col: 2, color: color.White},
			{row: 1, col: 3, color: color.White},
		},
	}
}

func lPiece() *piece {
	return &piece{
		t: "L",
		blocks: []*block{
			{row: 0, col: 0, color: color.White},
			{row: 1, col: 0, color: color.White},
			{row: 0, col: 1, color: color.White},
			{row: 0, col: 2, color: color.White},
		},
	}
}

func jPiece() *piece {
	return &piece{
		t: "J",
		blocks: []*block{
			{row: 0, col: 0, color: color.White},
			{row: 1, col: 0, color: color.White},
			{row: 1, col: 1, color: color.White},
			{row: 1, col: 2, color: color.White},
		},
	}
}

func oPiece() *piece {
	return &piece{
		t: "O",
		blocks: []*block{
			{row: 0, col: 0, color: color.White},
			{row: 1, col: 1, color: color.White},
			{row: 0, col: 1, color: color.White},
			{row: 1, col: 0, color: color.White},
		},
	}
}

func sPiece() *piece {
	return &piece{
		t: "S",
		blocks: []*block{
			{row: 1, col: 0, color: color.White},
			{row: 1, col: 1, color: color.White},
			{row: 0, col: 1, color: color.White},
			{row: 0, col: 2, color: color.White},
		},
	}
}

func tPiece() *piece {
	return &piece{
		t: "T",
		blocks: []*block{
			{row: 0, col: 0, color: color.White},
			{row: 0, col: 1, color: color.White},
			{row: 0, col: 2, color: color.White},
			{row: 1, col: 1, color: color.White},
		},
	}
}

func zPiece() *piece {
	return &piece{
		t: "Z",
		blocks: []*block{
			{row: 1, col: 0, color: color.White},
			{row: 0, col: 1, color: color.White},
			{row: 1, col: 1, color: color.White},
			{row: 1, col: 2, color: color.White},
		},
	}
}

func (p *piece) moveDown() *piece {
	blocks := make([]*block, 4)
	for i := range p.blocks {
		blocks[i] = &block{
			row: p.blocks[i].row+1,
			col: p.blocks[i].col,
			color: p.blocks[i].color,
		}
	}
	return &piece{
		t: p.t,
		blocks: blocks,
	}
}

func (p *piece) moveLeft() *piece {
	blocks := make([]*block, 4)
	for i := range p.blocks {
		blocks[i] = &block{
			row: p.blocks[i].row,
			col: p.blocks[i].col-1,
			color: p.blocks[i].color,
		}
	}
	return &piece{
		t: p.t,
		blocks: blocks,
	}
}

func (p *piece) moveRight() *piece {
	blocks := make([]*block, 4)
	for i := range p.blocks {
		blocks[i] = &block{
			row: p.blocks[i].row,
			col: p.blocks[i].col+1,
			color: p.blocks[i].color,
		}
	}
	return &piece{
		t: p.t,
		blocks: blocks,
	}
}

func (p *piece) rotate() *piece {
	return nil
}