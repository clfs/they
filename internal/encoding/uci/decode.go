package uci

import (
	"bufio"
	"bytes"
	"io"
)

// Decoder is a streaming decoder for UCI messages.
type Decoder struct {
	s *bufio.Scanner
}

// NewDecoder constructs a new streaming decoder reading from r.
func NewDecoder(r io.Reader) *Decoder {
	s := bufio.NewScanner(r)
	d := Decoder{
		s: s,
	}
	return &d
}

// ReadMessage reads the next [Message]. It returns [io.EOF] if there are no
// more messages.
func (d *Decoder) ReadMessage() (Message, error) {
	if ok := d.s.Scan(); !ok {
		err := d.s.Err()
		if err == nil {
			err = io.EOF
		}
		return nil, err
	}

	line := d.s.Bytes()

	var first []byte
	for field := range bytes.FieldsSeq(line) {
		first = field
		break
	}

	var m Message

	switch string(first) {
	case "uci":
		m = new(UCI)
	case "isready":
		m = new(IsReady)
	case "":
		m = new(Blank)
	default:
		m = new(Unknown)
	}

	if err := m.UnmarshalText(line); err != nil {
		return nil, err
	}
	return m, nil
}
