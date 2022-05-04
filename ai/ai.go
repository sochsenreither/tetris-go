package ai

import (
	"github.com/sochsenreither/tetris-go/game"
)

const (
	heightWeight    float64 = 0.510066
	linesWeight     float64 = 0.760666
	holesWeight     float64 = 0.35663
	bumpinessWeight float64 = 0.184483
)

type TetrisPlayer struct{}

func (t *TetrisPlayer) NextMove(board *game.Board, activePiece *game.Piece) *game.Piece {
	board.RemovePiece(activePiece)

	bestScore := 0.0
	var best *game.Piece

	for rotation := 0; rotation < 4; rotation++ {
		workingPiece := activePiece.Clone()
		// Rotate piece
		for i := 0; i < rotation; i++ {
			workingPiece = workingPiece.Rotate()
		}

		// Move piece to the left
		workingPiece = moveLeft(board, workingPiece)

		collision := false
		for !collision {
			tmpPiece := workingPiece.Clone()

			// Move piece down as far as possible
			tmpPiece = moveDown(board, tmpPiece)

			// Calculate score
			score := calculateWeight(board, tmpPiece)

			if score > bestScore || bestScore == 0.0 {
				bestScore = score
				best = workingPiece
			}

			p := workingPiece.MoveRight()
			if board.Collision(p) {
				collision = true
			} else {
				workingPiece = p
			}
		}
	}

	board.DrawPiece(best)
	return best
}

func calculateWeight(board *game.Board, piece *game.Piece) float64 {
	board.DrawPiece(piece)
	score := -heightWeight*float64(aggregateHeight(board)) +
		linesWeight*float64(completeLines(board)) -
		holesWeight*float64(holes(board)) -
		bumpinessWeight*float64(bumpiness(board))
	board.RemovePiece(piece)
	return score
}

func moveDown(board *game.Board, piece *game.Piece) *game.Piece {
	collision := false
	for !collision {
		p := piece.MoveDown()
		if board.Collision(p) {
			collision = true
		} else {
			piece = p
		}
	}
	return piece
}

func moveLeft(board *game.Board, piece *game.Piece) *game.Piece {
	collision := false
	for !collision {
		p := piece.MoveLeft()
		if board.Collision(p) {
			collision = true
		} else {
			piece = p
		}
	}
	return piece
}

func aggregateHeight(board *game.Board) int {
	count := 0
	for i := 0; i < game.COLS; i++ {
		count += columnHeight(board, i)
	}
	return count
}

func completeLines(board *game.Board) int {
	count := 0
	for i := 0; i < game.ROWS; i++ {
		if board.CanClearLine(i) {
			count++
		}
	}
	return count
}

func holes(board *game.Board) int {
	count := 0
	for col := 0; col < game.COLS; col++ {
		b := false
		for row := 0; row < game.ROWS; row++ {
			if board.Canvas[row][col] != nil {
				b = true
			} else if board.Canvas[row][col] == nil && b {
				count++
			}
		}
	}
	return count
}

func bumpiness(board *game.Board) int {
	count := 0
	for col := 0; col < game.COLS-1; col++ {
		diff := columnHeight(board, col) - columnHeight(board, col+1)
		if diff < 0 {
			count += diff * -1
		} else {
			count += diff
		}
	}
	return count
}

func columnHeight(board *game.Board, col int) int {
	count := 0
	for _, row := range board.Canvas {
		if row[col] == nil {
			count++
		} else {
			return game.ROWS - count
		}
	}
	return game.ROWS - count
}
