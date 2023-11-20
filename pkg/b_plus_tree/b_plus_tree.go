package bptree

type Node struct {
	isLeaf   bool
	keys     []int
	pointers []interface{}
	next     *Node
	parent   *Node
}

type BPTree struct {
	order int
	root  *Node
}

func (node *Node) update(key int, data interface{}) bool {
	for i, k := range node.keys {
		if k == key {
			node.pointers[i] = data

			return true
		}
	}

	return false
}

// Remove data associated with key from this node
func (node *Node) remove(key int) {
	// Find where is the key in this node
	place := -1
	for i, k := range node.keys {
		// found key in node
		if k == key {
			place = i
			break
		}
	}

	tempKeys := node.keys[:place]
	tempKeys = append(tempKeys, node.keys[place+1:]...)

	tempPointers := node.pointers[:place]
	tempPointers = append(tempPointers, node.pointers[place+1:]...)

	node.keys = tempKeys
	node.pointers = tempPointers
}

func NewBPTree(order int) *BPTree {
	root := &Node{
		isLeaf:   true,
		keys:     make([]int, 0, order-1),
		pointers: make([]interface{}, 0, order),
	}

	tree := &BPTree{
		order: order,
		root:  root,
	}

	return tree
}

func (tree *BPTree) findLeaf(key int) *Node {
	node := tree.root
	for !node.isLeaf {
		for i := 0; i < len(node.keys); i++ {
			if key <= node.keys[i] {
				node, _ = node.pointers[i].(*Node)
				break
			}

			if i == len(node.keys)-1 {
				node, _ = node.pointers[i+1].(*Node)
				break
			}
		}
	}

	return node
}

func (tree *BPTree) leafInsert(leaf *Node, key int, data interface{}) {
	// find where to insert the key
	place := -1
	for i, k := range leaf.keys {
		if k > key {
			place = i
			break
		}
	}

	// empty or last
	if place == -1 {
		leaf.keys = append(leaf.keys, key)
		leaf.pointers = append(leaf.pointers, data)
		return
	}

	// other
	leaf.keys = append(leaf.keys[:place], append([]int{key}, leaf.keys[place:]...)...)
	leaf.pointers = append(leaf.pointers[:place], append([]interface{}{data}, leaf.pointers[place:]...)...)
}

func (tree *BPTree) Find(key int) interface{} {
	// Find leaf node containing key range
	leaf := tree.findLeaf(key)

	for i, k := range leaf.keys {
		if k == key {
			data := leaf.pointers[i]

			return data
		}
	}

	return nil
}

