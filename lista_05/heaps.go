package main

import (
	"github.com/soaresrl/dsa/pkg/heap"
	"fmt"
)

func main() {
	max_heap := heap.NewMaxHeap(10)

	max_heap.Insert(10)
	max_heap.Insert(8)
	max_heap.Insert(5)
	max_heap.Insert(2)
	max_heap.Insert(12)
	max_heap.Insert(40)
	max_heap.Insert(13)
	max_heap.Insert(20)
	max_heap.Insert(90)
	max_heap.Insert(50)

	fmt.Println("-- Max heap priority array --")
	max_heap.Print()
	fmt.Println()

	rmvd, value := max_heap.Remove()

	if rmvd {
		fmt.Printf("removed = %v\n", value)
	}

	fmt.Println("-- Updated Max heap --")

	max_heap.Print()
	fmt.Println()

	hasEl, indexEl := max_heap.FindNode(20)

	if hasEl {
		fmt.Println("Found node 20, changing its priority to 14")
		max_heap.ChangePriority(indexEl, 14)
	}

	fmt.Println("-- Updated max heap --")

	max_heap.Print()
	fmt.Println()

	min_heap := heap.NewMinHeap(10)

	min_heap.Insert(10)
	min_heap.Insert(8)
	min_heap.Insert(5)
	min_heap.Insert(2)
	min_heap.Insert(12)
	min_heap.Insert(40)
	min_heap.Insert(13)
	min_heap.Insert(20)
	min_heap.Insert(90)
	min_heap.Insert(50)

	fmt.Println("-- Min heap priority array --")

	min_heap.Print()
	fmt.Println()

	rmvd, value = min_heap.Remove()

	if rmvd {
		fmt.Printf("removed = %v\n", value)
	}

	rmvd, value = min_heap.Remove()

	if rmvd {
		fmt.Printf("removed = %v\n", value)
	}

	fmt.Println("-- Updated Min heap --")

	min_heap.Print()
	
	fmt.Println()

	hasEl, indexEl = min_heap.FindNode(40)

	if hasEl {
		fmt.Println("Found node 40, changing its priority to 9")

		min_heap.ChangePriority(indexEl, 9)
	}

	fmt.Println("-- Updated Min heap --")

	min_heap.Print()
	
	fmt.Println()

	heap.FreeMaxHeap(&max_heap)
	heap.FreeMinHeap(&min_heap)
}