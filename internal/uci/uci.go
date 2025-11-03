// Package uci encodes and decodes UCI messages.
package uci

import (
	"errors"
	"fmt"
	"regexp"
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

// Quit represents a "quit" command.
type Quit struct{}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (m *Quit) UnmarshalText(text []byte) error {
	if string(text) != "quit" {
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

var (
	regexpIDName   = regexp.MustCompile(`^id name (.+)`)
	regexpIDAuthor = regexp.MustCompile(`^id author (.+)`)
)

// UnmarshalText implements [encoding.TextUnmarshaler].
func (m *ID) UnmarshalText(text []byte) error {
	subs := regexpIDName.FindSubmatch(text)
	if subs != nil {
		m.Name = string(subs[1])
		return nil
	}

	subs = regexpIDAuthor.FindSubmatch(text)
	if subs != nil {
		m.Author = string(subs[1])
		return nil
	}

	return errors.New("invalid id command")
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
