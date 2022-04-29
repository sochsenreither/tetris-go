package game

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := newBoard()

	if len(board.canvas) != ROWS {
		t.Errorf("Wrong length, expected %d, got %d", ROWS, len(board.canvas))
	}

	for _, row := range board.canvas {
		if len(row) != COLS {
			t.Errorf("Wrong row length, expected %d, got %d", COLS, len(row))
		}
	}
}

func TestDrawPiece(t *testing.T) {
	b1 := newBoard()
	b2 := newBoard()
	b1.drawPiece(randomPiece())

	if reflect.DeepEqual(b1.canvas, b2.canvas) {
		t.Errorf("Expected a different canvas")
	}
}

func TestRemovePiece(t *testing.T) {
	b1 := newBoard()
	b2 := newBoard()
	p := randomPiece()
	b1.drawPiece(p)
	b1.removePiece(p)

	if !reflect.DeepEqual(b1.canvas, b2.canvas) {
		t.Errorf("Expected the same canvas")
	}
}

func TestClearLines(t *testing.T) {
	b := newBoard()
	for i := 0; i < COLS; i++ {
		b.canvas[ROWS-1][i] = &Block{}
		b.canvas[ROWS-2][i] = &Block{}
		b.canvas[ROWS-3][i] = &Block{}
	}
	b.clearLines()
	if b.canvas[ROWS-1][1] != nil {
		t.Errorf("Expected row to slide down")
	}
}

func TestCollision(t *testing.T) {
	t.Run("collision with other piece", func(t *testing.T) {
		b := newBoard()
		p := oPiece()
		o1 := oPiece()
		for _, block := range p.blocks {
			block.inactive = true
		}
		b.drawPiece(p)
		if !b.collision(o1) {
			t.Errorf("Expected a collision")
		}
	})

	t.Run("collision out of bounds", func(t *testing.T) {
		b := newBoard()
		p := randomPiece()
		p.blocks[0].col, p.blocks[0].row = -1, 0
		if !b.collision(p) {
			t.Errorf("Expected a collision")
		}
		p.blocks[0].col, p.blocks[0].row = 0, ROWS
		if !b.collision(p) {
			t.Errorf("Expected a collision")
		}
	})
	t.Run("no collision", func(t *testing.T) {
		b := newBoard()
		p := randomPiece()
		if b.collision(p) {
			t.Errorf("Didn't expect a collision")
		}
	})
}

func TestCanClearLine(t *testing.T) {
	b := newBoard()
	if b.canClearLine(0) {
		t.Errorf("Expected false when checking if line 0 can be cleared")
	}
}

func debugPrint(canvas [][]*Block) {
	for _, row := range canvas {
		for _, b := range row {
			if b == nil {
				fmt.Printf("_")
			} else {
				fmt.Printf("o")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
