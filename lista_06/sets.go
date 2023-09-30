package main 

import (
	"github.com/soaresrl/dsa/pkg/set"
	"fmt"
)

func main(){
	set_1 := set.NewBitSet(24)

	set_1.Add(4)
	set_1.Add(10)
	set_1.Add(5)
	set_1.Add(2)

	copy_set_1 := set.NewBitSet(24)

	copy_set_1.Add(4)
	copy_set_1.Add(10)
	copy_set_1.Add(5)
	copy_set_1.Add(2)

	fmt.Println("Set 1 --")

	set_1.Print()

	fmt.Println()

	set_1_subset := set.NewBitSet(24)

	set_1_subset.Add(5)
	set_1_subset.Add(2)

	fmt.Println("SubSet of 1 --")

	set_1_subset.Print()

	fmt.Println()

	is_subset := set.IsSubset(set_1, set_1_subset)

	if is_subset {
		fmt.Println("SubSet of 1 Is subset")
	} else {
		fmt.Println("SubSet of 1 is not a subset of 1")
	}

	set_2 := set.NewBitSet(24)

	set_2.Add(5)
	set_2.Add(7)

	fmt.Println("Set 2 --")

	set_2.Print()

	fmt.Println()

	set_3 := set.NewBitSet(24)

	set_3.Add(5)
	set_3.Add(8)
	set_3.Add(9)
	
	fmt.Println("Set 3 --")

	set_3.Print()

	fmt.Println()

	are_equal := set.Equal(set_1, set_2)

	if are_equal {
		fmt.Println("1 is equal to 2")
	} else {
		fmt.Println("1 is not equal to 2")
	}
	
	are_equal = set.Equal(set_1, copy_set_1)

	if are_equal {
		fmt.Println("copy of 1 is equal to 1")
	} else {
		fmt.Println("copy of 1 is not equal to 1")
	}

	union_set := set.Union(set_1, set_2)

	fmt.Println("1 Union 2 --")

	union_set.Print()

	fmt.Println()

	inter_set := set.Intersection(set_3, set_2)

	fmt.Println("3 Inter 2 --")

	inter_set.Print()

	fmt.Println()

	diff_set := set.Difference(set_3, set_2)

	fmt.Println("3 Diff 2 --")

	diff_set.Print()

	fmt.Println()

	comp_set := set_1.Complement()

	fmt.Println("1 Complement --")

	comp_set.Print()

	fmt.Println()

	set_1.Remove(4)

	fmt.Println("1 After removing element 4 --")

	set_1.Print()

	fmt.Println()

	has_element := set_1.HasElement(4)

	if has_element {
		fmt.Println("Set 1 has element 4")
	} else {
		fmt.Println("Set 1 has not element 4")
	}

	has_element = set_1.HasElement(5)

	if has_element {
		fmt.Println("Set 1 has element 5")
	} else {
		fmt.Println("Set 1 has not element 5")
	}

	fmt.Printf("number of elements of Set 1's Complement = %d\n", comp_set.NumberOfElements())

	set.Free(&set_1)
	set.Free(&set_2)
	set.Free(&set_3)
	set.Free(&copy_set_1)

	set.Free(&inter_set)
	set.Free(&diff_set)
	set.Free(&union_set)
	set.Free(&comp_set)
}