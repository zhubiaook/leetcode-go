package main

import (
	"cmp"
	"math"
	"reflect"
	"testing"
)

var ints = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = []float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = []string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

type testData[T cmp.Ordered] struct {
	name  string
	items []T
	want  []T
}

func TestInsertionSort(t *testing.T) {
	tests1 := []testData[int]{
		{
			name:  "ints",
			items: ints,
			want:  []int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845},
		},
	}

	tests2 := []testData[float64]{
		{
			name:  "float64",
			items: float64s,
			want:  []float64{math.Inf(-1), -959.7485, -784, 2.3, 7.8, 7.8, 59, 74.3, 238.2, 905, 9845.768, math.Inf(1)},
		},
	}

	tests3 := []testData[string]{
		{
			name:  "strings",
			items: strings,
			want:  []string{"", "%*&^*&^&", "***", "Hello", "bar", "f00", "foo", "foo"},
		},
	}

	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.items)
			if !reflect.DeepEqual(tt.items, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", tt.items, tt.want)
			}
		})
	}

	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.items)
			if !reflect.DeepEqual(tt.items, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", tt.items, tt.want)
			}
		})
	}

	for _, tt := range tests3 {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.items)
			if !reflect.DeepEqual(tt.items, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", tt.items, tt.want)
			}
		})
	}

}

func TestBubbleSort(t *testing.T) {
	tests1 := []testData[int]{
		{
			name:  "ints",
			items: ints,
			want:  []int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845},
		},
	}

	tests2 := []testData[float64]{
		{
			name:  "float64",
			items: float64s,
			want:  []float64{math.Inf(-1), -959.7485, -784, 2.3, 7.8, 7.8, 59, 74.3, 238.2, 905, 9845.768, math.Inf(1)},
		},
	}

	tests3 := []testData[string]{
		{
			name:  "strings",
			items: strings,
			want:  []string{"", "%*&^*&^&", "***", "Hello", "bar", "f00", "foo", "foo"},
		},
	}

	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.items)
			if !reflect.DeepEqual(tt.items, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", tt.items, tt.want)
			}
		})
	}

	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.items)
			if !reflect.DeepEqual(tt.items, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", tt.items, tt.want)
			}
		})
	}

	for _, tt := range tests3 {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.items)
			if !reflect.DeepEqual(tt.items, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", tt.items, tt.want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	tests1 := []testData[int]{
		{
			name:  "ints",
			items: ints,
			want:  []int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845},
		},
	}

	tests2 := []testData[float64]{
		{
			name:  "float64",
			items: float64s,
			want:  []float64{math.Inf(-1), -959.7485, -784, 2.3, 7.8, 7.8, 59, 74.3, 238.2, 905, 9845.768, math.Inf(1)},
		},
	}

	tests3 := []testData[string]{
		{
			name:  "strings",
			items: strings,
			want:  []string{"", "%*&^*&^&", "***", "Hello", "bar", "f00", "foo", "foo"},
		},
	}

	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.items)
			if !reflect.DeepEqual(tt.items, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", tt.items, tt.want)
			}
		})
	}

	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.items)
			if !reflect.DeepEqual(tt.items, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", tt.items, tt.want)
			}
		})
	}

	for _, tt := range tests3 {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.items)
			if !reflect.DeepEqual(tt.items, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", tt.items, tt.want)
			}
		})
	}
}
