package core

// Board represents the positions of pieces on a board.
type Board struct {
	// Squares occupied by white pieces.
	white Bitboard

	// Squares occupied by black pieces.
	black Bitboard

	// Piece occupancy bitboards, indexed by [PieceType].
	pieces [6]Bitboard
}

// NewBoard returns a new [Board] with pieces in the starting position.
func NewBoard() Board {
	var b Board

	b.Set(Piece{Color: White, PieceType: Rook}, A1)
	b.Set(Piece{Color: White, PieceType: Knight}, B1)
	b.Set(Piece{Color: White, PieceType: Bishop}, C1)
	b.Set(Piece{Color: White, PieceType: Queen}, D1)
	b.Set(Piece{Color: White, PieceType: King}, E1)
	b.Set(Piece{Color: White, PieceType: Bishop}, F1)
	b.Set(Piece{Color: White, PieceType: Knight}, G1)
	b.Set(Piece{Color: White, PieceType: Rook}, H1)
	for f := FileA; f <= FileH; f++ {
		s := NewSquare(f, Rank2)
		b.Set(Piece{Color: White, PieceType: Pawn}, s)
	}

	b.Set(Piece{Color: Black, PieceType: Rook}, A8)
	b.Set(Piece{Color: Black, PieceType: Knight}, B8)
	b.Set(Piece{Color: Black, PieceType: Bishop}, C8)
	b.Set(Piece{Color: Black, PieceType: Queen}, D8)
	b.Set(Piece{Color: Black, PieceType: King}, E8)
	b.Set(Piece{Color: Black, PieceType: Bishop}, F8)
	b.Set(Piece{Color: Black, PieceType: Knight}, G8)
	b.Set(Piece{Color: Black, PieceType: Rook}, H8)
	for f := FileA; f <= FileH; f++ {
		s := NewSquare(f, Rank7)
		b.Set(Piece{Color: Black, PieceType: Pawn}, s)
	}

	return b
}

// White returns all squares occupied by white pieces.
func (b *Board) White() Bitboard {
	return b.white
}

// Black returns all squares occupied by black pieces.
func (b *Board) Black() Bitboard {
	return b.black
}

// PieceColor returns the color of the piece on s, if any.
func (b *Board) PieceColor(s Square) (Color, bool) {
	switch {
	case b.white.Get(s):
		return White, true
	case b.black.Get(s):
		return Black, true
	default:
		// [White] is the zero value for [Color].
		return White, false
	}
}

// PieceType returns the type of the piece on s, if any.
func (b *Board) PieceType(s Square) (PieceType, bool) {
	for i := range b.pieces {
		if b.pieces[i].Get(s) {
			return PieceType(i), true
		}
	}
	return 0, false
}

// Piece returns the piece on a square, if any.
func (b *Board) Piece(s Square) (Piece, bool) {
	c, ok := b.PieceColor(s)
	if !ok {
		return Piece{}, false
	}
	pt, _ := b.PieceType(s)
	p := Piece{
		Color:     c,
		PieceType: pt,
	}
	return p, true
}

// Set sets a piece on a square.
//
// If another piece is already on s, Set replaces it with p.
func (b *Board) Set(p Piece, s Square) {
	b.Clear(s)
	if p.Color == White {
		b.white.Set(s)
	} else {
		b.black.Set(s)
	}
	b.pieces[p.PieceType].Set(s)
}

// Move moves a piece between two squares.
func (b *Board) Move(p Piece, from, to Square) {
	b.Clear(from)
	b.Set(p, to)
}

// Clear clears a piece from a square.
//
// If the square is already empty, nothing happens.
func (b *Board) Clear(s Square) {
	b.white.Clear(s)
	b.black.Clear(s)
	for i := range b.pieces {
		b.pieces[i].Clear(s)
	}
}

// IsOccupied returns true if the given square is occupied.
func (b *Board) IsOccupied(s Square) bool {
	return b.white.Get(s) || b.black.Get(s)
}
