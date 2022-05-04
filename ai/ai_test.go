package ai

import (
	"fmt"
	"testing"

	"github.com/sochsenreither/tetris-go/game"
)

// Note: these tests are hardcoded for a canvas of size 10x22
func TestAggregateHeight(t *testing.T) {
	t.Run("empty board", func(t *testing.T) {
		assert(t, aggregateHeight(game.NewBoard()), 0)
	})

	t.Run("one row", func(t *testing.T) {
		board := game.NewBoard()
		fillRow(board.Canvas[game.ROWS-1], game.ROWS-1)
		assert(t, aggregateHeight(board), 10)
	})

	t.Run("testBoard", func(t *testing.T) {
		board := testBoard()
		assert(t, aggregateHeight(board), 48)
	})
}

func TestCompletedLines(t *testing.T) {
	board := game.NewBoard()
	fillRow(board.Canvas[21], 21)
	fillRow(board.Canvas[20], 20)

	fillRow(board.Canvas[19], 19)
	board.Canvas[19][1] = nil

	assert(t, completeLines(board), 2)
}

func TestHoles(t *testing.T) {
	t.Run("empty board", func(t *testing.T) {
		assert(t, holes(game.NewBoard()), 0)
	})

	t.Run("one row", func(t *testing.T) {
		board := game.NewBoard()
		fillRow(board.Canvas[game.ROWS-1], game.ROWS-1)
		assert(t, holes(board), 0)
	})

	t.Run("testBoard", func(t *testing.T) {
		board := testBoard()
		assert(t, holes(board), 2)
	})
}

func TestBumpiness(t *testing.T) {
	t.Run("empty board", func(t *testing.T) {
		assert(t, bumpiness(game.NewBoard()), 0)
	})

	t.Run("complete rows", func(t *testing.T) {
		board := game.NewBoard()
		fillRow(board.Canvas[21], 21)
		fillRow(board.Canvas[20], 20)
		assert(t, bumpiness(board), 0)
	})

	t.Run("testBoard", func(t *testing.T) {
		board := testBoard()
		assert(t, bumpiness(board), 6)
	})
}

func TestColumnHeight(t *testing.T) {
	t.Run("empty board", func(t *testing.T) {
		assert(t, columnHeight(game.NewBoard(), 0), 0)
	})

	t.Run("some blocks", func(t *testing.T) {
		board := game.NewBoard()
		for i := game.ROWS - 1; i > 5; i-- {
			board.Canvas[i][0] = &game.Block{}
		}

		assert(t, columnHeight(board, 0), 16)
	})
}

//////////////////////
// Helper functions //
//////////////////////
func fillRow(row []*game.Block, index int) {
	for col := range row {
		row[col] = &game.Block{
			Row:      index,
			Col:      col,
			Inactive: true,
		}
	}
}

func testBoard() *game.Board {
	board := game.NewBoard()

	for i := 16; i < 22; i++ {
		fillRow(board.Canvas[i], i)
	}

	board.Canvas[20][3] = nil
	board.Canvas[18][0] = nil
	board.Canvas[18][3] = nil
	board.Canvas[17][0] = nil
	board.Canvas[17][7] = nil
	board.Canvas[17][8] = nil
	for i := 0; i < 10; i++ {
		if i == 4 || i == 5 {
			continue
		}
		board.Canvas[16][i] = nil
	}
	return board
}

func assert(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func printBoard(board *game.Board) {
	for _, row := range board.Canvas {
		for _, b := range row {
			if b == nil {
				fmt.Printf("_ ")
			} else {
				fmt.Printf("o ")
			}
		}
		fmt.Printf("\n")
	}
}
