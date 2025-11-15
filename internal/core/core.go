// Package core implements core chess functionality.
package core

import (
	"fmt"
	"math/bits"
)

// A Bitboard stores one bit of information per board square.
type Bitboard uint64

// Count returns the number of set bits in b.
func (b *Bitboard) Count() int {
	return bits.OnesCount64(uint64(*b))
}

// IsEmpty returns true if no bits are set.
func (b *Bitboard) IsEmpty() bool {
	return *b == 0
}

// Get returns the bit at s.
func (b *Bitboard) Get(s Square) bool {
	v := *b & s.Bitboard()
	return v != 0
}

// Set sets the bit at s.
func (b *Bitboard) Set(s Square) {
	*b |= s.Bitboard()
}

// Clear clears the bit at s.
func (b *Bitboard) Clear(s Square) {
	*b &^= s.Bitboard()
}

// Color represents a color, like [White].
type Color bool

// [Color] constants.
const (
	White Color = false
	Black Color = true
)

// String implements [fmt.Stringer].
func (c Color) String() string {
	if c == White {
		return "White"
	}
	return "Black"
}

// Other returns the opposite color of c.
func (c Color) Other() Color {
	return Color(!c)
}

// PieceType represents a type of piece, like [Pawn].
type PieceType uint8

// [PieceType] constants.
const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

// String implements [fmt.Stringer].
func (p PieceType) String() string {
	switch p {
	case Pawn:
		return "Pawn"
	case Knight:
		return "Knight"
	case Bishop:
		return "Bishop"
	case Rook:
		return "Rook"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return fmt.Sprintf("PieceType(%d)", p)
	}
}

// Piece represents a piece.
//
// The zero value for Piece is a white pawn.
type Piece struct {
	Color     Color
	PieceType PieceType
}

// NewPiece returns a new [Piece].
func NewPiece(c Color, pt PieceType) Piece {
	return Piece{
		Color:     c,
		PieceType: pt,
	}
}

// File represents a file, like [FileA].
type File uint8

// [File] constants.
const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

// String implements [fmt.Stringer].
func (f File) String() string {
	if f > FileH {
		return fmt.Sprintf("File(%d)", f)
	}
	return fmt.Sprintf("File%c", 'A'+f)
}

// Bitboard returns a bitboard with just squares on f set.
func (f File) Bitboard() Bitboard {
	return Bitboard(0x101010101010101 << f)
}

// Rank represents a rank, like [Rank1].
type Rank uint8

// [Rank] constants.
const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

// String implements [fmt.Stringer].
func (r Rank) String() string {
	if r > Rank8 {
		return fmt.Sprintf("Rank(%d)", r)
	}
	return fmt.Sprintf("Rank%c", '1'+r)
}

// Bitboard returns a bitboard with just squares on r set.
func (r Rank) Bitboard() Bitboard {
	return Bitboard(0xff << (r * 8))
}

// Square represents a square, like [A1].
type Square uint8

// [Square] constants.
const (
	A1 Square = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

// NewSquare returns the square at the given coordinates.
func NewSquare(f File, r Rank) Square {
	return Square(r)*8 + Square(f)
}

// String implements [fmt.Stringer].
func (s Square) String() string {
	if s > H8 {
		return fmt.Sprintf("Square(%d)", s)
	}
	return fmt.Sprintf(
		"%c%c",
		'A'+s.File(),
		'1'+s.Rank(),
	)
}

// File returns the file that s is on.
func (s Square) File() File {
	return File(s % 8)
}

// Rank returns the rank that s is on.
func (s Square) Rank() Rank {
	return Rank(s / 8)
}

// Bitboard returns a bitboard with just s set.
func (s Square) Bitboard() Bitboard {
	return Bitboard(1 << s)
}

// Above returns the square above s, if any.
func (s Square) Above() (Square, bool) {
	if s.Rank() == Rank8 {
		return 0, false
	}
	return s + 8, true
}

// Below returns the square below s, if any.
func (s Square) Below() (Square, bool) {
	if s.Rank() == Rank1 {
		return 0, false
	}
	return s - 8, true
}

// Castling represents a set of castling rights.
//
// The zero value indicates neither player has castling rights.
//
// Note that a player may have a castling right even if they cannot legally make
// the corresponding move.
type Castling uint8

// [Castling] constants.
const (
	WhiteOO Castling = 1 << iota
	WhiteOOO
	BlackOO
	BlackOOO
)

// NewCastling returns a new [Castling] with all castling rights set.
func NewCastling() Castling {
	return WhiteOO | WhiteOOO | BlackOO | BlackOOO
}

// GetAll returns true if every castling right in x is also in c.
func (c *Castling) GetAll(x Castling) bool {
	return *c&x == x
}

// GetAny returns true if at least one castling right in x is also in c.
func (c *Castling) GetAny(x Castling) bool {
	return *c&x != 0
}

// Set sets in c all castling rights in x.
func (c *Castling) Set(x Castling) {
	*c |= x
}

// Clear clears from c all castling rights in x.
func (c *Castling) Clear(x Castling) {
	*c &^= x
}

// ClearColor clears from c both castling rights for the given color.
func (c *Castling) ClearColor(cl Color) {
	var lost Castling
	if cl == White {
		lost = WhiteOO | WhiteOOO
	} else {
		lost = BlackOO | BlackOOO
	}
	c.Clear(lost)
}

// EnPassant represents the right to capture en passant.
//
// The zero value indicates there is no right to capture en passant.
//
// Note that a player may have the right to capture en passant even if they
// cannot legally make the corresponding move.
type EnPassant uint8

// Square returns the square where the right to capture en passant exists, if
// any.
func (e *EnPassant) Square() (Square, bool) {
	return Square(*e), e.Exists()
}

// Exists returns true if the right to capture en passant exists.
func (e *EnPassant) Exists() bool {
	return *e != 0
}

// ExistsAt returns true if the right to capture en passant exists at s.
func (e *EnPassant) ExistsAt(s Square) bool {
	v, ok := e.Square()
	return ok && s == v
}

// Set sets the right to capture en passant at s.
func (e *EnPassant) Set(s Square) {
	*e = EnPassant(s)
}

// Clear clears the right to capture en passant.
func (e *EnPassant) Clear() {
	*e = 0
}

// Move represents a move.
type Move struct {
	// The moved piece, or king if castling, departs from this square.
	from Square

	// The moved piece, or king if castling, lands on this square.
	to Square

	// The moved piece promotes to this piece type.
	//
	// The zero value indicates no promotion occurs.
	promotion PieceType
}

// From returns the square the moved piece, or king if castling, departs from.
func (m Move) From() Square {
	return m.from
}

// To returns the square the moved piece, or king if castling, lands on.
func (m Move) To() Square {
	return m.to
}

// IsPromotion returns true if the move is a promotion move.
func (m Move) IsPromotion() bool {
	return m.promotion != Pawn
}

// PromotionTo returns the piece type that the moved piece promotes to, if any.
func (m Move) PromotionTo() (PieceType, bool) {
	return m.promotion, m.IsPromotion()
}
