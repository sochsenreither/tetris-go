package game

type Game struct {
	Board        *Board
	ActivePiece  *Piece
	NextPiece    *Piece
	pieceFactory *pieceFactory
	Score        uint
	Level        uint
	clearedLines uint
	Gameover     bool
}

func NewGame() *Game {
	return &Game{
		Board:        NewBoard(),
		ActivePiece:  nil,
		NextPiece:    nil,
		pieceFactory: newPieceFactory(),
		Score:        0,
		Level:        0,
		clearedLines: 0,
		Gameover:     false,
	}
}

func (g *Game) Step(direction string, tick bool) {
	// If there is no active piece spawn a new one
	if g.ActivePiece == nil {
		if g.NextPiece == nil {
			g.NextPiece = g.pieceFactory.randomPiece()
		}
		if g.Board.Collision(g.NextPiece) {
			g.Gameover = true
			return
		}
		g.ActivePiece = g.NextPiece
		g.NextPiece = g.pieceFactory.randomPiece()
	}

	if g.clearedLines >= 10 {
		g.Level++
		g.clearedLines = 0
	}

	var p *Piece

	if tick {
		p = g.ActivePiece.moveDown()
		if g.Board.Collision(p) {
			g.handleDroppedPiece()
			return
		}
		g.Board.RemovePiece(g.ActivePiece)
		g.ActivePiece = p
	}

	switch direction {
	case "DOWN":
		p = g.ActivePiece.moveDown()
		if g.Board.Collision(p) {
			g.handleDroppedPiece()
			return
		}
		g.Score += 1
	case "UP":
		p = g.ActivePiece.rotate()
	case "LEFT":
		p = g.ActivePiece.moveLeft()
	case "RIGHT":
		p = g.ActivePiece.moveRight()
	case "SPACE":
		for {
			p = g.ActivePiece.moveDown()
			if g.Board.Collision(p) {
				g.handleDroppedPiece()
				return
			}
			g.Board.RemovePiece(g.ActivePiece)
			g.ActivePiece = p
			g.Score += 2
		}
	default:
		g.Board.DrawPiece(g.ActivePiece)
		return
	}

	if !g.Board.Collision(p) {
		g.Board.RemovePiece(g.ActivePiece)
		g.ActivePiece = p
		g.Board.DrawPiece(g.ActivePiece)
	}
}

func (g *Game) handleDroppedPiece() {
	g.Board.DrawPiece(g.ActivePiece)
	g.ActivePiece.setInactive()
	count := g.Board.clearLines()
	g.ActivePiece = nil

	points := []uint{0, 40, 100, 300, 1200}

	g.Score += points[count] * (g.Level + 1)
	g.clearedLines += uint(count)
}
