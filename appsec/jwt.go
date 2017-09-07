package appsec

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// JWTAuthKey key for auth
	JWTAuthKey = []byte("authkey")

	// jwtPasswordKey key for password
	jwtPasswordKey = []byte("passwordkey")

	// jwtValidationKey key for password
	jwtValidationKey = []byte("validationkey")

	ScopeAccess     = "access"
	ScopeRefresh    = "refresh"
	ScopePassword   = "password"
	ScopeValidation = "validation"
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
	UserID  int64  `json:"user_id,omitempty"`
	IsAdmin *bool  `json:"is_admin,omitempty"`
	Scope   string `json:"scope,omitempty"`
}

// GetPasswordToken returns new password token
func GetPasswordToken(userID int64, email string) (string, error) {
	return getUserToken(userID, email, ScopePassword, jwtPasswordKey)
}

// ValidatePasswordToken validates a password token and return the user ID
func ValidatePasswordToken(token string) (int64, string, error) {
	return validateUserToken(token, ScopePassword, jwtPasswordKey)
}

// GetValidationToken returns new password token
func GetValidationToken(userID int64, email string) (string, error) {
	return getUserToken(userID, email, ScopeValidation, jwtValidationKey)
}

// ValidateValidationToken validates a password token and return the user ID
func ValidateValidationToken(token string) (int64, string, error) {
	return validateUserToken(token, ScopeValidation, jwtValidationKey)
}

func getToken(claims jwt.Claims, key interface{}) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
}

func getUserToken(userID int64, email, scope string, key interface{}) (string, error) {
	return getToken(
		UserClaims{
			UserID: userID,
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
	return getAuthToken(userID, isAdmin, ScopeAccess, time.Now().Add(time.Hour*72).Unix())
}

// GetRefreshToken returns new refresh token
func GetRefreshToken(userID int64, isAdmin bool) (string, error) {
	return getAuthToken(userID, isAdmin, ScopeRefresh, 0)
}

// ValidateRefreshToken returns userID if it's a valid refresh token
func ValidateRefreshToken(token string) (int64, error) {
	return validateTokenScope(token, ScopeRefresh, JWTAuthKey)
}

func validateTokenScope(rawtoken, scope string, key []byte) (int64, error) {
	token, err := jwt.ParseWithClaims(rawtoken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid && claims.Scope == scope {
		return claims.UserID, nil
	}
	return 0, ErrInvalidToken
}

func getAuthToken(userID int64, isAdmin bool, scope string, expiresat int64) (string, error) {
	return getToken(
		AuthClaims{
			UserID:         userID,
			IsAdmin:        &isAdmin,
			Scope:          scope,
			StandardClaims: jwt.StandardClaims{ExpiresAt: expiresat, Issuer: "test"},
		},
		JWTAuthKey,
	)
}
