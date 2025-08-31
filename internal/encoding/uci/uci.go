// Package uci implements encoding and decoding of UCI messages.
//
// A UCI message contains one or more tokens. A token is either a command, a
// parameter, or an argument.
//
// For example, the message "bestmove e2e4 ponder e7e5" contains four tokens:
//
//  1. "bestmove", a command
//  2. "e2e4", an argument to the "bestmove" command
//  3. "ponder", a parameter to the "bestmove" command
//  4. "e7e5",  an argument to the "ponder" parameter
//
// Leading and trailing whitespace is ignored when parsing or unmarshaling.
//
// TODO(clfs): Create custom error types.
package uci

import (
	"bytes"
	"encoding"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Message is the interface implemented by all messages.
type Message interface {
	encoding.TextAppender
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

// UCI represents a "uci" message.
type UCI struct{}

// AppendText implements the [encoding.TextAppender] interface.
func (m UCI) AppendText(b []byte) ([]byte, error) {
	return fmt.Append(b, "uci"), nil
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (m UCI) MarshalText() ([]byte, error) {
	return m.AppendText(nil)
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (m *UCI) UnmarshalText(text []byte) error {
	b := bytes.TrimSpace(text)
	if string(b) != "uci" {
		return errors.New("not a uci command")
	}
	return nil
}

// IsReady represents an "isready" message.
type IsReady struct{}

// AppendText implements the [encoding.TextAppender] interface.
func (m IsReady) AppendText(b []byte) ([]byte, error) {
	return fmt.Append(b, "isready"), nil
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (m IsReady) MarshalText() ([]byte, error) {
	return m.AppendText(nil)
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (m *IsReady) UnmarshalText(text []byte) error {
	b := bytes.TrimSpace(text)
	if string(b) != "isready" {
		return errors.New("not an isready command")
	}
	return nil
}

// SetOption represents a "setoption" message.
type SetOption struct {
	Name  string
	Value string
}

func (m SetOption) MarshalText() (text []byte, err error) {
	text = fmt.Append(text, "setoption")
	text = fmt.Appendf(text, " name %s", m.Name)
	if m.Value != "" {
		text = fmt.Appendf(text, " value %s", m.Value)
	}
	return
}

// UCINewGame represents a "ucinewgame" message.
type UCINewGame struct{}

func (m UCINewGame) MarshalText() ([]byte, error) {
	return []byte("ucinewgame"), nil
}

func (m *UCINewGame) UnmarshalText(text []byte) error {
	text = bytes.TrimSpace(text)

	if string(text) != "ucinewgame" {
		return errors.New("not a ucinewgame command")
	}

	return nil
}

// Position represents a "position" message.
type Position struct {
	Startpos bool
	FEN      string
	Moves    []string
}

func (m Position) MarshalText() (text []byte, err error) {
	text = fmt.Append(text, "position")
	if m.Startpos {
		text = fmt.Append(text, " startpos")
	}
	if m.FEN != "" {
		if m.Startpos {
			return nil, errors.New("cannot specify both startpos and fen")
		}
		text = fmt.Appendf(text, " fen %s", m.FEN)
	}
	if len(m.Moves) > 0 {
		text = fmt.Appendf(text, " moves %s", strings.Join(m.Moves, " "))
	}
	return
}

// Go represents a "go" message.
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

func (m *Go) MarshalText() (text []byte, err error) {
	text = fmt.Append(text, "go")
	if len(m.SearchMoves) > 0 {
		text = fmt.Appendf(text, " searchmoves %s", strings.Join(m.SearchMoves, " "))
	}
	if m.Ponder {
		text = fmt.Appendf(text, " ponder")
	}
	if m.WTime > 0 {
		text = fmt.Appendf(text, " wtime %d", m.WTime.Milliseconds())
	}
	if m.BTime > 0 {
		text = fmt.Appendf(text, " btime %d", m.BTime.Milliseconds())
	}
	if m.WInc > 0 {
		text = fmt.Appendf(text, " winc %d", m.WInc.Milliseconds())
	}
	if m.BInc > 0 {
		text = fmt.Appendf(text, " binc %d", m.BInc.Milliseconds())
	}
	if m.MovesToGo > 0 {
		text = fmt.Appendf(text, " movestogo %d", m.MovesToGo)
	}
	if m.Depth > 0 {
		text = fmt.Appendf(text, " depth %d", m.Depth)
	}
	if m.Nodes > 0 {
		text = fmt.Appendf(text, " nodes %d", m.Nodes)
	}
	if m.Mate > 0 {
		text = fmt.Appendf(text, " mate %d", m.Mate)
	}
	if m.MoveTime > 0 {
		text = fmt.Appendf(text, " movetime %d", m.MoveTime.Milliseconds())
	}
	if m.Infinite {
		text = fmt.Appendf(text, " infinite")
	}
	return
}

// Stop represents a "stop" message.
type Stop struct{}

func (m Stop) MarshalText() ([]byte, error) {
	return []byte("stop"), nil
}

// PonderHit represents a "ponderhit" message.
type PonderHit struct{}

func (m PonderHit) MarshalText() ([]byte, error) {
	return []byte("ponderhit"), nil
}

// Quit represents a "quit" message.
type Quit struct{}

func (m Quit) MarshalText() ([]byte, error) {
	return []byte("quit"), nil
}

// ID represents an "id" message.
type ID struct {
	Name   string
	Author string
}

func (m ID) MarshalText() ([]byte, error) {
	b := bytes.NewBufferString("id")

	if m.Name != "" {
		fmt.Fprintf(b, " name %s", m.Name)
	}

	if m.Author != "" {
		if m.Name != "" {
			return nil, errors.New("cannot specify both author and name")
		}
		fmt.Fprintf(b, " author %s", m.Author)
	}

	return b.Bytes(), nil
}

// UCIOK represents a "uciok" message.
type UCIOK struct{}

func (m UCIOK) MarshalText() ([]byte, error) {
	return []byte("uciok"), nil
}

// ReadyOK represents a "readyok" message.
type ReadyOK struct{}

func (m ReadyOK) MarshalText() ([]byte, error) {
	return []byte("readyok"), nil
}

// BestMove represents a "bestmove" command.
type BestMove struct {
	Move   string
	Ponder string
}

func (m BestMove) MarshalText() ([]byte, error) {
	b := bytes.NewBufferString("bestmove")

	fmt.Fprintf(b, " %s", m.Move)

	if m.Ponder != "" {
		fmt.Fprintf(b, " ponder %s", m.Ponder)
	}

	return b.Bytes(), nil
}

// Info represents an "info" message.
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

func (m Info) MarshalText() ([]byte, error) {
	b := bytes.NewBufferString("info")

	// TODO(clfs): Determine which fields are incompatible.

	if m.Depth > 0 {
		fmt.Fprintf(b, " depth %d", m.Depth)
	}

	if m.SelDepth > 0 {
		if !(m.Depth > 0) {
			return nil, errors.New("cannot specify seldepth without depth")
		}
	}

	if m.Time > 0 {
		fmt.Fprintf(b, " time %d", m.Time.Milliseconds())
	}

	if m.Nodes > 0 {
		fmt.Fprintf(b, " nodes %d", m.Nodes)
	}

	if len(m.PV) > 0 {
		fmt.Fprint(b, " pv")
		for _, m := range m.PV {
			fmt.Fprintf(b, " %s", m)
		}
	}

	if m.MultiPV > 0 {
		fmt.Fprintf(b, " multipv %d", m.MultiPV)
	}

	if m.ScoreCP {
		fmt.Fprintf(b, " score cp %d", m.Score)
	} else {
		fmt.Fprintf(b, " score mate %d", m.Score)
	}

	if m.CurrMove != "" {
		fmt.Fprintf(b, " currmove %s", m.CurrMove)
	}

	if m.CurrMoveNumber > 0 {
		fmt.Fprintf(b, " currmovenumber %d", m.CurrMoveNumber)
	}

	if m.NPS > 0 {
		fmt.Fprintf(b, " nps %d", m.NPS)
	}

	if m.TBHits > 0 {
		fmt.Fprintf(b, " tbhits %d", m.TBHits)
	}

	if m.Str != "" {
		fmt.Fprintf(b, " string %s", m.Str)
	}

	return b.Bytes(), nil
}

// Option represents an "option" message.
type Option struct {
	Name    string
	Type    string
	Default string
	Min     string
	Max     string
	Var     []string
}

// Blank is a placeholder that represents blank text.
type Blank struct{}

// AppendText implements the [encoding.TextAppender] interface.
func (m Blank) AppendText(b []byte) ([]byte, error) {
	return b, nil
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (m Blank) MarshalText() ([]byte, error) {
	return m.AppendText(nil)
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (m *Blank) UnmarshalText(text []byte) error {
	b := bytes.TrimSpace(text)
	if len(b) != 0 {
		return errors.New("uci: Blank.UnmarshalText: text not blank")
	}
	return nil
}

// Unknown is a placeholder that represents unknown text.
type Unknown struct {
	Text string
}

// AppendText implements the [encoding.TextAppender] interface.
func (m Unknown) AppendText(b []byte) ([]byte, error) {
	return fmt.Append(b, m.Text), nil
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (m Unknown) MarshalText() ([]byte, error) {
	return m.AppendText(nil)
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (m *Unknown) UnmarshalText(text []byte) error {
	m.Text = string(text)
	return nil
}
