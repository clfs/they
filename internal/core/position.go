package core

// Position describes a position.
type Position struct {
	// TODO(clfs): Make these fields unexported.

	// The board.
	Board Board

	// The player whose turn it is.
	Turn Color

	// Castling rights.
	Castling Castling

	// The right to capture en passant.
	EnPassant EnPassant

	// The number of plies since the start of the game.
	Plies uint16

	// A ply counter used to address the 50-move rule (FIDE Laws of Chess §9.3)
	// and 75-move rule (§9.6.2).
	//
	// If a ply is a capture or pawn move, the counter is reset to zero.
	// Otherwise, the counter is incremented.
	HalfmoveClock uint8
}

// NewPosition returns the starting position.
func NewPosition() Position {
	return Position{
		Board:    NewBoard(),
		Castling: NewCastling(),
	}
}

// Move makes a move.
//
// It does not check for invalid moves.
func (p *Position) Move(m Move) {
	// Find the move's from and to squares.
	from, to := m.From(), m.To()

	// Select the piece to move. For castling moves, this is the king.
	heldPiece, _ := p.Board.Piece(from)

	// Is the move a pawn move?
	isPawnMove := heldPiece.PieceType == Pawn

	// Is the move an en passant capture?
	isEnPassantCapture := isPawnMove && p.EnPassant.ExistsAt(to)

	// Is the move a regular capture, i.e., not en passant?
	isRegularCapture := p.Board.IsOccupied(to)

	// Is the move a capture?
	isCapture := isRegularCapture || isEnPassantCapture

	// Is it White's turn?
	isWhiteTurn := p.Turn == White

	// If the move is an en passant capture, remove the captured pawn.
	if isEnPassantCapture {
		s, _ := p.EnPassant.Square()
		if isWhiteTurn {
			s, _ = s.Down()
		} else {
			s, _ = s.Up()
		}
		p.Board.Clear(s)
	}

	// Is the move a king move? This includes castling moves.
	isKingMove := heldPiece.PieceType == King

	// If the held piece is a king, then the player making the move loses both
	// of their castling rights.
	if isKingMove {
		p.Castling.ClearColor(p.Turn)
	}

	// If the held piece leaves a corner square, then the castling right that
	// involves that square is lost.
	switch from {
	case A1:
		p.Castling.Clear(WhiteOOO)
	case H1:
		p.Castling.Clear(WhiteOO)
	case A8:
		p.Castling.Clear(BlackOOO)
	case H8:
		p.Castling.Clear(BlackOO)
	}

	// If the held piece lands on a corner square, then the castling right that
	// involves that square is lost.
	switch to {
	case A1:
		p.Castling.Clear(WhiteOOO)
	case H1:
		p.Castling.Clear(WhiteOO)
	case A8:
		p.Castling.Clear(BlackOOO)
	case H8:
		p.Castling.Clear(BlackOO)
	}

	// Is the move a double pawn push?
	fromRank, toRank := from.Rank(), to.Rank()
	isDoublePawnPush := isPawnMove &&
		((fromRank == Rank2 && toRank == Rank4) || (fromRank == Rank7 && toRank == Rank5))

	// Update the right to capture en passant.
	if isDoublePawnPush {
		var s Square
		if isWhiteTurn {
			s, _ = from.Up()
		} else {
			s, _ = from.Down()
		}
		p.EnPassant.Set(s)
	} else {
		p.EnPassant.Clear()
	}

	// Is the move a castling move?
	fromFile, toFile := from.File(), to.File()
	isCastlingMove := isKingMove && (fromFile == FileE) && (toFile == FileG || toFile == FileC)

	// If the move is a castling move, move the castling rook.
	if isCastlingMove {
		var rookFrom, rookTo Square
		switch {
		case from == E1 && to == G1: // WhiteOO
			rookFrom, rookTo = H1, F1
		case from == E1 && to == C1: // WhiteOOO
			rookFrom, rookTo = A1, D1
		case from == E8 && to == G8: // BlackOO
			rookFrom, rookTo = H8, F8
		case from == E8 && to == C8: // BlackOOO
			rookFrom, rookTo = A8, D8
		}
		rook := NewPiece(p.Turn, Rook)
		p.Board.Move(rook, rookFrom, rookTo)
	}

	// Move the held piece. If castling, this is the king.
	p.Board.Move(heldPiece, from, to)

	// Is the move a promotion?
	isPromotion := m.IsPromotion()

	// If the move is a promotion, replace the promoted pawn.
	if isPromotion {
		pt, _ := m.PromotionTo()
		piece := NewPiece(p.Turn, pt)
		p.Board.Set(piece, to)
	}

	// Increment the ply count.
	p.Plies++

	// Update the halfmove clock.
	if isCapture || isPawnMove {
		p.HalfmoveClock = 0
	} else {
		p.HalfmoveClock++
	}

	// Finish the turn.
	p.Turn = p.Turn.Other()
}

// Moves returns all likely legal moves in the position.
//
// If there are no legal moves in the position, Moves returns nil.
//
// Moves does not account for:
//   - Dead positions (FIDE Laws of Chess §5.2.2)
//   - Threefold repetition (§9.2.2)
//   - Fivefold repetition (§9.6.1)
func (p *Position) Moves() []Move {
	// If the halfmove clock is 75 or greater, there are no legal moves. The
	// cutoff is 75, rather than 50, since the 50-move rule (FIDE Laws of Chess
	// §9.3) involves an optional claim and the 75-move rule (§9.6.2) does not.
	//
	// TODO(clfs): Reword.
	if p.HalfmoveClock >= 75 {
		return nil
	}

	return nil
}
