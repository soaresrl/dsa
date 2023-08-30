package bank

const (
	Bank    = 0
	Saving  = 1
	Loyalty = 2
)

type Accountable interface {
	Credit(float32)
	Debit(amount float32) bool
	Yield(fee float32) bool
	GetBalance() float32
	GetBonus() (float32, bool)
	GetID() int
}

type BankAccount struct {
	Number  int
	balance float32
}

type SavingAccount struct {
	Number  int
	balance float32
}

type LoyaltyAccount struct {
	Number  int
	balance float32
	bonus   float32
}

func (b *BankAccount) Credit(amount float32) {
	b.balance += amount
}

func (b *BankAccount) Debit(amount float32) bool {
	if b.balance-amount < 0.0 {
		return false
	}

	b.balance -= amount

	return true
}

func (b *BankAccount) Yield(fee float32) bool {
	return false
}

func (b *BankAccount) GetID() int {
	return b.Number
}

func (b *BankAccount) GetBalance() float32 {
	return b.balance
}

func (b *BankAccount) GetBonus() (float32, bool) {
	return 0.0, false
}

func (b *SavingAccount) Credit(amount float32) {
	b.balance += amount
}

func (b *SavingAccount) Debit(amount float32) bool {
	if b.balance-amount < 0.0 {
		return false
	}

	b.balance -= amount

	return true
}

func (b *SavingAccount) Yield(fee float32) bool {
	b.balance += fee * b.balance / 100

	return true
}

func (b *SavingAccount) GetID() int {
	return b.Number
}

func (b *SavingAccount) GetBalance() float32 {
	return b.balance
}

func (b *SavingAccount) GetBonus() (float32, bool) {
	return 0.0, false
}

func (b *LoyaltyAccount) Credit(amount float32) {
	b.balance += amount
	b.bonus += 0.01 * amount
}

func (b *LoyaltyAccount) Debit(amount float32) bool {
	if b.balance-amount < 0.0 {
		return false
	}

	b.balance -= amount
	return true
}

func (b *LoyaltyAccount) Yield(fee float32) bool {
	b.balance += b.bonus
	b.bonus = 0.0

	return true
}

func (b *LoyaltyAccount) GetID() int {
	return b.Number
}

func (b *LoyaltyAccount) GetBalance() float32 {
	return b.balance
}

func (b *LoyaltyAccount) GetBonus() (float32, bool) {
	return b.bonus, true
}
