package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Persistencies
type TransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCard) error
	GetCreditCard(creditCard CreditCard) (CreditCard, error)
	CreateCreditCard(creditCard CreditCard) error
}

type Transaction struct {
	ID           string
	Amount       float64
	Description  string
	Status       string
	Store        string
	CreditCardId string
	CreatedAt    time.Time
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now().UTC()
	return t
}

func (t *Transaction) ProcessAndValidate(creditCard *CreditCard) {
	if t.Amount+creditCard.Balance > creditCard.Limit {
		t.Status = "rejected"
	} else {
		t.Status = "approved"
		creditCard.Balance = creditCard.Balance + t.Amount
	}
}
