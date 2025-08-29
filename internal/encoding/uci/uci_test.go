package uci

import (
	"bytes"
	"encoding"
	"testing"
)

func TestMarshalText(t *testing.T) {
	tests := []struct {
		in      encoding.TextMarshaler
		want    []byte
		wantErr bool
	}{
		{in: UCI{}, want: []byte("uci")},
		{in: IsReady{}, want: []byte("isready")},
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
