package uci

import (
	"fmt"
	"io"
)

// Encoder is a streaming encoder for UCI messages.
type Encoder struct {
	w io.Writer
}

// NewEncoder constructs a new streaming encoder writing to w.
func NewEncoder(w io.Writer) *Encoder {
	e := Encoder{
		w: w,
	}
	return &e
}

// WriteMessage writes the next [Message].
func (e *Encoder) WriteMessage(m Message) error {
	// TODO(clfs): Is there a way to use encoding.TextAppender?
	text, err := m.MarshalText()
	if err != nil {
		return err
	}

	fmt.Fprintln(e.w, string(text))

	return nil
}
