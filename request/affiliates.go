package request

import (
	"github.com/gin-gonic/gin"
)

type CreateAffiliateRequest struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Key    string `json:"key"`
	URL    string `json:"url"`
	Status bool   `json:"status"`
}

func NewCreateAffiliateRequest() *CreateAffiliateRequest {
	return &CreateAffiliateRequest{}
}

type CueLinkCallBackRequest struct {
	SubID         string `json:"subid"`
	SubID2        string `json:"subId2"`
	CampaignID    string `json:"campaign_id"`
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	OrderID       string `json:"order_id"`   // Corrected spelling
	Commission    string `json:"commission"` // Corrected spelling
	SaleAmount    string `json:"sale_amount"`

	ReferenceID string `json:"reference_id"`

	UserId               string `json:"user_id"`               // Corrected spelling
	CommissionPercentage string `json:"commission_percentage"` // Corrected spelling
	TransactionDate      string `json:"transaction_date"`
}

func NewCueLinkCallBackRequest() *CueLinkCallBackRequest {
	return &CueLinkCallBackRequest{}
}

func (req *CueLinkCallBackRequest) Bind(c *gin.Context) error {
	req.SubID = c.Query("subId")
	req.SubID2 = c.Query("subId2")
	req.CampaignID = c.Query("campaign_id")
	req.TransactionID = c.Query("transactionId")
	req.Status = c.Query("status")
	req.OrderID = c.Query("order_id")
	req.Commission = c.Query("commission")
	req.SaleAmount = c.Query("sale_amount")
	req.TransactionDate = c.Query("transaction_date")

	return nil
}
