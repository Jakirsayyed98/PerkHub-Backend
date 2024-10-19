package request

type LoginRequest struct {
	Number string `json:"number" binding:"required"`
	OTP    string `json:"otp"`
}

type SaveUserDetailReq struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	DOB    string `json:"dob"`
}

func NewSaveUserDetail() *SaveUserDetailReq {
	return &SaveUserDetailReq{}
}
