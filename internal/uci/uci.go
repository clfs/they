// Package uci encodes and decodes UCI messages.
package uci

import (
	"bytes"
	"errors"
	"fmt"
)

// UCI represents a "uci" command.
type UCI struct{}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (m *UCI) UnmarshalText(text []byte) error {
	if string(text) != "uci" {
		return errors.New("not a uci command")
	}
	return nil
}

// AppendText implements [encoding.TextAppender].
func (m *UCI) AppendText(b []byte) ([]byte, error) {
	return fmt.Append(b, "uci"), nil
}

// IsReady represents an "isready" command.
type IsReady struct{}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (m *IsReady) UnmarshalText(text []byte) error {
	if string(text) != "isready" {
		return errors.New("not an isready command")
	}
	return nil
}

// AppendText implements [encoding.TextAppender].
func (m *IsReady) AppendText(b []byte) ([]byte, error) {
	return fmt.Append(b, "isready"), nil
}

// UCINewGame represents a "ucinewgame" command.
type UCINewGame struct{}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (m *UCINewGame) UnmarshalText(text []byte) error {
	if string(text) != "ucinewgame" {
		return errors.New("not a ucinewgame command")
	}
	return nil
}

// AppendText implements [encoding.TextAppender].
func (m *UCINewGame) AppendText(b []byte) ([]byte, error) {
	return fmt.Append(b, "ucinewgame"), nil
}

// Position represents a "position" command.
type Position struct {
	FEN      string
	Startpos bool
	Moves    []string
}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (m *Position) UnmarshalText(text []byte) error {
	fields := bytes.Fields(text)

	if len(fields) == 0 {
		return errors.New("no command provided")
	}
	if string(fields[0]) != "position" {
		return errors.New("not a position command")
	}

	return nil
}

// AppendText implements [encoding.TextAppender].
func (m *Position) AppendText(b []byte) ([]byte, error) {
	if m.Startpos && m.FEN != "" {
		return nil, errors.New("cannot specify both startpos and fen")
	}
	if !m.Startpos && m.FEN == "" {
		return nil, errors.New("must specify either startpos or fen")
	}

	b = fmt.Append(b, "position")
	if m.Startpos {
		b = fmt.Append(b, " startpos")
	}
	if m.FEN != "" {
		b = fmt.Appendf(b, " fen %s", m.FEN)
	}
	if len(m.Moves) > 0 {
		b = fmt.Append(b, " moves")
		for _, m := range m.Moves {
			b = fmt.Appendf(b, " %s", m)
		}
	}
	return b, nil
}

// Go represents a "go" command.
type Go struct{}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (m *Go) UnmarshalText(text []byte) error {
	fields := bytes.Fields(text)

	if len(fields) == 0 {
		return errors.New("no command provided")
	}
	if string(fields[0]) != "go" {
		return errors.New("not a go command")
	}

	return nil
}

// AppendText implements [encoding.TextAppender].
func (m *Go) AppendText(b []byte) ([]byte, error) {
	return fmt.Append(b, "go"), nil
}

// Stop represents a "stop" command.
type Stop struct{}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (m *Stop) UnmarshalText(text []byte) error {
	if string(text) != "stop" {
		return errors.New("not a stop command")
	}
	return nil
}

// AppendText implements [encoding.TextAppender].
func (m *Stop) AppendText(b []byte) ([]byte, error) {
	return fmt.Append(b, "stop"), nil
}

// Quit represents a "quit" command.
type Quit struct{}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (m *Quit) UnmarshalText(text []byte) error {
	if string(text) != "stop" {
		return errors.New("not a quit command")
	}
	return nil
}

// AppendText implements [encoding.TextAppender].
func (m *Quit) AppendText(b []byte) ([]byte, error) {
	return fmt.Append(b, "quit"), nil
}

// ID represents an "id" command.
type ID struct {
	Name   string
	Author string
}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (m *ID) UnmarshalText(text []byte) error {
	fields := bytes.Fields(text)

	if len(fields) == 0 {
		return errors.New("no command provided")
	}
	if string(fields[0]) != "id" {
		return errors.New("not an id command")
	}

	return nil
}

// AppendText implements [encoding.TextAppender].
func (m *ID) AppendText(b []byte) ([]byte, error) {
	if m.Name != "" && m.Author != "" {
		return nil, errors.New("cannot specify both name and author")
	}
	if m.Name == "" && m.Author == "" {
		return nil, errors.New("must specify either name or author")
	}

	b = fmt.Append(b, "id ")
	if m.Name != "" {
		b = fmt.Appendf(b, "name %s", m.Name)
	}
	if m.Author != "" {
		b = fmt.Appendf(b, "author %s", m.Author)
	}

	return b, nil
}

// UCIOk represents a "uciok" command.
type UCIOk struct{}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (m *UCIOk) UnmarshalText(text []byte) error {
	if string(text) != "uciok" {
		return errors.New("not a uciok command")
	}
	return nil
}

// AppendText implements [encoding.TextAppender].
func (m *UCIOk) AppendText(b []byte) ([]byte, error) {
	return fmt.Append(b, "uciok"), nil
}

// ReadyOk represents a "readyok" command.
type ReadyOk struct{}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (m *ReadyOk) UnmarshalText(text []byte) error {
	if string(text) != "readyok" {
		return errors.New("not a readyok command")
	}
	return nil
}

// AppendText implements [encoding.TextAppender].
func (m *ReadyOk) AppendText(b []byte) ([]byte, error) {
	return fmt.Append(b, "readyok"), nil
}

// BestMove represents a "bestmove" command.
type BestMove struct {
	Move string
}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (m *BestMove) UnmarshalText(text []byte) error {
	fields := bytes.Fields(text)

	if len(fields) == 0 {
		return errors.New("no command provided")
	}
	if string(fields[0]) != "bestmove" {
		return errors.New("not an bestmove command")
	}

	return nil
}

// AppendText implements [encoding.TextAppender].
func (m *BestMove) AppendText(b []byte) ([]byte, error) {
	if m.Move == "" {
		return nil, errors.New("no move provided")
	}

	b = fmt.Appendf(b, "bestmove %s", m.Move)

	return b, nil
}
