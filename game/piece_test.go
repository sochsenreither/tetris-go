package game

import (
	"reflect"
	"testing"
)

func TestRandomPiece(t *testing.T) {
	pieces := make(map[string]bool)
	for i := 0; i < 200; i++ {
		pieces[randomPiece().t] = true
	}
	if len(pieces) != 7 {
		t.Errorf("Expected to get every piece eventually, only got %d", len(pieces))
	}
}

func TestMovement(t *testing.T) {
	pieces := []*piece{
		iPiece(),
		lPiece(),
		jPiece(),
		oPiece(),
		sPiece(),
		tPiece(),
		zPiece(),
	}

	for _, p := range pieces {
		movements := map[string]func() *piece{
			"down":   p.moveDown,
			"left":   p.moveLeft,
			"right":  p.moveRight,
			"rotate": p.rotate,
		}

		for n, m := range movements {
			movedPiece := m()
			if p.t == "O" && n == "rotate" {
				if !reflect.DeepEqual(p.blocks, movedPiece.blocks) {
					t.Errorf("Expected position to stay the same")
				}
			} else {
				if reflect.DeepEqual(p.blocks, movedPiece.blocks) {
					t.Errorf("Expected position to change with func %s", n)
				}
			}

		}
	}
}
