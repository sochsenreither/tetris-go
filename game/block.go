package game

import "image/color"

type Block struct {
	row      int
	col      int
	color    color.Color
	inactive bool
}

func (b *Block) Color() color.Color {
	return b.color
}

func (b *Block) Row() int {
	return b.row
}

func (b *Block) Col() int {
	return b.col
}