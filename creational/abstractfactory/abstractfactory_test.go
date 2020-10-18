package abstractfactory

import (
	"strings"
	"testing"
)

func TestBankFactoryMethod(t *testing.T) {
	factoryCreator := new(FactoryCreator)
	bankFactory, err := factoryCreator.GetFactory(BANK)

	if err != nil {
		t.Error("expected to return the BankFactory.")
	}

	bank, err := bankFactory.GetBank(HDFC)

	if err != nil {
		t.Fatal("expected to return the Bank.")
	}

	if strings.Compare(bank.GetShortBankName(), "HDFC") != 0 {
		t.Errorf("expected shortname is 'HDFC', but it has %s\n", bank.GetShortBankName())
	}
}

func TestLoanFactoryMethod(t *testing.T) {
	factoryCreator := new(FactoryCreator)
	loanFactory, err := factoryCreator.GetFactory(LOAN)

	if err != nil {
		t.Error("expected to return the LoanFactory.")
	}

	loan, err := loanFactory.GetLoan(HOME)

	if err != nil {
		t.Fatal("expected to return the Loan.")
	}

	if loan.GetInterestRate() != 10.4 {
		t.Errorf("expected interest rates for home loan is '10.6', but it has %0.2f\n", loan.GetInterestRate())
	}

}
