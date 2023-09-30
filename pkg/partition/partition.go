package partition

type Node struct {
	Parent		*Node
	Rank			int
	Data			int
}

type Partition struct {
	Sets			[]*Node
}

func NewPartition(n int) *Partition {
	partition := Partition{
		Sets: make([]*Node, n),
	}

	return &partition
}

func (p *Partition) MakeSet(x int) *Node {
	node := Node{
		Parent: 	nil,
		Rank:		0,
		Data:		x,
	}

	node.Parent = &node

	return &node
}

func (p *Partition) Link(x, y *Node) {
	if x.Rank > y.Rank {
		y.Parent = x
	} else {
		x.Parent = y

		if x.Rank == y.Rank {
			y.Rank++
		}
	}
}

func (p *Partition) FindSet(x *Node) *Node {
	if x != x.Parent {
		return p.FindSet(x.Parent)
	}

	return x
}

func (p *Partition) Union(x, y *Node) {
	p.Link(p.FindSet(x), p.FindSet(y))
}

func (p *Partition) SameSet(x, y *Node) bool {
	repX := p.FindSet(x)
	repY := p.FindSet(y)

	return repX == repY
}

func Free(p **Partition) {
	(*p).Sets = nil

	p = nil
}