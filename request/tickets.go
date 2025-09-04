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

type SendTicketMsg struct {
	TicketId string `json:"ticketId" binding:"required"`
	Message  string `json:"msg" binding:"required"`
}

func NewSendTicketMsg() *SendTicketMsg {
	return &SendTicketMsg{}
}

type AdminReplyTicketMsg struct {
	TicketId string `json:"ticketId" binding:"required"`
	Message  string `json:"msg" binding:"required"`
}

func NewAdminReplyTicketMsg() *AdminReplyTicketMsg {
	return &AdminReplyTicketMsg{}
}
