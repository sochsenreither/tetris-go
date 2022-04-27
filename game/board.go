package game

const ROWS = 22
const COLS = 10

type board struct {
	canvas [][]*block
}

func newBoard() *board {
	canvas := make([][]*block, ROWS)
	for i := 0; i < ROWS; i++ {
		canvas[i] = make([]*block, COLS)
	}
	return &board{
		canvas: canvas,
	}
}

func (b *board) drawPiece(p *piece) {
	for _, block := range p.blocks {
		b.canvas[block.row][block.col] = block
	}
}

func (b *board) removePiece(p *piece) {
	for _, block := range p.blocks {
		b.canvas[block.row][block.col] = nil
	}
}

// Returns false if there is a collision
func (b *board) collisionCheck(p *piece) bool {
	for _, block := range p.blocks {
		if block.row < 0 || block.row > ROWS-1 {
			return false
		}
		if block.col < 0 || block.col > COLS-1 {
			return false
		}
		if b.canvas[block.row][block.col] != nil {
			return false
		}
	}
	return true
}

func (b *board) clearLine() {

}
