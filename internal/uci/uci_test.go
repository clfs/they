package uci

import (
	"testing"
)

func TestID_UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		want    ID
		wantErr bool
	}{
		{
			name: "short name",
			text: "id name MyBot",
			want: ID{Name: "MyBot"},
		},
		{
			name: "long name",
			text: "id name Cool Bot",
			want: ID{Name: "Cool Bot"},
		},
		{
			name: "short author",
			text: "id author John",
			want: ID{Author: "John"},
		},
		{
			name: "long author",
			text: "id author John Smith",
			want: ID{Author: "John Smith"},
		},
		{
			name:    "empty string",
			text:    "",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got ID
			err := got.UnmarshalText([]byte(test.text))
			gotErr := (err != nil)

			if got != test.want {
				t.Errorf("ID.UnmarshalText(%q): got %#v, want %#v", test.text, got, test.want)
			}
			if gotErr != test.wantErr {
				t.Errorf("ID.UnmarshalText(%q): gotErr %v, wantErr %v", test.text, gotErr, test.wantErr)
			}
		})
	}
}

func TestUCI_AppendText(t *testing.T) {
	tests := []struct {
		name    string
		message ID
		want    string
		wantErr bool
	}{
		{
			name:    "short name",
			message: ID{Name: "MyBot"},
			want:    "id name MyBot",
		},
		{
			name:    "long name",
			message: ID{Name: "Cool Bot"},
			want:    "id name Cool Bot",
		},
		{
			name:    "short author",
			message: ID{Author: "John"},
			want:    "id author John",
		},
		{
			name:    "long author",
			message: ID{Author: "John Smith"},
			want:    "id author John Smith",
		},
		{
			name:    "no name or author",
			message: ID{},
			wantErr: true,
		},
		{
			name:    "name and author",
			message: ID{Name: "MyBot", Author: "John"},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.message.AppendText(nil)
			gotErr := (err != nil)

			if string(got) != test.want {
				t.Errorf("%#v.AppendText(nil): got %q, want %q", test.message, got, test.want)
			}
			if gotErr != test.wantErr {
				t.Errorf("%#v.AppendText(nil): gotErr %v, wantErr %v", test.message, gotErr, test.wantErr)
			}
		})
	}
}
