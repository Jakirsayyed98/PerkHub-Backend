package request

type CreateTicketRequest struct {
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

type GetTicketRequest struct {
	ID string `uri:"id" binding:"required"`
}

func NewCreateTicketRequest() *CreateTicketRequest {
	return &CreateTicketRequest{}
}
