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
			t.Errorf("%d: output mismatch: got %s, want %s", i, got, tt.want)
		}
	}
}
