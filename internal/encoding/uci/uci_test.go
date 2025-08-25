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
		wantErr error
	}{
		{in: UCI{}, want: []byte("uci")},
	}

	for i, tt := range tests {
		got, gotErr := tt.in.MarshalText()
		if gotErr != tt.wantErr {
			t.Errorf("%d: error mismatch: got %v, want %v", i, gotErr, tt.wantErr)
		}
		if !bytes.Equal(got, tt.want) {
			t.Errorf("%d: output mismatch: got %v, want %v", i, got, tt.want)
		}
	}
}
