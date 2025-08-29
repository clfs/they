package uci

import (
	"bytes"
	"testing"
)

func TestMarshalText(t *testing.T) {
	tests := []struct {
		in      Command
		want    []byte
		wantErr bool
	}{
		{in: &UCI{}, want: []byte("uci")},
		{in: &IsReady{}, want: []byte("isready")},
	}

	for i, tt := range tests {
		got, err := tt.in.MarshalText()
		gotErr := (err != nil)
		if gotErr != tt.wantErr {
			t.Errorf("%d: error mismatch: got %v, want %v", i, gotErr, tt.wantErr)
		}
		if !bytes.Equal(got, tt.want) {
			t.Errorf("%d: output mismatch: got %v, want %v", i, got, tt.want)
		}
	}
}

func TestParseString(t *testing.T) {
	tests := []struct {
		in      string
		want    Command
		wantErr bool
	}{
		{in: "uci", want: &UCI{}},
		{in: " uci", want: &UCI{}},
		{in: "uci\n", want: &UCI{}},
		{in: "uci\r\n", want: &UCI{}},

		{in: "isready", want: &IsReady{}},
	}

	for i, tt := range tests {
		got, err := ParseString(tt.in)
		gotErr := (err != nil)
		if gotErr != tt.wantErr {
			t.Errorf("%d: error mismatch: got %v, want %v", i, gotErr, tt.wantErr)
		}
		if got != tt.want {
			t.Errorf("%d: output mismatch: got %#v, want %#v", i, got, tt.want)
		}
	}
}
