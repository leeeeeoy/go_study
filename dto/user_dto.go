package dto

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}

type UserResponse struct {
	ID        int       `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type SignInRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type SignInResponse struct {
	AccessToken string `json:"access_token,omitempty"`
}

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}
