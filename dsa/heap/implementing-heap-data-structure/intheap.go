package heap

// intHeap is a max heap
type intHeap []int

// NewIntHeap returns a new intHeap
func NewIntHeap(items []int) *intHeap {
	h := make(intHeap, 0, len(items))
	for _, v := range items {
		h.Push(v)
	}

	return &h
}

// IsEmpty returns true if the heap is empty, false otherwise
func (h *intHeap) IsEmpty() bool {
	return len(*h) == 0
}

// Push adds an element to the heap
func (h *intHeap) Push(item int) {
	*h = append(*h, item)
	h.heapifyUp(len(*h) - 1)
}

// Pop removes and returns the maximum element from the heap
func (h *intHeap) Pop() int {
	if len(*h) == 0 {
		return 0
	}
	max := (*h)[0]
	lastIndex := len(*h) - 1
	(*h)[0] = (*h)[lastIndex]
	*h = (*h)[:lastIndex]
	h.heapifyDown(0)
	return max
}

// heapifyUp maintains the max heap property after an element is added
func (h *intHeap) heapifyUp(index int) {
	if index == 0 {
		return
	}

	parentIndex := (index - 1) / 2
	for index > 0 && (*h)[index] > (*h)[parentIndex] {
		(*h)[index], (*h)[parentIndex] = (*h)[parentIndex], (*h)[index]
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
}

// heapifyDown maintains the max heap property after an element is removed
func (h *intHeap) heapifyDown(index int) {
	maxIndex := h.maxChildIndex(index)

	for maxIndex > 0 && (*h)[index] < (*h)[maxIndex] {
		(*h)[index], (*h)[maxIndex] = (*h)[maxIndex], (*h)[index]
		index = maxIndex
		maxIndex = h.maxChildIndex(index)
	}
}

func (h *intHeap) maxChildIndex(index int) int {
	leftIndex := 2*index + 1
	if leftIndex >= len(*h) {
		return 0
	}

	rightIndex := leftIndex + 1
	if rightIndex >= len(*h) {
		return leftIndex
	}

	if (*h)[leftIndex] >= (*h)[rightIndex] {
		return leftIndex
	}
	return rightIndex
}
