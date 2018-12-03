package security

import (
	"fmt"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func TestGetPasswordToken(t *testing.T) {
	tokenString, err := JWTHelper{}.GetPasswordToken(123, "foo@bar.com")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtPasswordKey, nil
	})
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["user_id"] != float64(123) {
			t.Errorf("unexpected value for user id, %v", claims["user_id"])
		}
		if claims["email"] != "foo@bar.com" {
			t.Errorf("unexpected value for user email, %v", claims["email"])
		}
		if claims["scope"] != ScopePassword {
			t.Errorf("unexpected value for scope, %v", claims["scope"])
		}
		if time.Unix(int64(claims["exp"].(float64)), 0).After(time.Now().Add(1 * time.Hour)) {
			t.Error("unexpected value for expire at")
		}
	}
}

func TestValidatePasswordToken(t *testing.T) {
	j := JWTHelper{}
	tokenString, err := j.GetPasswordToken(123, "foo@bar.com")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	id, email, err := j.ValidatePasswordToken(tokenString)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if id != 123 {
		t.Fatalf("unexpected id value, exp [%d] got [%d]", 123, id)
	}
	if email != "foo@bar.com" {
		t.Fatalf("unexpected email value, exp [%s] got [%s]", "foo@bar.com", email)
	}

	tokenString, err = j.getUserToken(123, "foo@bar.com", ScopeAccess, jwtPasswordKey)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	_, _, err = j.ValidatePasswordToken(tokenString)
	if err != ErrInvalidToken {
		t.Errorf("unexpected error, exp [%v] got [%v]", ErrInvalidToken, err)
	}
}

func TestGetValidationToken(t *testing.T) {
	tokenString, err := JWTHelper{}.GetValidationToken(123, "foo@bar.com")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtValidationKey, nil
	})
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["user_id"] != float64(123) {
			t.Errorf("unexpected value for user id, %v", claims["user_id"])
		}
		if claims["scope"] != ScopeValidation {
			t.Errorf("unexpected value for scope, %v", claims["scope"])
		}
		if time.Unix(int64(claims["exp"].(float64)), 0).After(time.Now().Add(1 * time.Hour)) {
			t.Error("unexpected value for expire at")
		}
	}
}

func TestValidateValidationToken(t *testing.T) {
	j := JWTHelper{}
	tokenString, err := j.GetValidationToken(123, "foo@bar.com")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	id, email, err := j.ValidateValidationToken(tokenString)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if id != 123 {
		t.Fatalf("unexpected id value, exp [%d] got [%d]", 123, id)
	}
	if email != "foo@bar.com" {
		t.Fatalf("unexpected email value, exp [%s] got [%s]", "foo@bar.com", email)
	}

	tokenString, err = j.getUserToken(123, "foo@bar.com", ScopeAccess, jwtValidationKey)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	_, _, err = j.ValidateValidationToken(tokenString)
	if err != ErrInvalidToken {
		t.Errorf("unexpected error, exp [%v] got [%v]", ErrInvalidToken, err)
	}
}

func TestGetAuthToken(t *testing.T) {
	tokenString, err := JWTHelper{}.GetAuthToken(123, true)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return JWTAuthKey, nil
	})
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["user_id"] != float64(123) {
			t.Errorf("unexpected value for user id, %v", claims["user_id"])
		}
		if !claims["is_admin"].(bool) {
			t.Error("unexpected value for is admin")
		}
		if claims["scope"] != ScopeAccess {
			t.Errorf("unexpected value for scope, %v", claims["scope"])
		}
		if time.Unix(int64(claims["exp"].(float64)), 0).After(time.Now().Add(72 * time.Hour)) {
			t.Error("unexpected value for expire at")
		}
	}
}

func TestGetRefreshToken(t *testing.T) {
	tokenString, err := JWTHelper{}.GetRefreshToken(123, true)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return JWTAuthKey, nil
	})
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["user_id"] != float64(123) {
			t.Errorf("unexpected value for user id, %v", claims["user_id"])
		}
		if !claims["is_admin"].(bool) {
			t.Error("unexpected value for is admin")
		}
		if claims["scope"] != ScopeRefresh {
			t.Errorf("unexpected value for scope, %v", claims["scope"])
		}
		if _, ok := claims["exp"]; ok {
			t.Error("unexpected value for expire at")
		}
	}
}

func TestValidateRefreshToken(t *testing.T) {
	j := JWTHelper{}
	tokenString, err := j.GetRefreshToken(123, true)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	id, err := j.ValidateRefreshToken(tokenString)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if id != 123 {
		t.Fatalf("unexpected id value, exp [%d] got [%d]", 123, id)
	}
	tokenString, err = j.getAuthToken(123, true, ScopeAccess, 0)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	_, err = j.ValidateRefreshToken(tokenString)
	if err != ErrInvalidToken {
		t.Errorf("unexpected error, exp [%v] got [%v]", ErrInvalidToken, err)
	}
}

func TestValidateTokenScope(t *testing.T) {
	_, err := JWTHelper{}.validateTokenScope("", "", nil)
	if err == nil {
		t.Fatal("expecting error")
	}
}

func TestValidateUserToken(t *testing.T) {
	_, _, err := JWTHelper{}.validateUserToken("", "", nil)
	if err == nil {
		t.Fatal("expecting error")
	}
}
