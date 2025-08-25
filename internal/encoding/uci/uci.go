// Package uci implements encoding and decoding of UCI commands.
//
// Not all commands are supported.
//
// TODO(clfs): Create MarshalError and UnmarshalError.
package uci

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

// UCI represents a "uci" command.
type UCI struct{}

func (cmd UCI) MarshalText() ([]byte, error) {
	return []byte("uci"), nil
}

func (cmd UCI) UnmarshalText(text []byte) error {
	text = bytes.TrimSpace(text)

	if string(text) != "uci" {
		return errors.New("not a uci command")
	}

	return nil
}

// IsReady represents an "isready" command.
type IsReady struct{}

func (cmd IsReady) MarshalText() ([]byte, error) {
	return []byte("isready"), nil
}

func (cmd IsReady) UnmarshalText(text []byte) error {
	text = bytes.TrimSpace(text)

	if string(text) != "isready" {
		return errors.New("not an isready command")
	}

	return nil
}

// SetOption represents a "setoption" command.
type SetOption struct {
	Name  string
	Value string
}

func (cmd SetOption) MarshalText() ([]byte, error) {
	b := bytes.NewBufferString("setoption")

	fmt.Fprintf(b, " name %s", cmd.Name)

	if cmd.Value != "" {
		fmt.Fprintf(b, " value %s", cmd.Value)
	}

	return b.Bytes(), nil
}

// UCINewGame represents a "ucinewgame" command.
type UCINewGame struct{}

func (cmd UCINewGame) MarshalText() ([]byte, error) {
	return []byte("ucinewgame"), nil
}

func (cmd UCINewGame) UnmarshalText(text []byte) error {
	text = bytes.TrimSpace(text)

	if string(text) != "ucinewgame" {
		return errors.New("not a ucinewgame command")
	}

	return nil
}

// Position represents a "position" command.
type Position struct {
	Startpos bool
	FEN      string
	Moves    []string
}

func (cmd Position) MarshalText() ([]byte, error) {
	b := bytes.NewBufferString("position")

	if cmd.Startpos {
		fmt.Fprint(b, " startpos")
	}

	if cmd.FEN != "" {
		if cmd.Startpos {
			return nil, errors.New("cannot specify both startpos and fen")
		}
		fmt.Fprintf(b, " fen %s", cmd.FEN)
	}

	if len(cmd.Moves) > 0 {
		fmt.Fprint(b, " moves")
		for _, m := range cmd.Moves {
			fmt.Fprintf(b, " %s", m)
		}
	}

	return b.Bytes(), nil
}

