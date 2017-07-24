package appsec

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// JWTAuthKey key for auth
	JWTAuthKey = []byte("AllYourBase")

	// jwtPasswordKey key for password
	jwtPasswordKey = []byte("AllYourBase")

	// jwtValidationKey key for password
	jwtValidationKey = []byte("AllYourBase")

	scopePassword   = "password"
	scopeValidation = "validation"
)

// ErrInvalidToken error for invalid token
var ErrInvalidToken = errors.New("invalid token")

// UserClaims claims for user action
type UserClaims struct {
	jwt.StandardClaims
	UserID int64  `json:"user_id,omitempty"`
	Email  string `json:"email,omitempty"`
	Scope  string `json:"scope,omitempty"`
}

// AuthClaims jwt claims for authentification
type AuthClaims struct {
	jwt.StandardClaims
	UserID  int64 `json:"user_id,omitempty"`
	IsAdmin *bool `json:"is_admin,omitempty"`
}

func getToken(claims jwt.Claims, key interface{}) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
}

func getUserToken(userID int64, email, scope string, key interface{}) (string, error) {
	return getToken(
		UserClaims{
			UserID: userID,
			Email:  email,
			Scope:  scope,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour).Unix(),
				Issuer:    "test",
			},
		},
		key,
	)
}

func validateUserToken(rawtoken, scope string, key []byte) (int64, string, error) {
	token, err := jwt.ParseWithClaims(rawtoken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return 0, "", err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid && claims.Scope == scope {
		return claims.UserID, claims.Email, nil
	}
	return 0, "", ErrInvalidToken
}

// GetAuthToken returns new auth token
func GetAuthToken(userID int64, isAdmin bool) (string, error) {
	return getToken(
		AuthClaims{
			UserID:  userID,
			IsAdmin: &isAdmin,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
				Issuer:    "test",
			},
		},
		JWTAuthKey,
	)
}

// GetPasswordToken returns new password token
func GetPasswordToken(userID int64, email string) (string, error) {
	return getUserToken(userID, email, scopePassword, jwtPasswordKey)
}

// ValidatePasswordToken validates a password token and return the user ID
func ValidatePasswordToken(token string) (int64, string, error) {
	return validateUserToken(token, scopePassword, jwtPasswordKey)
}

// GetValidationToken returns new password token
func GetValidationToken(userID int64, email string) (string, error) {
	return getUserToken(userID, email, scopeValidation, jwtValidationKey)
}

// ValidateValidationToken validates a password token and return the user ID
func ValidateValidationToken(token string) (int64, string, error) {
	return validateUserToken(token, scopeValidation, jwtValidationKey)
}
