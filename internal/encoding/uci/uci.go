// Package uci implements encoding and decoding of UCI commands.
//
// Not all commands are supported.
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

// Command is the interface implemented by all commands.
type Command interface {
	encoding.TextAppender
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

// Parse parses text and returns the corresponding command.
func Parse(text []byte) (Command, error) {
	var first []byte
	for field := range bytes.FieldsSeq(text) {
		first = field
		break
	}

	var cmd Command

	// TODO(clfs): Does this string conversion allocate? If so, can we avoid it?
	switch string(first) {
	case "uci":
		cmd = new(UCI)
	case "isready":
		cmd = new(IsReady)
	default:
		return nil, fmt.Errorf("uci.Parse: unknown command %s", first)
	}

	err := cmd.UnmarshalText(text)
	return cmd, err
}

// ParseString wraps [Parse].
func ParseString(s string) (any, error) {
	return Parse([]byte(s))
}

// UCI represents a "uci" command.
type UCI struct{}

// AppendText implements the [encoding.TextAppender] interface.
func (cmd UCI) AppendText(b []byte) ([]byte, error) {
	return fmt.Append(b, "uci"), nil
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (cmd UCI) MarshalText() ([]byte, error) {
	return cmd.AppendText(nil)
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (cmd *UCI) UnmarshalText(text []byte) error {
	b := bytes.TrimSpace(text)
	if string(b) != "uci" {
		return errors.New("not a uci command")
	}
	return nil
}

// IsReady represents an "isready" command.
type IsReady struct{}

// AppendText implements the [encoding.TextAppender] interface.
func (cmd IsReady) AppendText(b []byte) ([]byte, error) {
	return fmt.Append(b, "isready"), nil
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (cmd IsReady) MarshalText() ([]byte, error) {
	return cmd.AppendText(nil)
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (cmd *IsReady) UnmarshalText(text []byte) error {
	b := bytes.TrimSpace(text)
	if string(b) != "isready" {
		return errors.New("not an isready command")
	}
	return nil
}

// SetOption represents a "setoption" command.
type SetOption struct {
	Name  string
	Value string
}

func (cmd SetOption) MarshalText() (text []byte, err error) {
	text = fmt.Append(text, "setoption")
	text = fmt.Appendf(text, " name %s", cmd.Name)
	if cmd.Value != "" {
		text = fmt.Appendf(text, " value %s", cmd.Value)
	}
	return
}

// UCINewGame represents a "ucinewgame" command.
type UCINewGame struct{}

func (cmd UCINewGame) MarshalText() ([]byte, error) {
	return []byte("ucinewgame"), nil
}

func (cmd *UCINewGame) UnmarshalText(text []byte) error {
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

func (cmd Position) MarshalText() (text []byte, err error) {
	text = fmt.Append(text, "position")
	if cmd.Startpos {
		text = fmt.Append(text, " startpos")
	}
	if cmd.FEN != "" {
		if cmd.Startpos {
			return nil, errors.New("cannot specify both startpos and fen")
		}
		text = fmt.Appendf(text, " fen %s", cmd.FEN)
	}
	if len(cmd.Moves) > 0 {
		text = fmt.Appendf(text, " moves %s", strings.Join(cmd.Moves, " "))
	}
	return
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

func (cmd *Go) MarshalText() (text []byte, err error) {
	text = fmt.Append(text, "go")
	if len(cmd.SearchMoves) > 0 {
		text = fmt.Appendf(text, " searchmoves %s", strings.Join(cmd.SearchMoves, " "))
	}
	if cmd.Ponder {
		text = fmt.Appendf(text, " ponder")
	}
	if cmd.WTime > 0 {
		text = fmt.Appendf(text, " wtime %d", cmd.WTime.Milliseconds())
	}
	if cmd.BTime > 0 {
		text = fmt.Appendf(text, " btime %d", cmd.BTime.Milliseconds())
	}
	if cmd.WInc > 0 {
		text = fmt.Appendf(text, " winc %d", cmd.WInc.Milliseconds())
	}
	if cmd.BInc > 0 {
		text = fmt.Appendf(text, " binc %d", cmd.BInc.Milliseconds())
	}
	if cmd.MovesToGo > 0 {
		text = fmt.Appendf(text, " movestogo %d", cmd.MovesToGo)
	}
	if cmd.Depth > 0 {
		text = fmt.Appendf(text, " depth %d", cmd.Depth)
	}
	if cmd.Nodes > 0 {
		text = fmt.Appendf(text, " nodes %d", cmd.Nodes)
	}
	if cmd.Mate > 0 {
		text = fmt.Appendf(text, " mate %d", cmd.Mate)
	}
	if cmd.MoveTime > 0 {
		text = fmt.Appendf(text, " movetime %d", cmd.MoveTime.Milliseconds())
	}
	if cmd.Infinite {
		text = fmt.Appendf(text, " infinite")
	}
	return
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

// Blank is a placeholder that represents blank text.
type Blank struct{}

func (cmd Blank) AppendText(b []byte) ([]byte, error) {
	return b, nil
}

func (cmd Blank) MarshalText() ([]byte, error) {
	return cmd.AppendText(nil)
}

func (cmd *Blank) UnmarshalText(text []byte) error {
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
func (cmd Unknown) AppendText(b []byte) ([]byte, error) {
	return fmt.Append(b, cmd.Text), nil
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (cmd Unknown) MarshalText() ([]byte, error) {
	return cmd.AppendText(nil)
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (cmd *Unknown) UnmarshalText(text []byte) error {
	cmd.Text = string(text)
	return nil
}
