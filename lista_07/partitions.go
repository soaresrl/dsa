package main

import (
	"github.com/soaresrl/dsa/pkg/partition"
	"fmt"
)

func main() {
	p := partition.NewPartition(10)

	el1 := p.MakeSet(1)
	el2 := p.MakeSet(2)
	el3 := p.MakeSet(3)
	el4 := p.MakeSet(4)
	el5 := p.MakeSet(5)
	el6 := p.MakeSet(6)
	el7 := p.MakeSet(7)
	el8 := p.MakeSet(8)
	el9 := p.MakeSet(9)
	el10 := p.MakeSet(10)

	fmt.Println("Make singleton sets for 1 to 10")
	fmt.Println("Apply Union for 1, 2, 3")
	fmt.Println("Apply Union for 4, 5, 6")
	fmt.Println("Apply Union for 7, 9, 10, 8")

	p.Union(el1, el2)
	p.Union(el1, el3)
	p.Union(el4, el5)
	p.Union(el4, el6)
	p.Union(el7, el9)
	p.Union(el7, el10)
	p.Union(el7, el8)

	el := p.FindSet(el8)

	fmt.Printf("Recovering Rep for %d, Rep = %d\n", el8.Data, el.Data)
	fmt.Printf("7 is on same set as 10: %v\n", p.SameSet(el7, el10))
	fmt.Printf("1 is on Same set as 10: %v\n", p.SameSet(el1, el10))

	partition.Free(&p)
}