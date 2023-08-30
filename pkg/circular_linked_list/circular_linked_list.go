package linked_list

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func NewList() *Node {
	return nil
}

func (lst *Node) Insert(value int) *Node {
	node := &Node{data: value}

	node.next = lst

	if lst != nil {
		last := lst

		for ; last.next != lst; last = last.next {
		}

		last.next = node
	} else {
		node.next = node
	}

	return node
}

func (lst *Node) IsEmpty() bool {
	return lst == nil
}

func (lst *Node) Find(value int) *Node {
	p := lst

	for ; p.next != lst; p = p.next {
		if p.data == value {
			return p
		}
	}

	// check if it's the last element
	if p.data == value {
		return p
	}

	return nil
}

func Free(lst **Node) {
	if *lst == nil {
		return
	}

	Free(&(*lst).next)
	(*lst).next = nil
	*lst = nil
}

func (lst *Node) Remove(value int) *Node {
	if lst == nil {
		return lst
	}

	// The list contains only one node and we want to remove it
	if lst.data == value && lst.next == lst {
		return nil
	}

	// Want to remove the first element
	if lst.data == value {
		last := lst

		// Find reference to last element
		for ; last.next != lst; last = last.next {
		}

		// Make it point no next of first
		last.next = lst.next

		return last.next
	}

	// Search middle or end of list case
	last := lst
	for ; last.next != lst && last.next.data != value; last = last.next {
	}

	if last.next.data == value {
		p := last.next

		last.next = p.next
		p = nil
	}

	return lst
}

func (lst *Node) RemoveRec(value int, sentinel *Node) *Node {
	if !lst.IsEmpty() {
		if lst.next.data == value {
			// lst is the last element and we want
			// to remove the first element
			if lst.next == sentinel {
				// p is the first element (to be removed)
				p := lst.next

				// the next of the last element
				// now points to the second element
				lst.next = p.next

				// lst now points to the second element
				lst = lst.next

				p = nil
			} else {
				// p is the element to be removed
				p := lst.next

				// the next of the previous element
				// should point to the next of the element to
				// be removed
				lst.next = p.next

				p = nil
			}
		} else {
			lst.next.RemoveRec(value, sentinel)
		}
	}

	return lst
}

func (lst *Node) Print() {
	fmt.Printf("start \n")

	for p := lst; p != nil; p = p.next {
		fmt.Printf("data=%d \n", p.data)
	}

	fmt.Printf("end \n")
}

func (lst *Node) PrintRec(sentinel *Node) {
	// while the next element is not the first element
	// we keep recursing
	if lst.next != sentinel {
		fmt.Printf("data = %d\n", lst.data)
		lst.next.PrintRec(sentinel)
	} else { // if the next is the first element just print the last
		fmt.Printf("data= %d\n", lst.data)
	}
}
