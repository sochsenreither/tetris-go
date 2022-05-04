package game

import "testing"

func TestGame(t *testing.T) {
	g := NewGame()
	g.Step("", false)
	if g.ActivePiece == nil {
		t.Errorf("Expected a piece to spawn")
	}
	if g.NextPiece == nil {
		t.Errorf("Expected the next piece to spawn")
	}
	g.clearedLines = 10
	g.Step("", false)
	if g.Level == 0 {
		t.Errorf("Expected the level to change")
	}
	g.Step("UP", false)
	g.Step("DOWN", false)
	g.Step("LEFT", false)
	g.Step("RIGHT", false)
}

func TestGameOver(t *testing.T) {
	g := NewGame()
	p := g.pieceFactory.randomPiece()
	p.setInactive()
	g.Board.DrawPiece(p)
	g.Step("", false)
	if !g.Gameover {
		t.Errorf("Expected game over")
	}
}
