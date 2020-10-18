package factory

import (
	"errors"
	"strings"
	"testing"
)

func TestCashPaymentMethod(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)

	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist.")
	}

	msg := payment.Pay(10.25)

	if !strings.Contains(msg, "paid using Cash") {
		t.Error("Payment has not done with cash payment.")
	}
}

func TestDebitCardPaymentMethod(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)

	if err != nil {
		t.Fatal("A payment method of type 'Debit Card' must exist.")
	}

	msg := payment.Pay(10.25)

	if !strings.Contains(msg, "paid using Debit Card") {
		t.Error("Payment has not done with Debit Card payment.")
	}
}

func TestCreditCardPaymentMethod(t *testing.T) {
	payment, err := GetPaymentMethod(CreditCard)

	if err != nil {
		t.Fatal("A payment method of type 'Credit Card' must exist.")
	}

	msg := payment.Pay(10.25)

	if !strings.Contains(msg, "paid using Credit Card") {
		t.Error("Payment has not done with Credit Card payment.")
	}
}

func TestUPIPaymentMethod(t *testing.T) {
	payment, err := GetPaymentMethod(UPI)

	if err != nil {
		t.Fatal("A payment method of type 'UPI' must exist.")
	}

	msg := payment.Pay(10.25)

	if !strings.Contains(msg, "paid using UPI") {
		t.Error("Payment has not done with UPI payment.")
	}
}

func TestNotExistingPaymentMethod(t *testing.T) {
	_, err := GetPaymentMethod(5)

	if err == nil {
		t.Fatal("expected to throw error")
	} else if err == errors.New("No implementation yet") {
		t.Error("The expected error of type No implementation yet")
	}
}
