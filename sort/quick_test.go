package sort_test

import (
	"reflect"
	"testing"

	"github.com/neepoo/BFS/sort"
)

// Test function for QuickSort using table-driven tests
func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"unsorted", []int{3, 6, 8, 10, 1, 2, 1}, []int{1, 1, 2, 3, 6, 8, 10}},
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.QuickSort(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("QuickSort() = %v, want %v", tt.input, tt.expected)
			}
		})
	}
}
