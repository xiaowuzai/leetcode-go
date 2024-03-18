package reversestring

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{
			name:  "Test case 1",
			input: []byte{'h', 'e', 'l', 'l', 'o'},
			want:  []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name:  "Test case 2",
			input: []byte{'H', 'i', '!', ' '},
			want:  []byte{' ', '!', 'i', 'H'},
		},
		{
			name:  "Test case 3",
			input: []byte{},
			want:  []byte{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.input)
			if string(tt.input) != string(tt.want) {
				t.Errorf("Got %v, wanted %v", tt.input, tt.want)
			}
		})
	}
}

func TestReverseString2(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{
			name:  "Test case 1",
			input: []byte{'h', 'e', 'l', 'l', 'o'},
			want:  []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name:  "Test case 2",
			input: []byte{'a', 'b', 'c', 'd'},
			want:  []byte{'d', 'c', 'b', 'a'},
		},
		{
			name:  "Test case 3",
			input: []byte{},
			want:  []byte{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString2(tt.input)
			if string(tt.input) != string(tt.want) {
				t.Errorf("got %v, want %v", tt.input, tt.want)
			}
		})
	}
}