// Go represents a "go" command.
type Go struct {
	SearchMoves []string
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

func (cmd *Go) MarshalText() ([]byte, error) {
	b := bytes.NewBufferString("go")

	// TODO(clfs): Determine which fields are incompatible.

	if len(cmd.SearchMoves) > 0 {
		fmt.Fprint(b, " searchmoves")
		for _, m := range cmd.SearchMoves {
			fmt.Fprintf(b, " %s", m)
		}
	}

	if cmd.Ponder {
		fmt.Fprint(b, " ponder")
	}

	if cmd.WTime > 0 {
		fmt.Fprintf(b, " wtime %d", cmd.WTime.Milliseconds())
	}

	if cmd.BTime > 0 {
		fmt.Fprintf(b, " btime %d", cmd.BTime.Milliseconds())
	}

	if cmd.WInc > 0 {
		fmt.Fprintf(b, " winc %d", cmd.WInc.Milliseconds())
	}

	if cmd.BInc > 0 {
		fmt.Fprintf(b, " binc %d", cmd.BInc.Milliseconds())
	}

	if cmd.MovesToGo > 0 {
		fmt.Fprintf(b, " movestogo %d", cmd.MovesToGo)
	}

	if cmd.Depth > 0 {
		fmt.Fprintf(b, " depth %d", cmd.Depth)
	}

	if cmd.Nodes > 0 {
		fmt.Fprintf(b, " nodes %d", cmd.Nodes)
	}

	if cmd.Mate > 0 {
		fmt.Fprintf(b, " mate %d", cmd.Mate)
	}

	if cmd.MoveTime > 0 {
		fmt.Fprintf(b, " movetime %d", cmd.MoveTime.Milliseconds())
	}

	if cmd.Infinite {
		fmt.Fprint(b, " infinite")
	}

	return b.Bytes(), nil
}

// Stop represents a "stop" command.
type Stop struct{}

func (cmd Stop) MarshalText() ([]byte, error) {
	return []byte("stop"), nil
}

// PonderHit represents a "ponderhit" command.
type PonderHit struct{}

func (cmd PonderHit) MarshalText() ([]byte, error) {
	return []byte("ponderhit"), nil
}

// Quit represents a "quit" command.
type Quit struct{}

func (cmd Quit) MarshalText() ([]byte, error) {
	return []byte("quit"), nil
}

// ID represents an "id" command.
type ID struct {
	Name   string
	Author string
}

func (cmd ID) MarshalText() ([]byte, error) {
	b := bytes.NewBufferString("id")

	if cmd.Name != "" {
		fmt.Fprintf(b, " name %s", cmd.Name)
	}

	if cmd.Author != "" {
		if cmd.Name != "" {
			return nil, errors.New("cannot specify both author and name")
		}
		fmt.Fprintf(b, " author %s", cmd.Author)
	}

	return b.Bytes(), nil
}

// UCIOK represents a "uciok" command.
type UCIOK struct{}

func (cmd UCIOK) MarshalText() ([]byte, error) {
	return []byte("uciok"), nil
}

// ReadyOK represents a "readyok" command.
type ReadyOK struct{}

func (cmd ReadyOK) MarshalText() ([]byte, error) {
	return []byte("readyok"), nil
}

// BestMove represents a "bestmove" command.
type BestMove struct {
	Move   string
	Ponder string
}

func (cmd BestMove) MarshalText() ([]byte, error) {
	b := bytes.NewBufferString("bestmove")

	fmt.Fprintf(b, " %s", cmd.Move)

	if cmd.Ponder != "" {
		fmt.Fprintf(b, " ponder %s", cmd.Ponder)
	}

	return b.Bytes(), nil
}

// Info represents an "info" command.
type Info struct {
	Depth          int
	SelDepth       int
	Time           time.Duration
	Nodes          int
	PV             []string
	MultiPV        int
	ScoreCP        bool
	Score          int
	CurrMove       string
	CurrMoveNumber int
	NPS            int
	TBHits         int
	Str            string
}

func (cmd Info) MarshalText() ([]byte, error) {
	b := bytes.NewBufferString("info")

	// TODO(clfs): Determine which fields are incompatible.

	if cmd.Depth > 0 {
		fmt.Fprintf(b, " depth %d", cmd.Depth)
	}

	if cmd.SelDepth > 0 {
		if !(cmd.Depth > 0) {
			return nil, errors.New("cannot specify seldepth without depth")
		}
	}

	if cmd.Time > 0 {
		fmt.Fprintf(b, " time %d", cmd.Time.Milliseconds())
	}

	if cmd.Nodes > 0 {
		fmt.Fprintf(b, " nodes %d", cmd.Nodes)
	}

	if len(cmd.PV) > 0 {
		fmt.Fprint(b, " pv")
		for _, m := range cmd.PV {
			fmt.Fprintf(b, " %s", m)
		}
	}

	if cmd.MultiPV > 0 {
		fmt.Fprintf(b, " multipv %d", cmd.MultiPV)
	}

	if cmd.ScoreCP {
		fmt.Fprintf(b, " score cp %d", cmd.Score)
	} else {
		fmt.Fprintf(b, " score mate %d", cmd.Score)
	}

	if cmd.CurrMove != "" {
		fmt.Fprintf(b, " currmove %s", cmd.CurrMove)
	}

	if cmd.CurrMoveNumber > 0 {
		fmt.Fprintf(b, " currmovenumber %d", cmd.CurrMoveNumber)
	}

	if cmd.NPS > 0 {
		fmt.Fprintf(b, " nps %d", cmd.NPS)
	}

	if cmd.TBHits > 0 {
		fmt.Fprintf(b, " tbhits %d", cmd.TBHits)
	}

	if cmd.Str != "" {
		fmt.Fprintf(b, " string %s", cmd.Str)
	}

	return b.Bytes(), nil
}

// Option represents an "option" command.
type Option struct {
	Name    string
	Type    string
	Default string
	Min     string
	Max     string
	Var     []string
}
