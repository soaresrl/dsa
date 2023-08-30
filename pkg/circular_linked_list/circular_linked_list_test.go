package linked_list

import "testing"

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

	if lst.data != 10 {
		t.Errorf("Add test failed got %d, wanted %d", lst.data, 10)
	}

	if lst.next.data != 5 {
		t.Errorf("Add test failed got (next) %d, wanted %d", lst.next.data, 5)
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

	lst.Print()

	lst = lst.Remove(10)
	lst = lst.Remove(2)
	lst = lst.Remove(5)

	lst.Print()

	if lst.data != 30 {
		t.Errorf("Remove test failed expected head = %d, got %d", 30, lst.data)
	}

	if lst.next.data != 29 {
		t.Errorf("Remove test failed expected next = %d, got %d", 29, lst.next.data)
	}

	if lst.next.next.data != 4 {
		t.Errorf("Remove test failed expected next.next = %d, got %d", 4, lst.next.next.data)
	}

	if lst.next.next.next.data != 30 {
		t.Errorf("Remove test failed expected next.next.next (start) = %d, got %d", 30, lst.next.next.next.data)
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

	if lst.data != 30 {
		t.Errorf("Remove test failed expected head = %d, got %d", 30, lst.data)
	}

	if lst.next.data != 29 {
		t.Errorf("Remove test failed expected next = %d, got %d", 29, lst.next.data)
	}

	if lst.next.next.data != 4 {
		t.Errorf("Remove test failed expected next.next = %d, got %d", 4, lst.next.next.data)
	}

	if lst.next.next.next.data != 30 {
		t.Errorf("Remove test failed expected next.next.next (start) = %d, got %d", 30, lst.next.next.next.data)
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

	if node1.data != 12 {
		t.Errorf("Find test failed expected %d, got %d", 12, node1.data)
	}
	if node2.data != 40 {
		t.Errorf("Find test failed expected %d, got %d", 40, node2.data)
	}
	if node3.data != 5 {
		t.Errorf("Find test failed expected %d, got %d", 5, node3.data)
	}
}
