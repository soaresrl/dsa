package main

import (
	LinkedList "github.com/soaresrl/dsa/pkg/linked_list"
)

//type Node struct {
//	data int
//	next *Node
//}
//
//func create() *Node {
//	return nil
//}
//
//func insert(lst *Node, value int) *Node {
//	node := &Node{data: value}
//
//	node.next = lst
//
//	return node
//}
//
//func is_empty(lst *Node) bool {
//	return lst == nil
//}
//
//func find(lst *Node, value int) *Node {
//	for p := lst; p != nil; p = p.next {
//		if p.data == value {
//			return p
//		}
//	}
//
//	return nil
//}
//
//func free(lst **Node) {
//	if *lst == nil {
//		return
//	}
//
//	free(&(*lst).next)
//	(*lst).next = nil
//	*lst = nil
//}
//
//func remove(lst *Node, value int) *Node {
//	var slow *Node
//
//	p := lst
//	for ; p != nil && p.data != value; p = p.next {
//		slow = p
//	}
//
//	if p == nil {
//		return lst
//	}
//	if slow == nil {
//		lst = p.next
//	} else {
//		slow.next = p.next
//		p = nil
//	}
//
//	return lst
//
//}
//
//func remove_rec(lst *Node, value int) *Node {
//	if !is_empty(lst) {
//		if lst.data == value {
//			lst = lst.next
//		} else {
//			lst.next = remove_rec(lst.next, value)
//		}
//	}
//
//	return lst
//}
//
//func print(lst *Node) {
//	fmt.Printf("start \n")
//
//	for p := lst; p != nil; p = p.next {
//		fmt.Printf("data=%d \n", p.data)
//	}
//
//	fmt.Printf("end \n")
//}
//
//func print_rec(lst *Node) {
//	if lst != nil {
//		fmt.Printf("data = %d\n", lst.data)
//		print_rec(lst.next)
//	}
//}
//
//func print_rec_inv(lst *Node) {
//	if lst != nil {
//		print_rec_inv(lst.next)
//		fmt.Printf("data = %d\n", lst.data)
//	}
//}

func main() {
	//lst := create()
	//fmt.Printf("is empty ? %v\n", is_empty(lst))
	//lst = insert(lst, 1)
	//lst = insert(lst, 4)
	//lst = insert(lst, 10)
	//lst = insert(lst, 5)

	//print(lst)

	////lst = remove(lst, 10)
	//lst = remove_rec(lst, 10)
	//print(lst)

	//print_rec(lst)
	//print_rec_inv(lst)

	//free(&lst)

	//print(lst)
	lst := LinkedList.NewList()

	lst = lst.Insert(4)
	lst = lst.Insert(10)

	lst.PrintRecInv()
}
