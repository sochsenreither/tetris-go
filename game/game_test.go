package game

import "testing"

func TestGame(t *testing.T) {
	g := NewGame()
	g.SpawnPiece()
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
