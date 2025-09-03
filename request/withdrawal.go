package request

type WithdrawalRequest struct {
	RequestedAmt  string `json:"request_amt"`
	Upi           string `json:"upi_id"`
	PaymentMethod string `json:"payment_method"`
}

func NewWithdrawalRequest() *WithdrawalRequest {
	return &WithdrawalRequest{}
}

type AddPaymentMethodRequest struct {
	Upi               string `json:"upi_id"`
	BankAccountNumber string `json:"bank_account_number"`
	IFSCCode          string `json:"ifsc_code"`
	BankName          string `json:"bank_name"`
}

func NewAddPaymentMethodRequest() *AddPaymentMethodRequest {
	return &AddPaymentMethodRequest{}
}
