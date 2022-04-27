package game

import (
	"fmt"
	"log"
)

type Game struct {
	board       *board
	activePiece *piece
}

func NewGame() *Game {
	return &Game{
		board:       newBoard(),
		activePiece: nil,
	}
}

func (g *Game) step(direction string) error {
	log.Printf("Step with direction %s", direction)
	// If there is no active piece spawn a new one
	if g.activePiece == nil {
		p := randomPiece()
		if !g.board.collisionCheck(p) {
			return fmt.Errorf("collision at start")
		}
		g.activePiece = p
	}

	g.board.removePiece(g.activePiece)

	switch direction {
		// TODO: implement faster speed when pressing down
	case "DOWN":
		// Move active piece one down
		// p := g.activePiece.moveDown()
		// if !g.board.collisionCheck(p) {
		// 	g.activePiece = p
		// } else {
		// 	// TODO: handle this case
		// 	// merge pieces and draw them
		// 	g.activePiece = nil
		// 	return nil
		// }
		p := g.activePiece.moveDown()
		if g.board.collisionCheck(p) {
			g.activePiece = p
		}
	case "UP":
		p := g.activePiece.rotate()
		if g.board.collisionCheck(p) {
			g.activePiece = p
		}
	case "LEFT":
		p := g.activePiece.moveLeft()
		if g.board.collisionCheck(p) {
			g.activePiece = p
		}
	case "RIGHT":
		p := g.activePiece.moveRight()
		if g.board.collisionCheck(p) {
			g.activePiece = p
		}
	case "":

	}

	g.board.drawPiece(g.activePiece)
	return nil
}
