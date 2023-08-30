package het_linked_list

import (
	"fmt"

	Bank "github.com/soaresrl/dsa/pkg/bank"
)

type Node struct {
	data Bank.Accountable
	next *Node
}

func NewList() *Node {
	return nil
}

func (lst *Node) Insert(account_type int) *Node {
	var account Bank.Accountable

	number := 0

	if lst != nil {
		number = lst.data.GetID()
	}

	switch account_type {
	case Bank.Bank:
		account = &Bank.BankAccount{Number: number + 1}

	case Bank.Saving:
		account = &Bank.SavingAccount{Number: number + 1}

	case Bank.Loyalty:
		account = &Bank.LoyaltyAccount{Number: number + 1}

	default:
		fmt.Println("Specified account type not supported.")
	}

	node := &Node{data: account}
	node.next = lst

	return node
}

func (lst *Node) IsEmpty() bool {
	return lst == nil
}

func (lst *Node) Find(account_number int) *Node {
	for p := lst; p != nil; p = p.next {
		if p.data.GetID() == account_number {
			return p
		}
	}

	return nil
}

func Free(lst **Node) {
	if *lst == nil {
		return
	}

	(*lst).data = nil

	Free(&(*lst).next)
	(*lst).next = nil
	*lst = nil
}

func (lst *Node) Remove(account_number int) *Node {
	var slow *Node

	p := lst
	for ; p != nil && p.data.GetID() != account_number; p = p.next {
		slow = p
	}

	if p == nil {
		return lst
	}
	if slow == nil {
		p.data = nil
		lst = p.next
	} else {
		slow.next = p.next
		p.data = nil
		p = nil
	}

	return lst
}

func (lst *Node) RemoveRec(account_number int) *Node {
	if !lst.IsEmpty() {
		if lst.data.GetID() == account_number {
			lst.data = nil
			lst = lst.next
		} else {
			lst.next = lst.next.RemoveRec(account_number)
		}
	}

	return lst
}

func (lst *Node) Print() {
	for p := lst; p != nil; p = p.next {
		fmt.Printf("** Account Number %d **\n", p.data.GetID())
		fmt.Printf("> Balance = %f \n", p.data.GetBalance())
	}
}

func (lst *Node) PrintRec() {
	if lst != nil {
		fmt.Printf("** Account Number %d **\n", lst.data.GetID())
		fmt.Printf("> Balance = %f \n", lst.data.GetBalance())
		lst.next.PrintRec()
	}
}

func (lst *Node) PrintRecInv() {
	if lst != nil {
		lst.next.PrintRecInv()
		fmt.Printf("** Account Number %d **\n", lst.data.GetID())
		fmt.Printf("> Balance = %f \n", lst.data.GetBalance())
	}
}

func (lst *Node) Credit(account_number int, amount float32) {
	account := lst.Find(account_number).data

	account.Credit(amount)
}

func (lst *Node) Debit(account_number int, amount float32) bool {
	account := lst.Find(account_number).data

	was_valid := account.Debit(amount)

	return was_valid
}

func (lst *Node) GetBalanceForAccount(account_number int) float32 {
	account := lst.Find(account_number).data

	balance := account.GetBalance()

	return balance
}

func (lst *Node) GetBonusForAccount(account_number int) (float32, bool) {
	account := lst.Find(account_number).data

	bonus, is_valid := account.GetBonus()

	return bonus, is_valid
}

func (lst *Node) TransferBetweenAccounts(amount float32, sender_number int, receiver_number int) bool {
	sender_account := lst.Find(sender_number).data
	receiver_account := lst.Find(receiver_number).data

	is_valid := sender_account.Debit(amount)

	if !is_valid {
		return false
	}

	receiver_account.Credit(amount)

	return true
}

func (lst *Node) YieldFromAccount(account_number int, fee float32) bool {
	account := lst.Find(account_number).data

	is_valid := account.Yield(fee)

	return is_valid
}
