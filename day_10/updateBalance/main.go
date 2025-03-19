package updatebalance

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line
func (c *customer) deposit(t transaction) {
	c.balance += t.amount
}
func (c *customer) withdrawl(t transaction) error {
	if c.balance < t.amount {
		return errors.New("insufficient funds")
	}
	c.balance -= t.amount
	return nil
}
func updateBalance(c *customer, t transaction) error {
	switch t.transactionType {
	case transactionDeposit:
		c.deposit(t)
		return nil
	case transactionWithdrawal:
		err := c.withdrawl(t)
		if err != nil {
			return err

		}

	default:
		return errors.New("unknown transaction type")
	}
	return nil

}
