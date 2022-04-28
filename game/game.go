package game

type game struct {
	board        *board
	activePiece  *piece
	nextPiece    *piece
	score        uint
	level        uint
	clearedLines uint
	gameover     bool
}

func NewGame() *game {
	return &game{
		board:        newBoard(),
		activePiece:  nil,
		nextPiece:    nil,
		score:        0,
		level:        0,
		clearedLines: 0,
		gameover:     false,
	}
}

func (g *game) step(direction string, tick bool) {
	// If there is no active piece spawn a new one
	if g.activePiece == nil {
		if g.nextPiece == nil {
			g.nextPiece = randomPiece()
		}
		if g.board.collision(g.nextPiece) {
			g.gameover = true
			return
		}
		g.activePiece = g.nextPiece
		g.nextPiece = randomPiece()
	}

	if g.clearedLines >= 10 {
		g.level++
		g.clearedLines -= 2
	}

	var p *piece

	if tick {
		p = g.activePiece.moveDown()
		if g.board.collision(p) {
			g.handleDroppedPiece()
			return
		}
		g.board.removePiece(g.activePiece)
		g.activePiece = p
	}

	switch direction {
	case "DOWN":
		p = g.activePiece.moveDown()
		if g.board.collision(p) {
			g.handleDroppedPiece()
			return
		}
		g.score += 1
	case "UP":
		p = g.activePiece.rotate()
	case "LEFT":
		p = g.activePiece.moveLeft()
	case "RIGHT":
		p = g.activePiece.moveRight()
	default:
		g.board.drawPiece(g.activePiece)
		return
	}

	if !g.board.collision(p) {
		g.board.removePiece(g.activePiece)
		g.activePiece = p
		g.board.drawPiece(g.activePiece)
	}
}

func (g *game) handleDroppedPiece() {
	g.board.drawPiece(g.activePiece)
	g.activePiece.setInactive()
	count := g.board.clearLines()
	g.activePiece = nil

	points := []uint{0, 40, 100, 300, 1200}

	g.score += points[count] * (g.level + 1)
	g.clearedLines += uint(count)
}
