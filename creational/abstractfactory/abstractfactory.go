package abstractfactory

import "errors"

//constant
const (
	BANK = "BANK"
	LOAN = "LOAN"
)

//AbstractFactory struct
type AbstractFactory interface {
	GetBank(bank int) (Bank, error)
	GetLoan(loanType int) (ILoan, error)
}

//FactoryCreator struct
type FactoryCreator struct{}

//GetFactory method
func (*FactoryCreator) GetFactory(factoryType string) (AbstractFactory, error) {
	switch factoryType {
	case BANK:
		return new(BankFactory), nil
	case LOAN:
		return new(LoanFactory), nil
	default:
		return nil, errors.New("No implementation yet")
	}
}
