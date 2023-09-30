package set

import (
	"fmt"
)

type Set struct {
	BitSet	[]byte
	Size		int
}

func NewBitSet(size int) *Set {
	bitSet := make([]byte, (size + 7) / 8)

	return &Set{BitSet: bitSet, Size: (size + 7) / 8}
}

func (s *Set) Add(elem int) {
	pos := int(elem / 8)
	mask := uint8((elem - 1) % 8)

	s.BitSet[pos] |= uint8(1) << mask
}

func (s *Set) Remove(elem int) {
	pos := int(elem / 8)
	mask := uint8((elem - 1) % 8)

	s.BitSet[pos] = s.BitSet[pos] & (^(uint8(1) << mask))
}

func (s *Set) Complement() *Set {
	bitSet := make([]byte, s.Size)

	complement := Set{BitSet: bitSet, Size: s.Size}

	i := 0
	for _, byteContainer := range s.BitSet {
		complement.BitSet[i] = ^byteContainer
		i++
	}

	return &complement
}

func (s *Set) HasElement(elem int) bool {
	pos := int(elem / 8)
	mask := uint8((elem - 1) % 8)

	test := s.BitSet[pos] & (uint8(1) << mask)

	if test == 0 {
		return false
	}

	return true
}

func (s *Set) Print() {
	for i := (s.Size - 1); i >= 0; i-- {
		fmt.Printf("%08b ", s.BitSet[i])
	}
}

func (s *Set) NumberOfElements() int {
	n := 0
	for i := 0; i < s.Size; i++ {
		for j := uint8(1); j < 128; j = j << 1 {
			if (s.BitSet[i] & j) != 0 {
				n++	
			}
		}

		if (s.BitSet[i] & 128) != 0 {
			n++	
		}
	}

	return n
}

func Free(s **Set) {
	(*s).BitSet = nil
	(*s).Size = 0

	s = nil
}

func Union(a *Set, b *Set) *Set {
	bitSet := make([]byte, a.Size)

	union := Set{BitSet: bitSet, Size: a.Size}

	for i := 0; i < a.Size; i++ {
		union.BitSet[i] = a.BitSet[i] | b.BitSet[i]
	}

	return &union
}

func Intersection(a *Set, b *Set) *Set {
	bitSet := make([]byte, a.Size)

	inter := Set{BitSet: bitSet, Size: a.Size}

	for i := 0; i < a.Size; i++ {
		inter.BitSet[i] = a.BitSet[i] & b.BitSet[i]
	}

	return &inter
}

func Difference(a *Set, b *Set) *Set {
	bitSet := make([]byte, a.Size)

	diff := Set{BitSet: bitSet, Size: a.Size}

	for i := 0; i < a.Size; i++ {
		diff.BitSet[i] = a.BitSet[i] & (^b.BitSet[i])
	}

	return &diff
}

func IsSubset(a *Set, b *Set) bool {
	bitSet := make([]byte, a.Size)

	test := Set{BitSet: bitSet, Size: a.Size}

	for i := 0; i < a.Size; i++ {
		test.BitSet[i] = a.BitSet[i] & b.BitSet[i]
	}

	for i := 0; i < test.Size; i++ {
		if(test.BitSet[i] != b.BitSet[i]){
			return false
		}
	}

	return true
}

func Equal(a *Set, b *Set) bool {
	for i := 0; i < a.Size; i++ {
		if a.BitSet[i] != b.BitSet[i] {
			return false
		}
	}

	return true
}


