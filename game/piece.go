package game

import (
	"image/color"
	"math/rand"
	"time"
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

type pieceFactory struct {
	pieces []*Piece
}

func newPieceFactory() *pieceFactory {
	pf := &pieceFactory{}
	pf.init()
	return pf
}

func (pf *pieceFactory) init() {
	rand.Seed(time.Now().Unix())
	pieces := []*Piece{
		iPiece(),
		jPiece(),
		lPiece(),
		oPiece(),
		sPiece(),
		tPiece(),
		zPiece(),
	}
	rand.Shuffle(len(pieces), func(i, j int) { pieces[i], pieces[j] = pieces[j], pieces[i] })
	pf.pieces = pieces
}

func (pf *pieceFactory) randomPiece() *Piece {
	if len(pf.pieces) == 0 {
		pf.init()
	}
	p := pf.pieces[0]
	pf.pieces = pf.pieces[1:]
	// Spawns the block in the middle of the canvas
	for _, b := range p.Blocks {
		b.Col += COLS/2 - 1
	}
	return p
}

// The pivot point is always at index 0.
type Piece struct {
	T      string
	Blocks []*Block
}

func (p *Piece) Clone() *Piece {
	piece := &Piece{
		T: p.T,
		Blocks: make([]*Block, 4),
	}
	for i, block := range p.Blocks {
		piece.Blocks[i] = block.clone()
	}
	return piece
}

func (p *Piece) MoveDown() *Piece {
	return p.move(1, 0)
}

func (p *Piece) MoveLeft() *Piece {
	return p.move(0, -1)
}

func (p *Piece) MoveRight() *Piece {
	return p.move(0, 1)
}

func (p *Piece) move(x, y int) *Piece {
	blocks := make([]*Block, 4)
	for i := range p.Blocks {
		blocks[i] = &Block{
			Row:      p.Blocks[i].Row + x,
			Col:      p.Blocks[i].Col + y,
			Color:    p.Blocks[i].Color,
			Inactive: p.Blocks[i].Inactive,
		}
	}
	return &Piece{
		T:      p.T,
		Blocks: blocks,
	}
}

func (p *Piece) Rotate() *Piece {
	if p.T == "O" {
		return p
	}

	blocks := make([]*Block, 4)
	pivot := p.Blocks[0]
	blocks[0] = pivot

	for i := range blocks {
		if i == 0 {
			continue
		}
		dRow := pivot.Row - p.Blocks[i].Row
		dCol := pivot.Col - p.Blocks[i].Col
		blocks[i] = &Block{
			Row:      pivot.Row + (dCol * -1),
			Col:      pivot.Col + dRow,
			Color:    p.Blocks[0].Color,
			Inactive: p.Blocks[0].Inactive,
		}
	}
	return &Piece{
		T:      p.T,
		Blocks: blocks,
	}
}

func (p *Piece) setInactive() {
	for _, b := range p.Blocks {
		b.Inactive = true
	}
}

func iPiece() *Piece {
	return &Piece{
		T: "I",
		Blocks: []*Block{
			{Row: 1, Col: 1, Color: mint, Inactive: false},
			{Row: 1, Col: 2, Color: mint, Inactive: false},
			{Row: 1, Col: 0, Color: mint, Inactive: false},
			{Row: 1, Col: 3, Color: mint, Inactive: false},
		},
	}
}

func lPiece() *Piece {
	return &Piece{
		T: "L",
		Blocks: []*Block{
			{Row: 1, Col: 1, Color: bittersweet, Inactive: false},
			{Row: 1, Col: 0, Color: bittersweet, Inactive: false},
			{Row: 1, Col: 2, Color: bittersweet, Inactive: false},
			{Row: 0, Col: 2, Color: bittersweet, Inactive: false},
		},
	}
}

func jPiece() *Piece {
	return &Piece{
		T: "J",
		Blocks: []*Block{
			{Row: 1, Col: 1, Color: bluejeans, Inactive: false},
			{Row: 1, Col: 0, Color: bluejeans, Inactive: false},
			{Row: 0, Col: 0, Color: bluejeans, Inactive: false},
			{Row: 1, Col: 2, Color: bluejeans, Inactive: false},
		},
	}
}

func oPiece() *Piece {
	return &Piece{
		T: "O",
		Blocks: []*Block{
			{Row: 0, Col: 0, Color: sunflower, Inactive: false},
			{Row: 1, Col: 1, Color: sunflower, Inactive: false},
			{Row: 0, Col: 1, Color: sunflower, Inactive: false},
			{Row: 1, Col: 0, Color: sunflower, Inactive: false},
		},
	}
}

func sPiece() *Piece {
	return &Piece{
		T: "S",
		Blocks: []*Block{
			{Row: 1, Col: 1, Color: grass, Inactive: false},
			{Row: 1, Col: 0, Color: grass, Inactive: false},
			{Row: 0, Col: 1, Color: grass, Inactive: false},
			{Row: 0, Col: 2, Color: grass, Inactive: false},
		},
	}
}

func tPiece() *Piece {
	return &Piece{
		T: "T",
		Blocks: []*Block{
			{Row: 1, Col: 1, Color: lavender, Inactive: false},
			{Row: 1, Col: 0, Color: lavender, Inactive: false},
			{Row: 0, Col: 1, Color: lavender, Inactive: false},
			{Row: 1, Col: 2, Color: lavender, Inactive: false},
		},
	}
}

func zPiece() *Piece {
	return &Piece{
		T: "Z",
		Blocks: []*Block{
			{Row: 1, Col: 1, Color: ruby, Inactive: false},
			{Row: 0, Col: 0, Color: ruby, Inactive: false},
			{Row: 0, Col: 1, Color: ruby, Inactive: false},
			{Row: 1, Col: 2, Color: ruby, Inactive: false},
		},
	}
}
