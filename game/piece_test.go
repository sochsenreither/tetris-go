package game

import (
	"reflect"
	"testing"
)

func TestRandomBlock(t *testing.T) {
	pieces := make(map[string]bool)
	for i := 0; i < 200; i++ {
		pieces[randomPiece().t] = true
	}
	if len(pieces) != 7 {
		t.Errorf("Expected to get every piece eventually, only got %d", len(pieces))
	}
}

func TestMovement(t *testing.T) {
	p := randomPiece()
	movements := map[string]func() *piece{
		"down": p.moveDown,
		"left": p.moveLeft,
		"right": p.moveRight,
		//"rotate": p.rotate,
	}

	for n, m := range movements {
		movedPiece := m()
		if reflect.DeepEqual(p.blocks, movedPiece.blocks) {
			t.Errorf("Expected position to change with func %s", n)
		}
	}
}