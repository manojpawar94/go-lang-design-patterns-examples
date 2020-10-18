package abstractfactory

import (
	"errors"
)

//factory of factory is known as abstract factory
//Bank constants
const (
	HDFC = 1
	SBI  = 2
)

//BankFactory struct
type BankFactory struct{}

//GetBank method
func (*BankFactory) GetBank(bank int) (Bank, error) {
	switch bank {
	case HDFC:
		return new(HDFCBank), nil
	case SBI:
		return new(SBIBank), nil
	default:
		return nil, errors.New("Banks has not found")
	}
}

//GetLoan method
func (*BankFactory) GetLoan(loanType int) (ILoan, error) {
	return nil, errors.New("No supported by BankFactory")
}

//Bank interface
type Bank interface {
	GetShortBankName() string
	GetBankName() string
}

//HDFCBank struct
type HDFCBank struct{}

//GetShortBankName method
func (*HDFCBank) GetShortBankName() string {
	return "HDFC"
}

//GetBankName method
func (*HDFCBank) GetBankName() string {
	return "Housing Development Finance Corporation Limited"
}

//SBIBank struct
type SBIBank struct{}

//GetShortBankName method
func (*SBIBank) GetShortBankName() string {
	return "SBI"
}

//GetBankName method
func (*SBIBank) GetBankName() string {
	return "State Bank of India"
}
