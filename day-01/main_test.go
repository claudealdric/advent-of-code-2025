package main

import "testing"

func TestApplyRotation(t *testing.T) {
	tests := []struct {
		currentValue int
		rotation     string
		want         int
	}{
		{currentValue: 50, rotation: "L68", want: 82},
		{currentValue: 82, rotation: "L30", want: 52},
		{currentValue: 52, rotation: "R48", want: 0},
		{currentValue: 0, rotation: "L5", want: 95},
		{currentValue: 95, rotation: "R60", want: 55},
		{currentValue: 55, rotation: "L55", want: 0},
		{currentValue: 0, rotation: "L1", want: 99},
		{currentValue: 99, rotation: "L99", want: 0},
		{currentValue: 0, rotation: "R14", want: 14},
		{currentValue: 14, rotation: "L82", want: 32},
	}

	for _, test := range tests {
		got, err := applyRotation(test.currentValue, test.rotation)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := test.want
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}

func TestParseRotationValue(t *testing.T) {
	tests := []struct {
		rotation string
		want     int
	}{
		{rotation: "R1", want: 1},
		{rotation: "R99", want: 99},
		{rotation: "L1", want: -1},
		{rotation: "L99", want: -99},
		{rotation: "L68", want: -68},
	}

	for _, test := range tests {
		got, err := parseRotationValue(test.rotation)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := test.want
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}
