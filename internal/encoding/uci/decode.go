package uci

import "bytes"

// Parse parses text and returns the corresponding message.
//
// If text is blank, Parse returns [Blank]. If text is an unrecognized message,
// Parse returns [Unknown].
func Parse(text []byte) (Message, error) {
	var first []byte
	for field := range bytes.FieldsSeq(text) {
		first = field
		break
	}

	var m Message

	// TODO(clfs): Does this string conversion allocate? If so, can we avoid it?
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

	err := m.UnmarshalText(text)
	return m, err
}

// ParseString wraps [Parse].
func ParseString(s string) (any, error) {
	return Parse([]byte(s))
}
