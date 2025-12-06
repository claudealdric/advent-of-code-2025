package main

import "testing"

func TestApplyRotation(t *testing.T) {
	tests := []struct {
		currentValue      int
		rotation          string
		expectedNewValue  int
		expectedZeroCount int
	}{
		{currentValue: 50, rotation: "L68", expectedNewValue: 82, expectedZeroCount: 1},
		{currentValue: 82, rotation: "L30", expectedNewValue: 52, expectedZeroCount: 0},
		{currentValue: 52, rotation: "R48", expectedNewValue: 0, expectedZeroCount: 1},
		{currentValue: 0, rotation: "L5", expectedNewValue: 95, expectedZeroCount: 0},
		{currentValue: 95, rotation: "R60", expectedNewValue: 55, expectedZeroCount: 1},
		{currentValue: 55, rotation: "L55", expectedNewValue: 0, expectedZeroCount: 1},
		{currentValue: 0, rotation: "L1", expectedNewValue: 99, expectedZeroCount: 0},
		{currentValue: 99, rotation: "L99", expectedNewValue: 0, expectedZeroCount: 1},
		{currentValue: 0, rotation: "R14", expectedNewValue: 14, expectedZeroCount: 0},
		{currentValue: 14, rotation: "L82", expectedNewValue: 32, expectedZeroCount: 1},
		{currentValue: 14, rotation: "L182", expectedNewValue: 32, expectedZeroCount: 2},
		{currentValue: 50, rotation: "R250", expectedNewValue: 0, expectedZeroCount: 3},
	}

	for _, test := range tests {
		got, err := applyRotation(test.currentValue, test.rotation)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		Equal(t, got.newValue, test.expectedNewValue)

		if got.zeroCount != test.expectedZeroCount {
			t.Errorf(
				"got %d, want %d (initial value: %d, rotation: %s, expected new value: %d)",
				got.zeroCount,
				test.expectedZeroCount,
				test.currentValue,
				test.rotation,
				test.expectedNewValue,
			)
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

func Equal(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
