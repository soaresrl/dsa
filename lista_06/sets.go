package main 

import (
	"github.com/soaresrl/dsa/pkg/set"
	"fmt"
)

func main(){
	set_obj := set.NewBitSet(24)

	set_obj.Add(4)
	set_obj.Add(10)

	set_obj.Print()

	fmt.Println()
}