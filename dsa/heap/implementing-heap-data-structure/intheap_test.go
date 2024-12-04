package heap

import "testing"

func TestIntHeap(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{3}, []int{3}},
		{[]int{}, []int{}},
		{[]int{9, 3, 1, 5, 2, 10}, []int{10, 9, 5, 3, 2, 1}},
	}
	for _, test := range tests {
		h := NewIntHeap(test.input)
		for i := 0; i < len(test.input); i++ {
			if h.Pop() != test.output[i] {
				t.Errorf("expected %d, got %d", test.output[i], h.Pop())
			}
		}
	}
}
