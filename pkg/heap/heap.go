package heap

import (
	"fmt"
)

type MaxHeap struct {
	Max			int
	Pos			int
	Priority		[]float32
}

func NewMaxHeap(max int) *MaxHeap {
	h := MaxHeap {
		Max: 			max,
		Pos: 			0,
		Priority:	make([]float32, max),
	}

	return &h
}

func (h *MaxHeap) Insert(priority float32) {
	if h.Pos <= h.Max - 1 {
		h.Priority[h.Pos] = priority

		h.FixAbove(h.Pos)

		h.Pos++
	} else {
		fmt.Println("Heap is full!!")
	}
}

func (h *MaxHeap) Remove() (bool, float32) {
	if h.Pos > 0 {
		top := h.Priority[0]

		h.Priority[0] = h.Priority[h.Pos - 1]
		h.Pos--

		h.FixBelow()

		return true, top
	} else {
		fmt.Println("Heap is empty!!")
	}

	return false, -1
}

func (h *MaxHeap) FixAbove(pos int) {
	for ; pos > 0 ; {
		fatherPos := (pos - 1) / 2

		if h.Priority[fatherPos] < h.Priority[pos]  {
			temp := h.Priority[pos]
			h.Priority[pos] = h.Priority[fatherPos]
			h.Priority[fatherPos] = temp
		} else {
			break
		}

		pos = fatherPos
	}
}

func (h *MaxHeap) FixBelow() {
	fatherPos := 0

	for ; h.Pos > 2*fatherPos + 1; {
		leftChild := 2 * fatherPos + 1
		rightChild := 2 * fatherPos + 2

		var childPos int
		if rightChild >= h.Pos {
			rightChild = leftChild
		}

		if h.Priority[leftChild] > h.Priority[rightChild] {
			childPos = leftChild
		} else {
			childPos = rightChild
		}

		if h.Priority[fatherPos] < h.Priority[childPos] {
			temp := h.Priority[fatherPos]
			h.Priority[fatherPos] = h.Priority[childPos]
			h.Priority[childPos] = temp
		} else {
			break
		}

		fatherPos = childPos
	}
}

func (h *MaxHeap) FindNode(a float32) (bool, int) {
	for i := 0; i < h.Pos; i++ {
		if a == h.Priority[i] {
			return true, i
		}
	}

	return false, -1
}

func (h *MaxHeap) ChangePriority(index int, value float32) {
	old := h.Priority[index]

	h.Priority[index] = value

	if value > old {
		h.FixAbove(index)
	} else {
		h.MaxHeapify(index)
	}
}

func (h *MaxHeap) MaxHeapify(i int) {
	l := 2*i + 1
	r := 2*i + 2

	largest := 0

	if l < h.Pos && h.Priority[l] > h.Priority[i] {
		largest = l
	} else {
		largest = i
	}

	if r < h.Pos && h.Priority[r] > h.Priority[largest] {
		largest = r
	}

	if largest != i {
		temp := h.Priority[i]
		h.Priority[i] = h.Priority[largest]
		h.Priority[largest] = temp

		h.MaxHeapify(largest)
	}
}

func (h *MaxHeap) Print() {
	for i := 0; i < h.Pos; i++ {
		fmt.Printf("%v\n", h.Priority[i])
	}
}

func FreeMaxHeap(h **MaxHeap) {
	(*h).Pos = 0
	(*h).Priority = nil

	h = nil
}