package factory

import (
	"errors"
	"fmt"
)

//PaymentMethod interface
type PaymentMethod interface {
	Pay(amount float32) string
}

//Payment method constants
const (
	Cash       = 1
	DebitCard  = 2
	CreditCard = 3
	UPI        = 4
)

//GetPaymentMethod function
func GetPaymentMethod(codePM int) (PaymentMethod, error) {
	switch codePM {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	case CreditCard:
		return new(CreditCardPM), nil
	case UPI:
		return new(UIPIPM), nil
	default:
		return nil, errors.New("No implementation yet")
	}
}

//CashPM struct
type CashPM struct{}

//Pay method
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using Cash payment\n", amount)
}

//DebitCardPM struct
type DebitCardPM struct{}

//Pay method
func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using Debit Card payment\n", amount)
}

//CreditCardPM struct
type CreditCardPM struct{}

//Pay method
func (c *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using Credit Card payment\n", amount)
}

//UIPIPM struct
type UIPIPM struct{}

//Pay method
func (c *UIPIPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using UPI payment\n", amount)
}
