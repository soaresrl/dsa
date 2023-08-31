package circular_linked_list

import (
	"testing"
)

func TestCreateList(t *testing.T) {
	lst := NewList()

	if lst != nil {
		t.Errorf("Node should be nil, got %v", lst)
	}
}

func TestAdd(t *testing.T) {
	lst := NewList()

	lst = lst.Insert(5)
	lst = lst.Insert(10)

	if lst.Data != 10 {
		t.Errorf("Add test failed got %d, wanted %d", lst.Data, 10)
	}

	if lst.Next.Data != 5 {
		t.Errorf("Add test failed got (Next) %d, wanted %d", lst.Next.Data, 5)
	}
}

func TestRemove(t *testing.T) {
	lst := NewList()

	lst = lst.Insert(5)
	lst = lst.Insert(10)
	lst = lst.Insert(4)
	lst = lst.Insert(29)
	lst = lst.Insert(30)
	lst = lst.Insert(2)

	lst = lst.Remove(10)
	lst = lst.Remove(2)
	lst = lst.Remove(5)

	if lst.Data != 30 {
		t.Errorf("Remove test failed expected head = %d, got %d", 30, lst.Data)
	}

	if lst.Next.Data != 29 {
		t.Errorf("Remove test failed expected Next = %d, got %d", 29, lst.Next.Data)
	}

	if lst.Next.Next.Data != 4 {
		t.Errorf("Remove test failed expected Next.Next = %d, got %d", 4, lst.Next.Next.Data)
	}

	if lst.Next.Next.Next.Data != 30 {
		t.Errorf("Remove test failed expected Next.Next.Next (start) = %d, got %d", 30, lst.Next.Next.Next.Data)
	}
}

func TestRemoveRec(t *testing.T) {
	lst := NewList()

	lst = lst.Insert(5)
	lst = lst.Insert(10)
	lst = lst.Insert(4)
	lst = lst.Insert(29)
	lst = lst.Insert(30)
	lst = lst.Insert(2)

	lst = lst.RemoveRec(10, lst)
	lst = lst.RemoveRec(2, lst)
	lst = lst.RemoveRec(5, lst)

	if lst.Data != 30 {
		t.Errorf("Remove test failed expected head = %d, got %d", 30, lst.Data)
	}

	if lst.Next.Data != 29 {
		t.Errorf("Remove test failed expected Next = %d, got %d", 29, lst.Next.Data)
	}

	if lst.Next.Next.Data != 4 {
		t.Errorf("Remove test failed expected Next.Next = %d, got %d", 4, lst.Next.Next.Data)
	}

	if lst.Next.Next.Next.Data != 30 {
		t.Errorf("Remove test failed expected Next.Next.Next (start) = %d, got %d", 30, lst.Next.Next.Next.Data)
	}
}

func TestFind(t *testing.T) {
	lst := NewList()

	lst = lst.Insert(5)
	lst = lst.Insert(12)
	lst = lst.Insert(40)

	node1 := lst.Find(12)
	node2 := lst.Find(40)
	node3 := lst.Find(5)

	if node1.Data != 12 {
		t.Errorf("Find test failed expected %d, got %d", 12, node1.Data)
	}
	if node2.Data != 40 {
		t.Errorf("Find test failed expected %d, got %d", 40, node2.Data)
	}
	if node3.Data != 5 {
		t.Errorf("Find test failed expected %d, got %d", 5, node3.Data)
	}
}