func (tree *BPTree) FixAbove(firstNode, secondNode *Node, key int) {
	// insert newNode least key in parent node
	parent := firstNode.parent

	if parent == nil { // if root node create new root
		newRoot := &Node{
			isLeaf:   false,
			keys:     make([]int, 0, tree.order-1),
			pointers: make([]interface{}, 0, tree.order),
		}

		newRoot.keys = append(newRoot.keys, firstNode.keys[len(firstNode.keys)-1])
		newRoot.pointers = append(newRoot.pointers, firstNode, secondNode)

		firstNode.parent = newRoot
		secondNode.parent = newRoot

		tree.root = newRoot
	} else if len(parent.keys) < tree.order-1 { // can insert key in parent
		place := -1

		for i, p := range parent.pointers {
			if p == firstNode {
				place = i
				break
			}
		}

		if place == -1 {
			parent.keys = append(parent.keys, secondNode.keys[0])
			parent.pointers = append(parent.pointers, secondNode)
		} else {
			parent.keys = append(parent.keys[:place], append([]int{firstNode.keys[len(firstNode.keys)-1]}, parent.keys[place:]...)...)
			parent.pointers = append(parent.pointers[:place+1], append([]interface{}{secondNode}, parent.pointers[place+1:]...)...)
		}
		secondNode.parent = parent
	} else { // need to split parent and fix above recursively
		temp := &Node{
			isLeaf:   parent.isLeaf,
			keys:     parent.keys,
			pointers: parent.pointers,
			parent:   parent.parent,
		}

		newNode := &Node{
			isLeaf:   false,
			keys:     make([]int, 0, tree.order-1),
			pointers: make([]interface{}, 0, tree.order),
		}

		// Find where to put the new key in the temp node
		place := -1

		for i, p := range temp.pointers {
			if p == firstNode {
				place = i
				break
			}
		}

		if place == -1 {
			temp.keys = append(temp.keys, secondNode.keys[0])
			temp.pointers = append(temp.pointers, secondNode)
		} else {
			temp.keys = append(temp.keys[:place], append([]int{firstNode.keys[len(firstNode.keys)-1]}, temp.keys[place:]...)...)
			temp.pointers = append(temp.pointers[:place+1], append([]interface{}{secondNode}, temp.pointers[place+1:]...)...)
		}

		idxMedian := (tree.order)/2 - 1

		parent.keys = make([]int, 0, tree.order-1)
		parent.pointers = make([]interface{}, 0, tree.order)

		parent.keys = append(parent.keys, temp.keys[0:idxMedian+1]...)
		parent.pointers = append(parent.pointers, temp.pointers[0:idxMedian+1]...)

		newNode.keys = append(newNode.keys, temp.keys[idxMedian+1:]...)
		newNode.pointers = append(newNode.pointers, temp.pointers[idxMedian+1:]...)

		newNode.parent = parent.parent

		tree.FixAbove(parent, newNode, parent.keys[len(parent.keys)-1])
	}
}

func (tree *BPTree) Insert(key int, data interface{}) {
	// search for the leaf node to insert the key
	leaf := tree.findLeaf(key)

	if leaf.update(key, data) {
		return
	}

	if len(leaf.keys) < tree.order-1 {
		tree.leafInsert(leaf, key, data)
	} else { // need to split
		newNode := &Node{
			isLeaf:   true,
			keys:     make([]int, 0, tree.order-1),
			pointers: make([]interface{}, 0, tree.order),
		}

		// allocate temp node with greater order to
		// store the key to be inserted and all the previous inserted
		temp := &Node{
			isLeaf:   true,
			keys:     make([]int, 0, tree.order),
			pointers: make([]interface{}, 0, tree.order+1),
		}

		temp.keys = append(temp.keys, leaf.keys...)
		temp.pointers = append(temp.pointers, leaf.pointers...)

		tree.leafInsert(temp, key, data)

		newNode.next = leaf.next
		leaf.next = newNode

		idxMedian := (tree.order)/2 - 1

		leaf.keys = make([]int, 0)
		leaf.pointers = make([]interface{}, 0)

		leaf.keys = append(leaf.keys, temp.keys[0:idxMedian+1]...)
		leaf.pointers = append(leaf.pointers, temp.pointers[0:idxMedian+1]...)

		newNode.keys = append(newNode.keys, temp.keys[idxMedian+1:]...)
		newNode.pointers = append(newNode.pointers, temp.pointers[idxMedian+1:]...)

		// Fix Above
		tree.FixAbove(leaf, newNode, leaf.keys[len(leaf.keys)-1])
	}
}

func (tree *BPTree) updateParentKeys(node *Node, removedKey int) {

}

