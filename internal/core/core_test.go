package core

import "testing"

func TestSquare_String(t *testing.T) {
	tests := []struct {
		name   string
		square Square
		want   string
	}{
		{name: "zero value", want: "A1"},
		{"F6", F6, "F6"},
		{"167", 167, "Square(167)"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.square.String()
			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}
