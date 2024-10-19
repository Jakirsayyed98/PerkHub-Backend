package request

import (
	"github.com/gin-gonic/gin"
)

type CueLinkCallBackRequest struct {
	CampaignID           string `json:"campaign_id"`
	Commission           string `json:"commission"` // Corrected spelling
	ReferenceID          string `json:"reference_id"`
	SaleAmount           string `json:"sale_amount"`
	Status               string `json:"status"`
	SubID                string `json:"subid"`
	SubID1               string `json:"subid1"`
	SubID2               string `json:"subid2"`                // Corrected spelling
	SubID3               string `json:"subid3"`                // Corrected spelling
	CommissionPercentage string `json:"commission_percentage"` // Corrected spelling
	TransactionDate      string `json:"transaction_date"`
	TransactionID        string `json:"transaction_id"`
}

func NewCueLinkCallBackRequest() *CueLinkCallBackRequest {
	return &CueLinkCallBackRequest{}
}

func (req *CueLinkCallBackRequest) Bind(c *gin.Context) error {
	req.CampaignID = c.Query("campaign_id")
	req.ReferenceID = c.Query("reference_id")
	req.Status = c.Query("status")
	req.SubID = c.Query("subid")
	req.SubID1 = c.Query("subid1")
	req.SubID2 = c.Query("subid2")
	req.SubID3 = c.Query("subid3")
	req.TransactionDate = c.Query("transaction_date")
	req.TransactionID = c.Query("transaction_id")
	req.Commission = c.Query("commission")
	req.SaleAmount = c.Query("sale_amount")

	return nil
}
