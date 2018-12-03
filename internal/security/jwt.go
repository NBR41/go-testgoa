package security

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
)

//List of constants
const (
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

//JWTHelper struct for jwt token helper
type JWTHelper struct {
}

// GetPasswordToken returns new password token
func (j JWTHelper) GetPasswordToken(userID int64, email string) (string, error) {
	return j.getUserToken(userID, email, ScopePassword, jwtPasswordKey)
}

// ValidatePasswordToken validates a password token and return the user ID
func (j JWTHelper) ValidatePasswordToken(token string) (int64, string, error) {
	return j.validateUserToken(token, ScopePassword, jwtPasswordKey)
}

// GetValidationToken returns new password token
func (j JWTHelper) GetValidationToken(userID int64, email string) (string, error) {
	return j.getUserToken(userID, email, ScopeValidation, jwtValidationKey)
}

// ValidateValidationToken validates a password token and return the user ID
func (j JWTHelper) ValidateValidationToken(token string) (int64, string, error) {
	return j.validateUserToken(token, ScopeValidation, jwtValidationKey)
}

// GetAuthToken returns new auth token
func (j JWTHelper) GetAuthToken(userID int64, isAdmin bool) (string, error) {
	return j.getAuthToken(userID, isAdmin, ScopeAccess, time.Now().Add(time.Hour*72).Unix())
}

// GetRefreshToken returns new refresh token
func (j JWTHelper) GetRefreshToken(userID int64, isAdmin bool) (string, error) {
	return j.getAuthToken(userID, isAdmin, ScopeRefresh, 0)
}

// ValidateRefreshToken returns userID if it's a valid refresh token
func (j JWTHelper) ValidateRefreshToken(token string) (int64, error) {
	return j.validateTokenScope(token, ScopeRefresh, JWTAuthKey)
}

func (j JWTHelper) getToken(claims jwt.Claims, key interface{}) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
}

func (j JWTHelper) getUserToken(userID int64, email, scope string, key interface{}) (string, error) {
	return j.getToken(j.getUserClaims(userID, email, scope), key)
}

func (j JWTHelper) validateUserToken(rawtoken, scope string, key []byte) (int64, string, error) {
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

func (j JWTHelper) validateTokenScope(rawtoken, scope string, key []byte) (int64, error) {
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

func (j JWTHelper) getAuthToken(userID int64, isAdmin bool, scope string, expiresat int64) (string, error) {
	return j.getToken(j.getAuthClaims(userID, isAdmin, scope, expiresat), JWTAuthKey)
}

func (j JWTHelper) getUserClaims(userID int64, email, scope string) UserClaims {
	return UserClaims{
		UserID: userID,
		Email:  email,
		Scope:  scope,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "test",
		},
	}
}

func (j JWTHelper) getAuthClaims(userID int64, isAdmin bool, scope string, expiresat int64) AuthClaims {
	return AuthClaims{
		UserID:         userID,
		IsAdmin:        &isAdmin,
		Scope:          scope,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expiresat, Issuer: "test"},
	}
}
