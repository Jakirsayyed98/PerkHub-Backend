package request

type LoginRequest struct {
	Number string `json:"number" binding:"required"`
	OTP    string `json:"otp"`
}

type AdminLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminRegister struct {
	Email    string `json:"email" binding:"required,email"`    // "email" should also be a valid email format
	Password string `json:"password" binding:"required,min=6"` // Password must be at least 6 characters long
	Name     string `json:"name" binding:"required,min=3"`     // Name should be at least 3 characters long
}

type SaveUserDetailReq struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	DOB    string `json:"dob"`
}

func NewSaveUserDetail() *SaveUserDetailReq {
	return &SaveUserDetailReq{}
}
