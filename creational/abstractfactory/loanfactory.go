package abstractfactory

import "errors"

//Loan type constant
const (
	HOME     = 1
	BUSINESS = 2
	EDUCTION = 3
)

//LoanFactory struct
type LoanFactory struct{}

//GetLoan method
func (*LoanFactory) GetLoan(loanType int) (ILoan, error) {
	switch loanType {
	case HOME:
		return &HomeLoan{Loan: Loan{10.4}}, nil
	case BUSINESS:
		return &BusinessLoan{Loan: Loan{15.6}}, nil
	case EDUCTION:
		return &EductionLoan{Loan: Loan{11.22}}, nil
	default:
		return nil, errors.New("Loan type not implemented yet")
	}
}

//GetBank method
func (*LoanFactory) GetBank(bank int) (Bank, error) {
	return nil, errors.New("Not supported by loan factory")
}

//ILoan interface an abstract interface
type ILoan interface {
	GetInterestRate() float32
	SetInterestRate(rate float32)
	CalculatLoanPayment(loanAmount float64, year int) float64
}

//Loan struct an abstract class
type Loan struct {
	interestRate float32
}

//CalculatLoanPayment method
func (l *Loan) CalculatLoanPayment(loanAmount float64, year int) float64 {
	interestAmount := loanAmount * float64(l.interestRate) * float64(year)
	return loanAmount + interestAmount
}

//HomeLoan struct an implementing class
type HomeLoan struct {
	Loan
}

//GetInterestRate method
func (l *HomeLoan) GetInterestRate() float32 {
	return l.Loan.interestRate
}

//SetInterestRate method
func (l *HomeLoan) SetInterestRate(interestRate float32) {
	l.Loan.interestRate = interestRate
}

//BusinessLoan struct an implementing class
type BusinessLoan struct {
	Loan
}

//GetInterestRate method
func (l *BusinessLoan) GetInterestRate() float32 {
	return l.Loan.interestRate
}

//SetInterestRate method
func (l *BusinessLoan) SetInterestRate(interestRate float32) {
	l.Loan.interestRate = interestRate
}

//EductionLoan struct an implementing class
type EductionLoan struct {
	Loan
}

//GetInterestRate method
func (l *EductionLoan) GetInterestRate() float32 {
	return l.Loan.interestRate
}

//SetInterestRate method
func (l *EductionLoan) SetInterestRate(interestRate float32) {
	l.Loan.interestRate = interestRate
}
