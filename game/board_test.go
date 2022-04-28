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
		b.canvas[ROWS-1][i] = &block{}
		b.canvas[ROWS-2][i] = &block{}
		b.canvas[ROWS-3][i] = &block{}
	}
	b.clearLines()
	if b.canvas[ROWS-1][1] != nil {
		t.Errorf("Expected row to slide down")
	}
}

func TestCollision(t *testing.T) {
	// TODO: implement me
}


func debugPrint(canvas [][]*block) {
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