func (tree *BPTree) Remove(key int) {
	data := tree.Find(key)

	if data == nil {
		return
	}

	leaf := tree.findLeaf(key)

	leaf.remove(key)

	if len(leaf.keys) >= (tree.order/2)-1 {
		parent := leaf.parent

		place := -1
		for i, k := range parent.keys {
			if k == key {
				place = i
				break
			}
		}

		if place == -1 {
			return
		}

		parent.keys[place] = leaf.keys[len(leaf.keys)-1]
		return
	}

	// The node has less elements than the minimum allowed
	// for the tree's order
	if len(leaf.keys) < (tree.order/2)-1 {
		leftIdx := -1
		rightIdx := -1

		for i, p := range leaf.parent.pointers {
			if p == leaf {
				leftIdx = i - 1
				rightIdx = i + 1

				break
			}
		}

		var leftNode *Node
		var rightNode *Node

		if leftIdx != -1 {
			leftNode = leaf.parent.pointers[leftIdx].(*Node)
		}
		if rightIdx != -1 {
			rightNode = leaf.parent.pointers[rightIdx].(*Node)
		}
		// see if it can borrow from left brother
		if leftNode != nil && len(leftNode.keys) > (tree.order/2)-1 {
			tempKeys := make([]int, 0, tree.order-1)
			tempPointers := make([]interface{}, 0, tree.order)

			tempKeys = append(tempKeys, leftNode.keys[len(leftNode.keys)-1])
			tempKeys = append(tempKeys, leaf.keys...)
			tempPointers = append(tempPointers, leftNode.pointers[len(leftNode.pointers)-1])
			tempPointers = append(tempPointers, leaf.pointers...)

			leaf.keys = tempKeys
			leaf.pointers = tempPointers

			// delete key from sibling
			leftNode.keys = leftNode.keys[:len(leftNode.keys)-1]
			leftNode.pointers = leftNode.pointers[:len(leftNode.pointers)-1]

			// update parent
			// update parent
			parent := leaf.parent

			place := -1
			for i, k := range parent.keys {
				if k == key {
					place = i
					break
				}
			}

			if place == -1 {
				return
			}

			leaf.parent.keys[place-1] = leftNode.keys[len(leftNode.keys)-1]
		} else if rightNode != nil && len(rightNode.keys) > (tree.order/2)-1 {
			leaf.keys = append(leaf.keys, rightNode.keys[0])
			leaf.pointers = append(leaf.pointers, rightNode.pointers[0])

			// delete key from sibling
			rightNode.keys = rightNode.keys[1:]
			rightNode.pointers = rightNode.pointers[1:]

			// update parent
			parent := leaf.parent

			place := -1
			for i, k := range parent.keys {
				if k == key {
					place = i
					break
				}
			}

			if place == -1 {
				return
			}

			parent.keys[place] = leaf.keys[len(leaf.keys)-1]
		} else { // merge two nodes
			if leftNode != nil && len(leftNode.keys) < tree.order-1 {
				leftNode.keys = append(leftNode.keys, leaf.keys...)
				leftNode.pointers = append(leftNode.pointers, leaf.pointers...)

				tempKeys := leaf.parent.keys[:leftIdx-1]
				tempKeys = append(tempKeys, leaf.parent.keys[leftIdx+1:]...)

				tempPointers := leaf.parent.pointers[:leftIdx]
				tempPointers = append(tempPointers, leaf.parent.pointers[leftIdx+2:]...)

				tempKeys[leftIdx] = leaf.keys[len(leaf.keys)-1]

				leaf.parent.keys = tempKeys
				leaf.parent.pointers = tempPointers

				leaf.parent.keys[0] = leftNode.keys[len(leftNode.keys)-1]
			} else if rightNode != nil && len(rightNode.keys) < tree.order-1 {
				leaf.keys = append(leaf.keys, rightNode.keys...)
				leaf.pointers = append(leaf.pointers, rightNode.pointers...)

				tempKeys := leaf.parent.keys[:rightIdx-3]
				tempKeys = append(tempKeys, leaf.parent.keys[rightIdx-1:]...)

				tempPointers := leaf.parent.pointers[:rightIdx-1]
				tempPointers = append(tempPointers, leaf.parent.pointers[rightIdx+1:]...)

				leaf.parent.keys = tempKeys
				leaf.parent.pointers = tempPointers
			}
		}
	}
}

func Free(node **Node) {
	for _, p := range (*node).pointers {
		// if pointer is to a node instead to data
		if _, ok := p.(*Node); ok {
			n, _ := p.(*Node)
			Free(&n)
		} else { // pointer to data
			p = nil
		}
	}

	(*node).parent = nil

	*node = nil
}

func FreeBTree(tree **BPTree) {
	Free(&(*tree).root)

	*tree = nil
}
