package adapter

import (
	"strings"
	"testing"
)

func TestBankCustomerAdapeter(t *testing.T) {
	bankDetails := ReadBankDetailsFromJSON()

	bankCustomer := &BankCustomer{BankDetail: bankDetails.BankDetails[0]}

	status, msg := bankCustomer.CheckEligibility(bankDetails.BankDetails[0])
	if !status {
		t.Error("status is expected to true, but actual is false")
	}

	if !strings.Contains(msg, "Customer is eligiable for Credit Card") {
		t.Error("Msg must contain 'Customer is eligiable for Credit Card'.")
	}

	status, msg = bankCustomer.CheckEligibility(bankDetails.BankDetails[1])
	if status {
		t.Error("status is expected to false, but actual is true")
	}

	if !strings.Contains(msg, "Customer is not eligiable for Credit Card") {
		t.Error("Msg must contain 'Customer is not eligiable for Credit Card'.")
	}
}
