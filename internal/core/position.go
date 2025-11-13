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

func (p *Position) Move(m Move) {
	// TODO(clfs): Implement.
}
