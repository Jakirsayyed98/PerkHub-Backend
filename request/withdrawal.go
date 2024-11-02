package request

type WithdrawalRequest struct {
	RequestedAmt string `json:"request_amt"`
	Upi          string `json:"upi_id"`
}

func NewWithdrawalRequest() *WithdrawalRequest {
	return &WithdrawalRequest{}
}
