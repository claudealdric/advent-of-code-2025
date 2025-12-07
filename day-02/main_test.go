package main

import "testing"

func TestIsInvalidId(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{input: 1010, want: true},
		{input: 12, want: false},
		{input: 123, want: false},
		{input: 1188511885, want: true},
		{input: 1188511880, want: false},
		{input: 446446, want: true},
		{input: 38593859, want: true},
	}

	for _, test := range tests {
		got := isInvalidId(test.input)
		want := test.want
		if got != want {
			t.Errorf("got %v, want %v (input: %d)", got, want, test.input)
		}
	}
}
