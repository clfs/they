package uci

import (
	"bytes"
	"errors"
	"testing"
)

func TestMarshalText(t *testing.T) {
	tests := []struct {
		in      Message
		want    []byte
		wantErr error
	}{
		{
			in:   &UCI{},
			want: []byte("uci"),
		},
		{
			in:   &IsReady{},
			want: []byte("isready"),
		},
		{
			in:   &SetOption{Name: "foo", Value: "bar"},
			want: []byte("setoption name foo value bar"),
		},
		{
			in:   &SetOption{Name: "foo"},
			want: []byte("setoption name foo"),
		},
		{
			in:   &SetOption{Name: "foo bar"},
			want: []byte("setoption name foo bar"),
		},
		{
			in:      &SetOption{Value: "foo"},
			wantErr: errMissingArg,
		},
	}

	for i, tt := range tests {
		got, gotErr := tt.in.MarshalText()
		if !errors.Is(gotErr, tt.wantErr) {
			t.Errorf("%d: error mismatch: got %v, want %v", i, gotErr, tt.wantErr)
		}
		if !bytes.Equal(got, tt.want) {
			t.Errorf("%d: output mismatch: got %q, want %q", i, got, tt.want)
		}
	}
}

func TestUnmarshalText(t *testing.T) {

}
