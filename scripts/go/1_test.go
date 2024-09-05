package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicatesKeepLastEasy(t *testing.T) {
	tests := []struct {
		name     string
		slice    []Element
		expected []Element
	}{
		{
			name:     "remove duplicates from slice",
			slice:    []Element{{key: "a"}, {key: "b"}, {key: "a"}, {key: "c"}},
			expected: []Element{{key: "b"}, {key: "a"}, {key: "c"}},
		},
		{
			name:     "empty slice",
			slice:    []Element{},
			expected: []Element{},
		},
		{
			name:     "no duplicates",
			slice:    []Element{{key: "a"}, {key: "b"}, {key: "c"}},
			expected: []Element{{key: "a"}, {key: "b"}, {key: "c"}},
		},
		{
			name: "complex1",
			slice: []Element{
				{"banana", 20},
				{"apple", 1},
				{"apple", 10},
				{"banana", 2},
				{"apple", 3},
				{"orange", 4},
				{"banana", 5},
			},
			expected: []Element{{"apple", 3}, {"orange", 4}, {"banana", 5}},
		},
		{
			name: "complex2",
			slice: []Element{
				{"apple", 10},
				{"banana", 2},
				{"orange", 4},
				{"banana", 2},
				{"grape", 5},
			},
			expected: []Element{{"apple", 10}, {"orange", 4},
				{"banana", 2},

				{"grape", 5}},
		},
		{

			name: "complex",
			slice: []Element{
				{"banana", 20},
				{"apple", 1},
				{"apple", 10},
				{"banana", 2},
				{"apple", 3},
				{"orange", 4},
				{"banana", 5},
				{"grape", 6},
				{"grape", 8},
				{"banana", 10},
				{"orange", 5},
				{"orange", 6},
				{"grape", 8},
				{"parrot", 1},
				{"grape", 80},
				{"banana", 888},
				{"apple", 3},
				{"apple", 3},
				{"apple", 4},
				{"parrot", 1},
			},
			expected: []Element{{"orange", 6}, {"grape", 80}, {"banana", 888}, {"apple", 4}, {"parrot", 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := removeDuplicatesKeepLastEasy(tt.slice)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("removeDuplicatesKeepLast(%v) got %v, want %v", tt.slice, actual, tt.expected)
			}
		})
	}
}

// TestRemoveDuplicatesKeepLast 测试 removeDuplicatesKeepLast 函数
func TestRemoveDuplicatesKeepLast(t *testing.T) {
	tests := []struct {
		name     string
		slice    []Element
		expected []Element
	}{
		{
			name:     "remove duplicates from slice",
			slice:    []Element{{key: "a"}, {key: "b"}, {key: "a"}, {key: "c"}},
			expected: []Element{{key: "b"}, {key: "a"}, {key: "c"}},
		},
		{
			name:     "empty slice",
			slice:    []Element{},
			expected: []Element{},
		},
		{
			name:     "no duplicates",
			slice:    []Element{{key: "a"}, {key: "b"}, {key: "c"}},
			expected: []Element{{key: "a"}, {key: "b"}, {key: "c"}},
		},
		{
			name: "complex1",
			slice: []Element{
				{"banana", 20},
				{"apple", 1},
				{"apple", 10},
				{"banana", 2},
				{"apple", 3},
				{"orange", 4},
				{"banana", 5},
			},
			expected: []Element{{"apple", 3}, {"orange", 4}, {"banana", 5}},
		},
		{
			name: "complex2",
			slice: []Element{
				{"apple", 10},
				{"banana", 2},
				{"orange", 4},
				{"banana", 2},
				{"grape", 5},
			},
			expected: []Element{{"apple", 10}, {"orange", 4},
				{"banana", 2},

				{"grape", 5}},
		},
		{
			name: "complex",
			slice: []Element{
				{"banana", 20},
				{"apple", 1},
				{"apple", 10},
				{"banana", 2},
				{"apple", 3},
				{"orange", 4},
				{"banana", 5},
				{"grape", 6},
				{"grape", 8},
				{"banana", 10},
				{"orange", 5},
				{"orange", 6},
				{"grape", 8},
				{"parrot", 1},
				{"grape", 80},
				{"banana", 888},
				{"apple", 3},
				{"apple", 3},
				{"apple", 4},
				{"parrot", 1},
			},
			expected: []Element{{"orange", 6}, {"grape", 80}, {"banana", 888}, {"apple", 4}, {"parrot", 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := removeDuplicatesKeepLast(tt.slice)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("removeDuplicatesKeepLast(%v) got %v, want %v", tt.slice, actual, tt.expected)
			}
		})
	}
}
