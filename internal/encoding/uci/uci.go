// Package uci implements encoding and decoding of UCI commands.
//
// Not all commands are supported.
package uci

import (
	"time"

	"github.com/clfs/they/internal/encoding/fen"
	"github.com/clfs/they/internal/encoding/pacn"
)

// UCI represents a "uci" command.
type UCI struct{}

// IsReady represents an "isready" command.
type IsReady struct{}

// UCINewGame represents a "ucinewgame" command.
type UCINewGame struct{}

// PositionStartpos represents a "position startpos" command.
type PositionStartpos struct {
	Moves []pacn.Move
}

// PositionFEN represents a "position fen" command.
type PositionFEN struct {
	FEN   fen.Position
	Moves []pacn.Move
}

// Go represents a "go" command.
type Go struct {
	SearchMoves []pacn.Move
	Ponder      bool
	WTime       time.Duration
	BTime       time.Duration
	WInc        time.Duration
	BInc        time.Duration
	MovesToGo   int
	Depth       int
	Nodes       int
	Mate        int
	MoveTime    time.Duration
	Infinite    bool
}

// Stop represents a "stop" command.
type Stop struct{}

// PonderHit represents a "ponderhit" command.
type PonderHit struct{}

// Quit represents a "quit" command.
type Quit struct{}

// IDName represents an "id name" command.
type IDName struct {
	Name string
}

// IDAuthor represents an "id author" command.
type IDAuthor struct {
	Author string
}

// UCIOK represents a "uciok" command.
type UCIOK struct{}

// ReadyOK represents a "readyok" command.
type ReadyOK struct{}

// BestMove represents a "bestmove" command.
type BestMove struct {
	Move   pacn.Move
	Ponder pacn.Move
}

// Info represents an "info" command.
type Info struct {
	Depth          int
	SelDepth       int
	Time           time.Duration
	Nodes          int
	PV             []pacn.Move
	MultiPV        int
	ScoreCP        bool
	Score          int
	CurrMove       pacn.Move
	CurrMoveNumber int
	NPS            int
	TBHits         int
	String         string
}
