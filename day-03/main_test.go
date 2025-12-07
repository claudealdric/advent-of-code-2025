package main

import (
	"testing"
)

func TestGetLargestJoltage(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "987654321111111", want: "98"},
		{input: "811111111111119", want: "89"},
		{input: "234234234234278", want: "78"},
		{input: "818181911112111", want: "92"},
	}

	for _, test := range tests {
		got := getLargestJoltage(test.input)
		want := test.want
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}

func TestGetLargestJoltage2(t *testing.T) {
	t.Run("subsequence length of 2", func(t *testing.T) {
		tests := []struct {
			input string
			want  string
		}{
			{input: "81119", want: "89"},
			{input: "98765", want: "98"},
			{input: "987654321111111", want: "98"},
			{input: "234234234234278", want: "78"},
			{input: "818181911112111", want: "92"},
		}

		for _, test := range tests {
			got := getLargestJoltage2(test.input, 2)
			want := test.want
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		}
	})

	t.Run("subsequence length of 12", func(t *testing.T) {
		tests := []struct {
			input string
			want  string
		}{
			{input: "987654321111111", want: "987654321111"},
			{input: "811111111111119", want: "811111111119"},
		}

		for _, test := range tests {
			got := getLargestJoltage2(test.input, 12)
			want := test.want
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		}
	})
}
