package redblack

import "fmt"

type Color int

const (
	Red Color = iota
	Black
)

type Node struct {
	Data   int
	Color  Color
	Left   *Node
	Right  *Node
	Parent *Node
}

type Tree struct {
	Root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (tree *Tree) rotateLeft(node *Node) {
	y := node.Right

	node.Right = y.Left

	if y.Left != nil {
		y.Left.Parent = node
	}

	y.Parent = node.Parent

	if node.Parent == nil {
		tree.Root = y
	} else if node == node.Parent.Left {
		node.Parent.Left = y
	} else {
		node.Parent.Right = y
	}

	y.Left = node
	node.Parent = y
}

func (tree *Tree) rotateRight(node *Node) {
	x := node.Left

	node.Left = x.Right

	if x.Right != nil {
		x.Right.Parent = node
	}

	x.Parent = node.Parent

	if node.Parent == nil {
		tree.Root = x
	} else if node.Parent.Left == node {
		node.Parent.Left = x
	} else {
		node.Parent.Right = x
	}

	x.Right = node

	node.Parent = x
}

func (tree *Tree) Insert(data int) {
	var y *Node = nil

	z := &Node{
		Data:  data,
		Color: Black,
	}

	x := tree.Root

	if x == nil {
		tree.Root = z
		return
	}

	for x != nil {
		y = x

		if z.Data < x.Data {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	z.Parent = y

	if y == nil {
		tree.Root = z
	} else if z.Data < y.Data {
		y.Left = z
	} else {
		y.Right = z
	}

	z.Color = Red

	tree.fixup(z)
}

func (tree *Tree) fixup(node *Node) {
	for node != tree.Root && node.Parent.Color == Red {
		if node.Parent == node.Parent.Parent.Left {
			y := node.Parent.Parent.Right

			if y != nil && y.Color == Red {
				node.Parent.Color = Black

				y.Color = Black

				node.Parent.Parent.Color = Red

				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {
					node = node.Parent

					tree.rotateLeft(node)
				}
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red

				tree.rotateRight(node.Parent.Parent)
			}
		} else {
			y := node.Parent.Parent.Left

			if y != nil && y.Color == Red {
				node.Parent.Color = Black

				y.Color = Black

				node.Parent.Parent.Color = Red

				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					node = node.Parent

					tree.rotateRight(node)
				}
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red

				tree.rotateLeft(node.Parent.Parent)
			}
		}
	}

	tree.Root.Color = Black
}

func (node *Node) Find(value int) *Node {
	if value == node.Data {
		return node
	}

	if value < node.Data {
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

func (tree *Tree) Find(value int) *Node {
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

	fmt.Printf("Data = %d, Color=%d \n", root.Data, root.Color)

	PrintPreOrder(root.Left)
	PrintPreOrder(root.Right)
}

func (tree *Tree) PrintPreOrder() {
	fmt.Println("---- Pre Order Print ----")
	if tree.Root == nil {
		return
	}

	fmt.Println(">Color 0 = Red\n>Color 1 = Black")

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

func FreeAbb(tree **Tree) {
	Free(&(*tree).Root)
}
