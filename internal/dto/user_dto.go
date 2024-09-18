package dto

type CreateUserRequest struct {
	FullName string `json:"fullname" binding:"required"`
	// PhoneNumber string `json:"phone_number" binding:"required"`
	// Email       string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserRequest struct {
	ID int `uri:"id"`
}
