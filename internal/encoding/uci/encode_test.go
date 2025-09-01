package uci

import (
	"bytes"
	"testing"
)

func TestEncoder(t *testing.T) {
	tests := []struct {
		in   []Message
		want string
	}{
		{
			[]Message{
				&UCI{},
			},
			"uci\n",
		},
		{
			[]Message{
				&IsReady{},
			},
			"isready\n",
		},
		{
			[]Message{
				&SetOption{
					Name: "foo",
				},
			},
			"setoption name foo\n",
		},
		{
			[]Message{
				&SetOption{
					Name:  "foo",
					Value: "bar baz",
				},
			},
			"setoption name foo value bar baz\n",
		},
	}

	for i, tt := range tests {
		var buf bytes.Buffer
		enc := NewEncoder(&buf)

		for j, m := range tt.in {
			if err := enc.WriteMessage(m); err != nil {
				t.Errorf("%d: message %d: error: %v", i, j, err)
			}
		}

		got := buf.String()
		if got != tt.want {
			t.Errorf("%d: output mismatch: got %q, want %q", i, got, tt.want)
		}
	}
}
