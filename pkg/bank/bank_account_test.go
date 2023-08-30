package bank

import (
	"testing"
)

func BankAccountCreditTest(t *testing.T) {
	account := BankAccount{
		balance: 1000,
		number:  1,
	}

	account.Credit(50)

	if account.balance != 1050 {
		t.Errorf("Credit test failed. expected %v got %v.\n", 1050, account.balance)
	}
}
