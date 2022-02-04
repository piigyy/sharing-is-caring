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
