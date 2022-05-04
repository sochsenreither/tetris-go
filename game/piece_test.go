package game

import (
	"reflect"
	"testing"
)

func TestRandomPiece(t *testing.T) {
	pieces := make(map[string]bool)
	pf := newPieceFactory()
	for i := 0; i < 7; i++ {
		pieces[pf.randomPiece().T] = true
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
			"down":   p.MoveDown,
			"left":   p.MoveLeft,
			"right":  p.MoveRight,
			"rotate": p.Rotate,
		}

		for n, m := range movements {
			movedPiece := m()
			if p.T == "O" && n == "rotate" {
				if !reflect.DeepEqual(p.Blocks, movedPiece.Blocks) {
					t.Errorf("Expected position to stay the same")
				}
			} else {
				if reflect.DeepEqual(p.Blocks, movedPiece.Blocks) {
					t.Errorf("Expected position to change with func %s", n)
				}
			}

		}
	}
}

func TestClone(t *testing.T) {
	pf := newPieceFactory()
	p := pf.randomPiece()
	pCloned := p.Clone()

	if &p.T == &pCloned.T {
		t.Errorf("Same memory address")
	}

	for i, block := range p.Blocks {
		if block == pCloned.Blocks[i] {
			t.Errorf("Same memory address")
		}
	}
}
