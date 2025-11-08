// Package core implements core chess functionality.
package core

import "fmt"

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
