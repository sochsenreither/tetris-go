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

// Returns true if a piece was dropped
func (g *Game) Step(direction string, tick bool) bool {
	if g.clearedLines >= 10 {
		g.Level++
		g.clearedLines = 0
	}

	var p *Piece

	if tick {
		p = g.ActivePiece.MoveDown()
		if g.Board.Collision(p) {
			g.handleDroppedPiece()
			return true
		}
		g.Board.RemovePiece(g.ActivePiece)
		g.ActivePiece = p
	}

	switch direction {
	case "DOWN":
		p = g.ActivePiece.MoveDown()
		if g.Board.Collision(p) {
			g.handleDroppedPiece()
			return true
		}
		g.Score += 1
	case "UP":
		p = g.ActivePiece.Rotate()
	case "LEFT":
		p = g.ActivePiece.MoveLeft()
	case "RIGHT":
		p = g.ActivePiece.MoveRight()
	case "SPACE":
		for {
			p = g.ActivePiece.MoveDown()
			if g.Board.Collision(p) {
				g.handleDroppedPiece()
				return true
			}
			g.Board.RemovePiece(g.ActivePiece)
			g.ActivePiece = p
			g.Score += 2
		}
	default:
		g.Board.DrawPiece(g.ActivePiece)
		return false
	}

	if !g.Board.Collision(p) {
		g.Board.RemovePiece(g.ActivePiece)
		g.ActivePiece = p
		g.Board.DrawPiece(g.ActivePiece)
	}
	return false
}

func (g *Game) SpawnPiece() {
	if g.ActivePiece == nil {
		if g.NextPiece == nil {
			g.NextPiece = g.pieceFactory.randomPiece()
		}
		if g.Board.Collision(g.NextPiece) {
			g.Gameover = true
		}
		g.ActivePiece = g.NextPiece
		g.NextPiece = g.pieceFactory.randomPiece()
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
