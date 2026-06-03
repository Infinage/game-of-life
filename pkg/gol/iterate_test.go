package gol

import (
	"reflect"
	"testing"
)

func TestGrid_Next(t *testing.T) {
	tests := []struct {
		name     string
		initial  Grid
		expected Grid
	}{
		{
			name: "Underpopulation: Single cell dies",
			initial: Grid{
				Cell{0, 0}: nil,
			},
			expected: Grid{},
		},
		{
			name: "Block: Still life (does not change)",
			initial: Grid{
				Cell{0, 0}: nil, Cell{0, 1}: nil,
				Cell{1, 0}: nil, Cell{1, 1}: nil,
			},
			expected: Grid{
				Cell{0, 0}: nil, Cell{0, 1}: nil,
				Cell{1, 0}: nil, Cell{1, 1}: nil,
			},
		},
		{
			name:     "Blinker: Oscillator (Horizontal to Vertical)",
			initial:  Grid{Cell{-1, 0}: nil, Cell{0, 0}: nil, Cell{1, 0}: nil},
			expected: Grid{Cell{0, -1}: nil, Cell{0, 0}: nil, Cell{0, 1}: nil},
		},
		{
			name: "Overpopulation: Center cell dies",
			initial: Grid{
				Cell{0, -1}: nil, Cell{0, 1}: nil,
				Cell{0, 0}:  nil,
				Cell{-1, 0}: nil, Cell{1, 0}: nil,
			},
			expected: Grid{
				Cell{-1, -1}: nil, Cell{1, -1}: nil,
				Cell{-1, 1}: nil, Cell{1, 1}: nil,
				Cell{0, -1}: nil, Cell{0, 1}: nil,
				Cell{-1, 0}: nil, Cell{1, 0}: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initial.Next()
			if !reflect.DeepEqual(tt.initial, tt.expected) {
				t.Errorf("values mismatch => want: %v, but got: %v", tt.expected, tt.initial)
			}
		})
	}
}
