package het_linked_list

import (
	"testing"

	Bank "github.com/soaresrl/dsa/pkg/bank"
)

func TestCreateList(t *testing.T) {
	lst := NewList()

	if lst != nil {
		t.Errorf("Node should be nil, got %v", lst)
	}

}

func TestAdd(t *testing.T) {
	lst := NewList()

	lst = lst.Insert(Bank.Bank)
	lst = lst.Insert(Bank.Saving)
	lst = lst.Insert(Bank.Loyalty)

	if lst.data.GetID() != 3 {
		t.Errorf("Insert test failed got %d, wanted %d", lst.data.GetID(), 2)
	}

	if lst.next.data.GetID() != 2 {
		t.Errorf("Add test failed got (next) %d, wanted %d", lst.next.data.GetID(), 1)
	}
}

func TestRemove(t *testing.T) {
	lst := NewList()

	lst = lst.Insert(Bank.Bank)
	lst = lst.Insert(Bank.Saving)
	lst = lst.Insert(Bank.Loyalty)

	lst = lst.Remove(2)

	if lst.data.GetID() != 3 {
		t.Errorf("Remove test failed expected head = %d, got %d", 3, lst.data)
	}

	if lst.next.data.GetID() != 1 {
		t.Errorf("Remove test failed expected next = %d, got %d", 1, lst.next.data)
	}
}

func TestCredit(t *testing.T) {
	lst := NewList()

	lst = lst.Insert(Bank.Bank)
	lst = lst.Insert(Bank.Saving)
	lst = lst.Insert(Bank.Loyalty)

	lst.Credit(1, 50.0)

	if lst.Find(1).data.GetBalance() != 50.0 {
		t.Errorf("Credit operation failed, expected %f got %f", 50.0, lst.Find(1).data.GetBalance())
	}
}

func TestDebit(t *testing.T) {
	lst := NewList()

	lst = lst.Insert(Bank.Bank)
	lst = lst.Insert(Bank.Saving)
	lst = lst.Insert(Bank.Loyalty)

	lst.Credit(1, 50.0)

	is_valid := lst.Debit(1, 25.0)

	if is_valid != true {
		t.Errorf("Debit operation failed, expected valid to be %v got %v", true, is_valid)
	}

	if lst.Find(1).data.GetBalance() != 25.0 {
		t.Errorf("Debit operation failed, expected %f got %f", 25.0, lst.Find(1).data.GetBalance())
	}

	is_valid = lst.Debit(1, 30.0)

	if is_valid != false {
		t.Errorf("Debit operation failed, expected valid to be %v got %v", false, is_valid)
	}

	if lst.Find(1).data.GetBalance() != 25.0 {
		t.Errorf("Debit operation failed, expected %f got %f", 25.0, lst.Find(1).data.GetBalance())
	}
}

func TestGetBonus(t *testing.T) {
	lst := NewList()

	lst = lst.Insert(Bank.Bank)
	lst = lst.Insert(Bank.Saving)
	lst = lst.Insert(Bank.Loyalty)

	lst.Credit(3, 100.0)

	bonus, is_valid := lst.GetBonusForAccount(3)

	if is_valid != true {
		t.Errorf("Get Bonus operation failed, expected to valid to be %v got %v", true, is_valid)
	}

	if bonus != 1.0 {
		t.Errorf("Bonus failed, expected %f got %f.", 1.0, bonus)
	}
}

func TestTransferBetweenAccounts(t *testing.T) {
	lst := NewList()

	lst = lst.Insert(Bank.Bank)
	lst = lst.Insert(Bank.Saving)
	lst = lst.Insert(Bank.Loyalty)

	lst.Credit(1, 1000.0)

	is_valid := lst.TransferBetweenAccounts(100.0, 1, 2)

	if is_valid != true {
		t.Errorf("Trasnfer failed expected valid to be %v got %v", true, is_valid)
	}

	amount_sender := lst.GetBalanceForAccount(1)
	amount_receiver := lst.GetBalanceForAccount(2)

	if amount_sender != 900.0 {
		t.Errorf("Transfer failed expected sender amount to be %f got %f", 900.0, amount_sender)
	}

	if amount_receiver != 100.0 {
		t.Errorf("Transfer failed expected receiver amount to be %f got %f", 100.0, amount_receiver)
	}

	is_valid = lst.TransferBetweenAccounts(1000.0, 1, 3)

	if is_valid != false {
		t.Errorf("Trasnfer failed expected valid to be %v got %v", false, is_valid)
	}
}

func TestYield(t *testing.T) {
	lst := NewList()

	lst = lst.Insert(Bank.Bank)
	lst = lst.Insert(Bank.Saving)
	lst = lst.Insert(Bank.Loyalty)

	lst.Credit(2, 1000.0)
	lst.Credit(3, 1000.0)

	is_valid := lst.YieldFromAccount(2, 2.0)

	if is_valid != true {
		t.Errorf("Yield operation failed expected valid to be %v got %v", true, is_valid)
	}

	balance := lst.GetBalanceForAccount(2)

	if balance != 1020.0 {
		t.Errorf("Yield operation test failed expected balance = %f, got %f.\n", 1020.0, balance)
	}

	is_valid = lst.YieldFromAccount(3, 0.0)

	if is_valid != true {
		t.Errorf("Yield operation test failed expected valid to be %v got %v", true, is_valid)
	}

	balance = lst.GetBalanceForAccount(3)

	if balance != 1010.0 {
		t.Errorf("Yield operation test failed expected balance = %f, got %f.\n", 1010.0, balance)
	}
}
