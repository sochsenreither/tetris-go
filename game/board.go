package game

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

// Returns true if there is a collision
func (b *board) collision(p *piece) bool {
	for _, block := range p.blocks {
		if block.row < 0 || block.row > ROWS-1 {
			return true
		}
		if block.col < 0 || block.col > COLS-1 {
			return true
		}
		if b.canvas[block.row][block.col] != nil && b.canvas[block.row][block.col].inactive {
			return true
		}
	}
	return false
}

func (b *board) clearLines() int {
	count := 0
	for i := ROWS - 1; i != 0; i-- {
		for b.canClearLine(i) {
			count++
			b.moveCanvasDown(i)
		}
	}
	return count
}

func (b *board) moveCanvasDown(index int) {
	for i := index; i != 0; i-- {
		for j := range b.canvas[i] {
			if b.canvas[i-1][j] == nil {
				b.canvas[i][j] = nil
			} else {
				b.canvas[i][j] = b.canvas[i-1][j].clone()
				b.canvas[i][j].row += 1
				b.canvas[i-1][j] = nil
			}
		}
	}
}

func (b *board) canClearLine(index int) bool {
	if index <= 0 {
		return false
	}
	canClear := true
	for _, block := range b.canvas[index] {
		if block == nil {
			canClear = false
			break
		}
	}
	return canClear
}
