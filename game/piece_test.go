package game

import "testing"

func TestRandomBlock(t *testing.T) {
	pieces := make(map[string]bool)
	for i := 0; i < 20; i++ {
		pieces[randomPiece().t] = true
	}
	if len(pieces) == 1 {
		t.Errorf("Expected to get every piece eventually, only got %d", len(pieces))
	}
}