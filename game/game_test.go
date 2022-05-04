package game

import "testing"

func TestGame(t *testing.T) {
	g := NewGame()
	g.Step("", false)
	if g.activePiece == nil {
		t.Errorf("Expected a piece to spawn")
	}
	if g.nextPiece == nil {
		t.Errorf("Expected the next piece to spawn")
	}
	g.clearedLines = 10
	g.Step("", false)
	if g.level == 0 {
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
	g.board.drawPiece(p)
	g.Step("", false)
	if !g.GameOVer() {
		t.Errorf("Expected game over")
	}
}
