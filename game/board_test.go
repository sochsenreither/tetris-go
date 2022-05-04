package game

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()

	if len(board.Canvas) != ROWS {
		t.Errorf("Wrong length, expected %d, got %d", ROWS, len(board.Canvas))
	}

	for _, row := range board.Canvas {
		if len(row) != COLS {
			t.Errorf("Wrong row length, expected %d, got %d", COLS, len(row))
		}
	}
}

func TestDrawPiece(t *testing.T) {
	b1 := NewBoard()
	b2 := NewBoard()
	pf := newPieceFactory()
	b1.DrawPiece(pf.randomPiece())

	if reflect.DeepEqual(b1.Canvas, b2.Canvas) {
		t.Errorf("Expected a different canvas")
	}
}

func TestRemovePiece(t *testing.T) {
	b1 := NewBoard()
	b2 := NewBoard()
	pf := newPieceFactory()
	p := pf.randomPiece()
	b1.DrawPiece(p)
	b1.RemovePiece(p)

	if !reflect.DeepEqual(b1.Canvas, b2.Canvas) {
		t.Errorf("Expected the same canvas")
	}
}

func TestClearLines(t *testing.T) {
	b := NewBoard()
	for i := 0; i < COLS; i++ {
		b.Canvas[ROWS-1][i] = &Block{}
		b.Canvas[ROWS-2][i] = &Block{}
		b.Canvas[ROWS-3][i] = &Block{}
	}
	b.clearLines()
	if b.Canvas[ROWS-1][1] != nil {
		t.Errorf("Expected row to slide down")
	}
}

func TestCollision(t *testing.T) {
	t.Run("collision with other piece", func(t *testing.T) {
		b := NewBoard()
		p := oPiece()
		o1 := oPiece()
		for _, block := range p.blocks {
			block.Inactive = true
		}
		b.DrawPiece(p)
		if !b.Collision(o1) {
			t.Errorf("Expected a collision")
		}
	})

	t.Run("collision out of bounds", func(t *testing.T) {
		b := NewBoard()
		pf := newPieceFactory()
		p := pf.randomPiece()
		p.blocks[0].Col, p.blocks[0].Row = -1, 0
		if !b.Collision(p) {
			t.Errorf("Expected a collision")
		}
		p.blocks[0].Col, p.blocks[0].Row = 0, ROWS
		if !b.Collision(p) {
			t.Errorf("Expected a collision")
		}
	})
	t.Run("no collision", func(t *testing.T) {
		b := NewBoard()
		pf := newPieceFactory()
		p := pf.randomPiece()
		if b.Collision(p) {
			t.Errorf("Didn't expect a collision")
		}
	})
}

func TestCanClearLine(t *testing.T) {
	b := NewBoard()
	if b.CanClearLine(0) {
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
