package extendible_hash

import (
	"fmt"
	"strconv"
)

type Bucket struct {
	LocalDepth	int
	Size		int
	Items		[]int
}

type HashTable struct {
	GlobalDepth	int
	BucketSize 	int
	Buckets		[]*Bucket
}

func hashfunction(key, depth int) int {
   return key & ((1 << depth) - 1)
}

func NewHashTable(bucketSize int) *HashTable {
	var buckets []*Bucket

	for i := 0; i < 2; i++ {
		bucket := Bucket{
			LocalDepth: 1,		
			Size: bucketSize,
			Items: make([]int, 0),
		}

		buckets = append(buckets, &bucket)
	}

	table := HashTable{
		GlobalDepth: 1,
		BucketSize: bucketSize,
		Buckets: buckets,
	}

	return &table
}

func (h *HashTable) Insert(value int) {
	bucketID := hashfunction(value, h.GlobalDepth)

	status := h.Buckets[bucketID].Insert(value)

	if status == 0 {
		// case 2 - can be split
		if h.Buckets[bucketID].LocalDepth < h.GlobalDepth {
			h.Buckets[bucketID].LocalDepth++

			bucket := Bucket{
				LocalDepth: h.Buckets[bucketID].LocalDepth,
				Size: h.BucketSize,
				Items: make([]int, 0),
			}

			temp := make([]int, len(h.Buckets[bucketID].Items))

			copy(temp, h.Buckets[bucketID].Items)

			h.Buckets[bucketID].Items = h.Buckets[bucketID].Items[:0]

			for _, item := range temp {
				index := hashfunction(item, h.GlobalDepth)

				if index == bucketID {
					bucket.Items = append(bucket.Items, item)
				} else {
					h.Buckets[bucketID].Items = append(h.Buckets[bucketID].Items, item)
				}
			}

			bucket.Insert(value)

			h.Buckets[bucketID] = &bucket
		} else { // case 3 bucket is complete and has one reference
			h.GlobalDepth++
			h.Buckets[bucketID].LocalDepth++
			temp := make([]int, len(h.Buckets[bucketID].Items))

			copy(temp, h.Buckets[bucketID].Items)

			for i := 1 << (h.GlobalDepth-1); i < 1 << h.GlobalDepth; i++ {
				pairID := i ^ (1 << (h.GlobalDepth - 1))
				
				if pairID == bucketID {
					bucket := Bucket {
						LocalDepth: h.GlobalDepth,
						Size: h.BucketSize,
						Items: make([]int, 0),
					}
					h.Buckets = append(h.Buckets, &bucket)
					continue
				}

				h.Buckets = append(h.Buckets, h.Buckets[pairID])
			}
			
			h.Buckets[bucketID].Items = h.Buckets[bucketID].Items[:0]

			for _, item := range temp {
				index := hashfunction(item, h.GlobalDepth)

				h.Buckets[index].Insert(item)
			}

			bucketID = hashfunction(value, h.GlobalDepth)

			h.Buckets[bucketID].Insert(value)
		}
	}
}

func (h *HashTable) Get(key int) int {
	pos := hashfunction(key, h.GlobalDepth)

	for _, v := range h.Buckets[pos].Items {
		if v == key {
			return v
		}
	}

	return -1
}

func (h *HashTable) Remove(key int) {
	pos := hashfunction(key, h.GlobalDepth)

	status := h.Buckets[pos].Remove(key)

	if status == -1 {
		return
	}

	// 1. Try to merge two buckets
	// search for pair bucket
	pairID := pos ^ (1 << (h.GlobalDepth - 1))
	
	// if the two entries on address list 
        // don't point to the same bucket
	if h.Buckets[pos] != h.Buckets[pairID] {
		// see if the current size of two buckets 
		// don't exceeds the capacity of one bucket
		size := len(h.Buckets[pos].Items) + len(h.Buckets[pairID].Items)

		if size <= h.BucketSize {
			h.Buckets[pairID].Items = append(h.Buckets[pairID].Items, h.Buckets[pos].Items...)

			h.Buckets[pos] = h.Buckets[pairID]
			h.Buckets[pos].LocalDepth--
		}
	}
	// end of 1.

	// 2. Shrink the addresses list if needed
	count := 0
	for _, b := range h.Buckets {
		if h.GlobalDepth - b.LocalDepth == 1 {
			count++
		}
	}

	if count == len(h.Buckets) {
		h.GlobalDepth--
		h.Buckets = h.Buckets[0:1<<h.GlobalDepth]
	}


	// end of 2.
}

func Free(h **HashTable) {
	(*h).Buckets = nil

	*h = nil
}

func (h *HashTable) Print() {
	fmt.Println("Hash Global Depth: ", h.GlobalDepth)
	for i, v := range h.Buckets {
		fmt.Printf("Bucket index: %d \n", i)
		v.Print()
	}
}

func (b *Bucket) Insert(key int) int {
	for _, v := range b.Items {
		if v == key {
			return -1
		}
	}

	// Case 1 - bucket has space
	if len(b.Items) < b.Size {
		b.Items = append(b.Items, key)
		
		return 1
	}

	return 0
}

func (b *Bucket) Remove(key int) int {
	i := -1
	for j, k := range b.Items {
		if k == key {
			i = j
			break
		}
	}

	if i == -1 {
		return -1
	}

	b.Items[i] = b.Items[len(b.Items) - 1]
	b.Items = b.Items[:len(b.Items)-1]

	return 1
}

func (b *Bucket) Print() {
	fmt.Printf("Address %p\n", b)
	fmt.Printf("Local Depth: %d\n", b.LocalDepth)

	for _, v := range b.Items {
		n := int64(v)

		fmt.Printf("value = %d (%v)\n", v, strconv.FormatInt(n, 2))
	}
}
