package game

import (
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

func TestDrawBlock(t *testing.T) {
	board := newBoard()
	p := randomPiece()
	board.drawPiece(p)

	blockExists := false
	for i := 0; i < 2; i++ {
		for _, c := range board.canvas[i] {
			if c != nil {
				blockExists = true
			}
		}
	}
	if !blockExists {
		t.Errorf("Expected piece, got empty canvas")
	}
}