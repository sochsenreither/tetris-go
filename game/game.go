package game

import (
	"fmt"
)

type Game struct {
	board       *board
	activePiece *piece
	nextPiece   *piece
	score       uint
}

func NewGame() *Game {
	return &Game{
		board:       newBoard(),
		activePiece: nil,
		nextPiece:   nil,
		score:       0,
	}
}

func (g *Game) step(direction string, tick bool) error {
	// If there is no active piece spawn a new one
	if g.activePiece == nil {
		if g.nextPiece == nil {
			g.nextPiece = randomPiece()
		}
		if g.board.collision(g.nextPiece) {
			return fmt.Errorf("collision at start")
		}
		g.activePiece = g.nextPiece
		g.nextPiece = randomPiece()
	}

	var p *piece

	if tick {
		p = g.activePiece.moveDown()
		if g.board.collision(p) {
			g.board.drawPiece(g.activePiece)
			g.activePiece.setInactive()
			g.board.clearLines()
			g.activePiece = nil
			return nil
		}
		g.board.removePiece(g.activePiece)
		g.activePiece = p
	}

	switch direction {
	case "DOWN":
		p = g.activePiece.moveDown()
		if g.board.collision(p) {
			g.board.drawPiece(g.activePiece)
			g.activePiece.setInactive()
			g.board.clearLines()
			g.activePiece = nil
			return nil
		}
	case "UP":
		p = g.activePiece.rotate()
	case "LEFT":
		p = g.activePiece.moveLeft()
	case "RIGHT":
		p = g.activePiece.moveRight()
	default:
		g.board.drawPiece(g.activePiece)
		return nil
	}

	if !g.board.collision(p) {
		g.board.removePiece(g.activePiece)
		g.activePiece = p
		g.board.drawPiece(g.activePiece)
	}

	return nil
}
