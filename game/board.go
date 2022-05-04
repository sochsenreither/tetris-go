package game

const (
	ROWS = 22
	COLS = 10
)

type Board struct {
	Canvas [][]*Block
}

func NewBoard() *Board {
	canvas := make([][]*Block, ROWS)
	for i := 0; i < ROWS; i++ {
		canvas[i] = make([]*Block, COLS)
	}
	return &Board{
		Canvas: canvas,
	}
}

func (b *Board) DrawPiece(p *Piece) {
	for _, block := range p.Blocks {
		b.Canvas[block.Row][block.Col] = block
	}
}

func (b *Board) RemovePiece(p *Piece) {
	for _, block := range p.Blocks {
		b.Canvas[block.Row][block.Col] = nil
	}
}

// Returns true if there is a Collision
func (b *Board) Collision(p *Piece) bool {
	for _, block := range p.Blocks {
		if block.Row < 0 || block.Row > ROWS-1 {
			return true
		}
		if block.Col < 0 || block.Col > COLS-1 {
			return true
		}
		if b.Canvas[block.Row][block.Col] != nil && b.Canvas[block.Row][block.Col].Inactive {
			return true
		}
	}
	return false
}

func (b *Board) clearLines() int {
	count := 0
	for i := ROWS - 1; i != 0; i-- {
		for b.CanClearLine(i) {
			count++
			b.moveCanvasDown(i)
		}
	}
	return count
}

func (b *Board) moveCanvasDown(index int) {
	for i := index; i != 0; i-- {
		for j := range b.Canvas[i] {
			if b.Canvas[i-1][j] == nil {
				b.Canvas[i][j] = nil
			} else {
				b.Canvas[i][j] = b.Canvas[i-1][j]
				b.Canvas[i][j].Row += 1
				b.Canvas[i-1][j] = nil
			}
		}
	}
}

func (b *Board) CanClearLine(index int) bool {
	if index <= 0 {
		return false
	}
	canClear := true
	for _, block := range b.Canvas[index] {
		if block == nil {
			canClear = false
			break
		}
	}
	return canClear
}
