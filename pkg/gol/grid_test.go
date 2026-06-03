package gol

import (
	"reflect"
	"slices"
	"testing"
)

func Test_GridSortingOrder(t *testing.T) {
	tests := []struct {
		initial  []Cell
		expected []Cell
	}{
		{
			initial:  []Cell{{1, 0}, {0, 1}, {0, 0}},
			expected: []Cell{{0, 0}, {1, 0}, {0, 1}},
		},
		{
			initial:  []Cell{{-1, 0}, {0, -1}, {0, 0}},
			expected: []Cell{{0, -1}, {-1, 0}, {0, 0}},
		},
	}

	for _, test := range tests {
		got := test.initial
		slices.SortFunc(got, SortCell)
		if !slices.Equal(got, test.expected) {
			t.Errorf("values mismatch.\nWant: %v\nGot: %v", test.expected, got)
		}
	}
}

func Test_CellNeighbours(t *testing.T) {
	tests := []struct {
		initial  Cell
		expected []Cell
	}{
		{
			initial:  Cell{-1, 0},
			expected: []Cell{{-2, -1}, {-2, 0}, {-2, 1}, {-1, -1}, {-1, 1}, {0, -1}, {0, 0}, {0, 1}},
		},
		{
			initial:  Cell{-1, -1},
			expected: []Cell{{-2, -2}, {-2, -1}, {-2, 0}, {-1, -2}, {-1, 0}, {0, -2}, {0, -1}, {0, 0}},
		},
		{
			initial:  Cell{0, 0},
			expected: []Cell{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}},
		},
		{
			initial:  Cell{-1, 1},
			expected: []Cell{{-2, 0}, {-2, 1}, {-2, 2}, {-1, 0}, {-1, 2}, {0, 0}, {0, 1}, {0, 2}},
		},
	}

	for _, test := range tests {
		if got := test.initial.Neighbours(); !reflect.DeepEqual(test.expected, got) {
			t.Errorf("values mismatch for %v.\nWant: %v\nGot: %v", test.initial, test.expected, got)
		}
	}
}
