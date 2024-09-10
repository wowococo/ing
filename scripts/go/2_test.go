package main

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		element  int
		expected []int
	}{
		{
			name:     "remove element from slice",
			slice:    []int{1, 2, 3, 4, 5},
			element:  3,
			expected: []int{1, 2, 4, 5},
		},
		{
			name:     "remove element from slice",
			slice:    []int{1, 2, 3, 4, 5},
			element:  1,
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "remove element from slice",
			slice:    []int{1, 2, 3, 4, 5},
			element:  2,
			expected: []int{1, 3, 4, 5},
		},
		{
			name:     "remove element from slice",
			slice:    []int{1, 2, 3, 4, 5},
			element:  4,
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "remove element from slice",
			slice:    []int{1, 2, 3, 4, 5},
			element:  5,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "remove non-existing element",
			slice:    []int{1, 2, 3, 4, 5},
			element:  6,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "remove from empty slice",
			slice:    []int{},
			element:  3,
			expected: []int{},
		},
		{
			name:     "remove from nil slice",
			slice:    nil,
			element:  3,
			expected: nil,
		},
		{
			name:     "remove from slice with duplicates",
			slice:    []int{1, 2, 3, 3, 4},
			element:  3,
			expected: []int{1, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := removeElement(tt.slice, tt.element)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("removeElement(%v, %d) got %v, want %v", tt.slice, tt.element, actual, tt.expected)
			}
		})
	}
}
