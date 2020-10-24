package adapter

//BankDetails struct
type BankDetails struct {
	BankDetails []BankDetail `json:"bankDetails"`
}

//BankDetail struct - an adaptee class
type BankDetail struct {
	BankName          string `json:"bankName"`
	AccountHolderName string `json:"accountHolderName"`
	AccountNumber     int64  `json:"accountNumber"`
}

//CreditCard interface
type CreditCard interface {
	CheckEligibility(BankDetail)
}

//BankCustomer struct - an adpater class
type BankCustomer struct {
	BankDetail
}

//CheckEligibility method
func (b *BankCustomer) CheckEligibility(bankDetail BankDetail) (bool, string) {
	if b.BankName == bankDetail.BankName && b.AccountHolderName == bankDetail.AccountHolderName && b.AccountNumber == bankDetail.AccountNumber {
		return true, "Bank details Authenticated. Customer is eligiable for Credit Card."
	}
	return false, "Bank details did not Authenticate. Customer is not eligiable for Credit Card."
}
