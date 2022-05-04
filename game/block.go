package game

import "image/color"

type Block struct {
	Row      int
	Col      int
	Color    color.Color
	Inactive bool
}

func (b *Block) clone() *Block {
	return &Block{
		Row:      b.Row,
		Col:      b.Col,
		Color:    b.Color,
		Inactive: b.Inactive,
	}
}
