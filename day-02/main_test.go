package main

import (
	"slices"
	"testing"
)

func TestAllElementsAreEqual(t *testing.T) {
	tests := []struct {
		input []string
		want  bool
	}{
		{input: []string{"12", "12", "12"}, want: true},
		{input: []string{"12", "12", "13"}, want: false},
	}

	for _, test := range tests {
		got := allElementsAreEqual(test.input)
		want := test.want
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestGetChunks(t *testing.T) {
	tests := []struct {
		input       string
		totalChunks int
		want        []string
	}{
		{input: "121212", totalChunks: 3, want: []string{"12", "12", "12"}},
		{input: "121212", totalChunks: 2, want: []string{"121", "212"}},
		{input: "121212", totalChunks: 6, want: []string{"1", "2", "1", "2", "1", "2"}},
	}

	for _, test := range tests {
		got, err := getChunks(test.input, test.totalChunks)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := test.want
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

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

func TestIsInvalidId2(t *testing.T) {
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
		{input: 565656, want: true},
		{input: 824824824, want: true},
		{input: 2121212121, want: true},
	}

	for _, test := range tests {
		got := isInvalidId2(test.input)
		want := test.want
		if got != want {
			t.Errorf("got %v, want %v (input: %d)", got, want, test.input)
		}
	}
}
