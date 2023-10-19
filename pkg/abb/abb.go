package abb

import "fmt"

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Parent *Node
}

type Abb struct {
	Root *Node
}

func NewAbb() *Abb {
	a := Abb{}

	return &a
}

func (node *Node) Insert(value int) {
	if value > node.Key {
		if node.Right != nil {
			node.Right.Insert(value)
		} else {
			node.Right = &Node{Key: value, Parent: node}
		}
	} else if value < node.Key {
		if node.Left != nil {
			node.Left.Insert(value)
		} else {
			node.Left = &Node{Key: value, Parent: node}
		}
	}
}

func (node *Node) Remove(value int) {
	if value < node.Key {
		node.Left.Remove(value)
	} else if value > node.Key {
		node.Right.Remove(value)
	} else {
		if node.Left == nil && node.Right == nil { // nó sem filhos
			node = nil
		} else if node.Left == nil {
			temp := node
			node = node.Right
			node.Parent = temp.Parent
		} else if node.Right == nil {
			temp := node
			node = node.Left
			node.Parent = temp.Parent
		} else { // nó tem os dois filhos, busca o sucessor
			sucessor := node.Right

			for ; sucessor.Left != nil; sucessor = sucessor.Left {
			}

			node.Key = sucessor.Key
			sucessor.Key = value

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
}

func (tree *Abb) Insert(value int) {
	if tree.Root == nil {
		node := Node{Key: value}

		tree.Root = &node
		return
	}

	tree.Root.Insert(value)
}

func (tree *Abb) InsertIterative(value int) {
	if tree.Root == nil {
		node := Node{Key: value}

		tree.Root = &node
		return
	}

	node := &Node{Key: value}
	var parent *Node
	x := tree.Root

	for x != nil {
		parent = x
		if value < parent.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	node.Parent = parent

	if value < parent.Key {
		parent.Left = node
	} else {
		parent.Right = node
	}

}

func (tree *Abb) Remove(value int) {
	if tree.Root == nil {
		return
	}

	tree.Root.Remove(value)
}

func (node *Node) Find(value int) *Node {
	if value == node.Key {
		return node
	}

	if value < node.Key {
		if node.Left != nil {
			return node.Left.Find(value)
		} else {
			return nil
		}
	} else {
		if node.Right != nil {
			return node.Right.Find(value)
		} else {
			return nil
		}
	}
}

func (tree *Abb) Find(value int) *Node {
	if tree.Root == nil {
		return nil
	}

	node := tree.Root.Find(value)

	return node
}

func PrintPreOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Printf("%d \n", root.Key)

	PrintPreOrder(root.Left)
	PrintPreOrder(root.Right)
}

func (tree *Abb) PrintPreOrder() {
	fmt.Println("---- Pre Order Print ----")
	if tree.Root == nil {
		return
	}

	PrintPreOrder(tree.Root)

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

func FreeAbb(tree **Abb) {
	Free(&(*tree).Root)
}
