package game

import "image/color"

type Block struct {
	Row      int
	Col      int
	Color    color.Color
	Inactive bool
}