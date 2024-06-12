package model

type UserAuth struct {
	UserID   uint `json:"user_id"`
	Email    uint `json:"email"`
	Password uint `json:"password"`
}
