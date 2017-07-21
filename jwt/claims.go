package jwt

import (
	jwtgo "github.com/dgrijalva/jwt-go"
)

type AppClaims struct {
	jwtgo.StandardClaims
	UserID  int64 `json:"user_id,omitempty"`
	IsAdmin *bool `json:"is_admin,omitempty"`
}
