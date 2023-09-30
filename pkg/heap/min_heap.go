package heap

import (
	"fmt"
)

type MinHeap struct {
	Max			int
	Pos			int
	Priority		[]float32
}

func NewMinHeap(max int) *MinHeap {
	h := MinHeap {
		Max: 			max,
		Pos: 			0,
		Priority:	make([]float32, max),
	}

	return &h
}

func (h *MinHeap) Insert(priority float32) {
	if h.Pos <= h.Max - 1 {
		h.Priority[h.Pos] = priority

		h.FixAbove(h.Pos)

		h.Pos++
	} else {
		fmt.Println("Heap is full!!")
	}
}

func (h *MinHeap) Remove() (bool, float32) {
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

func (h *MinHeap) FixAbove(pos int) {
	for ; pos > 0 ; {
		fatherPos := (pos - 1) / 2

		if h.Priority[fatherPos] > h.Priority[pos]  {
			temp := h.Priority[pos]
			h.Priority[pos] = h.Priority[fatherPos]
			h.Priority[fatherPos] = temp
		} else {
			break
		}

		pos = fatherPos
	}
}

func (h *MinHeap) FixBelow() {
	fatherPos := 0

	for ; h.Pos > 2*fatherPos + 1; {
		leftChild := 2 * fatherPos + 1
		rightChild := 2 * fatherPos + 2

		var childPos int
		if rightChild >= h.Pos {
			rightChild = leftChild
		}

		if h.Priority[leftChild] < h.Priority[rightChild] {
			childPos = leftChild
		} else {
			childPos = rightChild
		}

		if h.Priority[fatherPos] > h.Priority[childPos] {
			temp := h.Priority[fatherPos]
			h.Priority[fatherPos] = h.Priority[childPos]
			h.Priority[childPos] = temp
		} else {
			break
		}

		fatherPos = childPos
	}
}

func (h *MinHeap) FindNode(a float32) (bool, int) {
	for i := 0; i < h.Pos; i++ {
		if a == h.Priority[i] {
			return true, i
		}
	}

	return false, -1
}

func (h *MinHeap) ChangePriority(index int, value float32) {
	old := h.Priority[index]

	h.Priority[index] = value
	
	if value > old {
		h.MinHeapify(index)
	} else {
		h.FixAbove(index)
	}
}

func (h *MinHeap) MinHeapify(i int) {
	l := 2*i + 1
	r := 2*i + 2

	smallest := 0

	if l < h.Pos && h.Priority[l] < h.Priority[i] {
		smallest = l
	} else {
		smallest = i
	}

	if r < h.Pos && h.Priority[r] < h.Priority[smallest] {
		smallest = r
	}

	if smallest != i {
		temp := h.Priority[i]
		h.Priority[i] = h.Priority[smallest]
		h.Priority[smallest] = temp

		h.MinHeapify(smallest)
	}
}

func (h *MinHeap) Print() {
	for i := 0; i < h.Pos; i++ {
		fmt.Printf("%v\n", h.Priority[i])
	}
}

func FreeMinHeap(h **MinHeap) {
	(*h).Pos = 0
	(*h).Priority = nil

	h = nil
}