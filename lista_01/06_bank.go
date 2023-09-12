package main

import (
	"fmt"

	Bank "github.com/soaresrl/dsa/pkg/bank"
	Het "github.com/soaresrl/dsa/pkg/het_linked_list"
)

func main () {
	lst := Het.NewList()

	lst = lst.Insert(Bank.Bank)
	lst = lst.Insert(Bank.Saving)
	lst = lst.Insert(Bank.Loyalty)

	lst.Print()

	fmt.Println()

	fmt.Println("Recursive")

	lst.PrintRec()

	fmt.Println()

	fmt.Println("Reverse Order")

	lst.PrintRecInv()

	fmt.Println()

	fmt.Printf("Is empty = %v\n", lst.IsEmpty())

	fmt.Println()
	
	fmt.Println("> Credit 100 for 2")

	lst.Credit(2, 100)

	fmt.Printf("Balance (2) = %f\n", lst.GetBalanceForAccount(2))

	fmt.Println("> Debit 50 for 2")

	lst.Debit(2, 50)

	fmt.Printf("Balance (2) = %f\n", lst.GetBalanceForAccount(2))

	fmt.Println()

	fmt.Println("> Debit 55 for 2")

	lst.Debit(2, 55)

	fmt.Printf("Balance (2) = %f\n", lst.GetBalanceForAccount(2))

	fmt.Println()

	fmt.Println("> Credit 100 for 3")

	lst.Credit(3, 100)

	fmt.Printf("Balance (3) = %f\n", lst.GetBalanceForAccount(3))

	bonus, _ := lst.GetBonusForAccount(3)

	fmt.Printf("Bonus = %f\n", bonus)
	
	fmt.Println("> Transfer 50 from 3 to 2")

	lst.TransferBetweenAccounts(50, 3, 2)

	fmt.Println()

	fmt.Printf("Balance (3) = %f\n", lst.GetBalanceForAccount(3))

	fmt.Println()

	fmt.Printf("Balance (2) = %f\n", lst.GetBalanceForAccount(2))

	fmt.Println("> Yield for account 2 (2.0)")

	lst.YieldFromAccount(2, 2.0)
	
	fmt.Println()
	
	fmt.Printf("Balance (2) = %f\n", lst.GetBalanceForAccount(2))

	fmt.Println("> Yield for account 3 (1.0)")

	lst.YieldFromAccount(3, 0.0)

	fmt.Println()

	fmt.Printf("Balance (3) = %f\n", lst.GetBalanceForAccount(3))

	fmt.Println()

	lst.Print()

	fmt.Println()

	fmt.Println("> Remove account 1")

	lst.Remove(1)

	lst.Print()
}
