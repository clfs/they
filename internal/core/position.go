package core

// Position describes a position.
type Position struct {
	// Squares occupied by white pieces.
	whiteBB Bitboard

	// Squares occupied by black pieces.
	blackBB Bitboard

	// Piece occupancy bitboards, indexed by [PieceType].
	pieceBBs [6]Bitboard

	// The player whose turn it is.
	turn Color

	// Castling rights.
	castling Castling

	// The right to capture en passant.
	enPassant EnPassant

	// The number of plies since the start of the game.
	plies uint16

	// The number of plies since the last capture, the last pawn advance, or the
	// start of the game, whichever is least.
	fiftyMoveRule uint8
}
