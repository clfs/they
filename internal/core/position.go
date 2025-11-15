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

	// The number of plies since the most recent capture or pawn advance. If no
	// captures or pawn advances have occurred, this is the number of plies
	// since the start of the game.
	FiftyMoveRule uint8
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
			s, _ = s.Below()
		} else {
			s, _ = s.Above()
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
		(fromRank == Rank2 && toRank == Rank4) || (fromRank == Rank7 && toRank == Rank5)

	// If the move is a double pawn push, update the right to capture en
	// passant.
	if isDoublePawnPush {
		var s Square
		if isWhiteTurn {
			s, _ = from.Above()
		} else {
			s, _ = from.Below()
		}
		p.EnPassant.Set(s)
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

	// Update the fifty move rule counter.
	if isCapture || isPawnMove {
		p.FiftyMoveRule++
	} else {
		p.FiftyMoveRule = 0
	}

	// Finish the turn.
	p.Turn = p.Turn.Other()
}
