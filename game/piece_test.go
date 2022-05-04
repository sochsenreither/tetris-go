package game

import (
	"reflect"
	"testing"
)

func TestRandomPiece(t *testing.T) {
	pieces := make(map[string]bool)
	pf := newPieceFactory()
	for i := 0; i < 7; i++ {
		pieces[pf.randomPiece().t] = true
	}
	if len(pieces) != 7 {
		t.Errorf("Expected to get every piece within seven tries, only got %d", len(pieces))
	}
}

func TestMovement(t *testing.T) {
	pieces := []*Piece{
		iPiece(),
		lPiece(),
		jPiece(),
		oPiece(),
		sPiece(),
		tPiece(),
		zPiece(),
	}

	for _, p := range pieces {
		movements := map[string]func() *Piece{
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
