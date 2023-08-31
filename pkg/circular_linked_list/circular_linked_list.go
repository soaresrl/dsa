package circular_linked_list

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

func NewList() *Node {
	return nil
}

func (lst *Node) Insert(value int) *Node {
	node := &Node{Data: value}

	node.Next = lst

	if lst != nil {
		last := lst

		for ; last.Next != lst; last = last.Next {
		}

		last.Next = node
	} else {
		node.Next = node
	}

	return node
}

func (lst *Node) IsEmpty() bool {
	return lst == nil
}

func (lst *Node) Find(value int) *Node {
	p := lst

	for ; p.Next != lst; p = p.Next {
		if p.Data == value {
			return p
		}
	}

	// check if it's the last element
	if p.Data == value {
		return p
	}

	return nil
}

func Free(lst **Node) {
	if *lst == nil {
		return
	}

	if (*lst).Next == *lst {
		(*lst).Next = nil

		*lst = nil

		return
	}

	p := *lst

	for ; p.Next != *lst; p = p.Next {
		temp := p.Next

		p.Next = nil

		p = temp
	}

	p.Next = nil

	*lst = nil
}

func (lst *Node) Remove(value int) *Node {
	if lst == nil {
		return lst
	}

	// The list contains only one node and we want to remove it
	if lst.Data == value && lst.Next == lst {
		return nil
	}

	// Want to remove the first element
	if lst.Data == value {
		last := lst

		// Find reference to last element
		for ; last.Next != lst; last = last.Next {
		}

		// Make it point no Next of first
		last.Next = lst.Next

		return last.Next
	}

	// Search middle or end of list case
	last := lst
	for ; last.Next != lst && last.Next.Data != value; last = last.Next {
	}

	if last.Next.Data == value {
		p := last.Next

		last.Next = p.Next
		p = nil
	}

	return lst
}

func (lst *Node) RemoveRec(value int, sentinel *Node) *Node {
	if !lst.IsEmpty() {
		if lst.Next.Data == value {
			// lst is the last element and we want
			// to remove the first element
			if lst.Next == sentinel {
				// p is the first element (to be removed)
				p := lst.Next
				 
				// the Next of the last element
				// now points to the second element
				lst.Next = p.Next

				// lst now points to the second element
				lst = lst.Next
				
				p = nil
			} else {
				// p is the element to be removed
				p := lst.Next

				// the Next of the previous element
				// should point to the Next of the element to
				// be removed
				lst.Next = p.Next

				p = nil
			}
		} else {
			return lst.Next.RemoveRec(value, sentinel)
		}
	}

	return lst
}

func (lst *Node) Print() {
	fmt.Printf("start \n")

	p := lst

	for ; p.Next != lst; p = p.Next {
		fmt.Printf("Data=%d \n", p.Data)
	}
	fmt.Printf("Data=%d \n", p.Data)

	fmt.Printf("end \n")
}

func (lst *Node) PrintRec(sentinel *Node) {
	// while the Next element is not the first element
	// we keep recursing
	if lst.Next != sentinel {
		fmt.Printf("Data = %d\n", lst.Data)
		lst.Next.PrintRec(sentinel)
	} else { // if the Next is the first element just print the last
		fmt.Printf("Data= %d\n", lst.Data)
	}
}
