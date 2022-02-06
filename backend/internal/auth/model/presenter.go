package model

import "time"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID                   string    `json:"id"`
	AccessToken          string    `json:"accessToken"`
	Email                string    `json:"email"`
	AccessTokenExpiredAt time.Time `json:"accessTokenExpiredAt"`
}

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
	Email       string
}
