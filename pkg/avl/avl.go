package avl

import "fmt"

type Node struct {
	Key    int
	Factor int
	Parent *Node
	Left   *Node
	Right  *Node
}

func rotateRight(root *Node) *Node {
	parent := root.Parent
	t := root.Left
	m := t.Right

	t.Right = root
	root.Parent = t
	root.Left = m

	if m != nil {
		m.Parent = root
	}

	if parent != nil {
		if parent.Right == root {
			parent.Right = t
		} else {
			parent.Left = t
		}
	}

	t.Factor = 0
	root.Factor = 0

	return t
}

func rotateLeft(root *Node) *Node {
	parent := root.Parent

	t := root.Right
	m := t.Left

	root.Right = m

	if m != nil {
		m.Parent = root
	}

	t.Parent = parent

	if parent != nil {
		if parent.Right == root {
			parent.Right = t
		} else {
			parent.Left = t
		}
	}

	t.Factor = 0
	root.Factor = 0

	return t
}

func rotateRightLeft(root *Node) *Node {
	t := root.Right
	s := t.Left

	t_2 := s.Left
	t_3 := s.Right

	s.Left = root
	s.Right = t

	root.Right = t_2
	t.Left = t_3

	if s.Factor == 1 {
		root.Factor = -1
	} else {
		root.Factor = 0
	}

	if s.Factor == -1 {
		t.Factor = 1
	} else {
		t.Factor = 0
	}

	s.Factor = 0
	return s
}

func rotateLeftRight(root *Node) *Node {
	p := root.Left
	q := p.Right

	t_2 := q.Left
	t_3 := q.Right

	q.Left = p
	q.Right = root

	p.Right = t_2
	root.Left = t_3

	if q.Factor == 1 {
		p.Factor = -1
	} else {
		p.Factor = 0
	}

	if q.Factor == -1 {
		root.Factor = 1
	} else {
		root.Factor = 0
	}

	q.Factor = 0

	return q
}

func insert(node *Node, key int, flag *int) *Node {
	if node == nil {
		node = &Node{
			Key: key,
		}

		*flag = 1
	} else if node.Key > key {
		node.Left = insert(node.Left, key, flag)

		if node.Left.Parent == nil {
			node.Left.Parent = node
		}

		if *flag != 0 {
			switch node.Factor {
			case 1:
				node.Factor = 0
				*flag = 0
			case 0:
				node.Factor = -1
			case -1:
				if node.Left.Factor == -1 {
					node = rotateRight(node)
				} else {
					node = rotateLeftRight(node)
				}
				*flag = 0
			}
		}
	} else if node.Key < key {
		node.Right = insert(node.Right, key, flag)

		if node.Right.Parent == nil {
			node.Right.Parent = node
		}

		if *flag != 0 {
			switch node.Factor {
			case -1:
				node.Factor = 0
				*flag = 0
			case 0:
				node.Factor = -1
			case 1:
				if node.Right.Factor == -1 {
					node = rotateLeft(node)
				} else {
					node = rotateRightLeft(node)
				}
				*flag = 0
			}
		}
	}

	return node
}

func (root *Node) Insert(key int) *Node {
	var flag int = 0

	return insert(root, key, &flag)
}

func remove(node *Node, key int, delta_h *int) *Node {
	if node == nil {
		return nil
	} else if key < node.Key {
		node.Left = remove(node.Left, key, delta_h)
		node.Factor -= *delta_h

		if node.Factor == 2 {
			switch node.Factor {
			case 1:
				node = rotateLeft(node)
				*delta_h = -1
			case 0:
				node = rotateLeft(node)
				*delta_h = 0
			case -1:
				node = rotateRightLeft(node)
				*delta_h = -1
			}
		} else {
			if node.Factor == 1 {
				*delta_h = 0
			} else {
				*delta_h = -1
			}
		}
	} else if key > node.Key {
		node.Right = remove(node.Right, key, delta_h)
		node.Factor -= *delta_h

		if node.Factor == 2 {
			switch node.Factor {
			case -1:
				node = rotateRight(node)
				*delta_h = -1
			case 0:
				node = rotateRight(node)
				*delta_h = 0
			case 1:
				node = rotateLeftRight(node)
				*delta_h = -1
			}
		} else {
			if node.Factor == 1 {
				*delta_h = 0
			} else {
				*delta_h = -1
			}
		}
	} else {
		if node.Left == nil && node.Right == nil {
			*delta_h = -1
			node = nil
		} else if node.Left == nil {
			temp := node
			node = node.Right
			node.Parent = temp.Parent
		} else if node.Right == nil {
			temp := node
			node = node.Left
			node.Parent = temp.Parent
		} else {
			sucessor := node.Right

			for ; sucessor.Left != nil; sucessor = sucessor.Left {
			}

			node.Key = sucessor.Key
			sucessor.Key = key

			if sucessor.Parent.Left == sucessor {
				sucessor.Parent.Left = sucessor.Right
			} else {
				sucessor.Parent.Right = sucessor.Right
			}

			if sucessor.Right != nil {
				sucessor.Right.Parent = sucessor.Parent
			}
		}
	}

	return node
}

func (root *Node) Remove(key int) *Node {
	var delta_h int = 0

	return remove(root, key, &delta_h)
}

func PrintPreOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Printf("%d \n", root.Key)

	PrintPreOrder(root.Left)
	PrintPreOrder(root.Right)
}

func (node *Node) PrintPreOrder() {
	fmt.Println("---- Pre Order Print ----")
	if node == nil {
		return
	}

	PrintPreOrder(node)

	fmt.Println("---- Pre Order Print End ----")
}

func Free(node **Node) {
	if (*node).Left != nil {
		Free(&(*node).Left)
	}

	if (*node).Right != nil {
		Free(&(*node).Right)
	}

	(*node).Parent = nil

	*node = nil
}
