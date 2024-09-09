package dto

type CreateUserRequest struct {
	FullName    string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type GetUserRequest struct {
	ID int `uri:"id"`
}
