package main

import (
	"context"
	"net/http"

	goajwt "github.com/goadesign/goa/middleware/security/jwt"

	"github.com/NBR41/go-testgoa/appsec"
)

var jwtUserValiadtion = func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	token := goajwt.ContextJWT(ctx)

	claims, ok := token.Claims.(appsec.AuthClaims)
	if !ok {
		return goajwt.ErrJWTError("wrong type of claims")
	}
	if claims.IsAdmin == nil || !*claims.IsAdmin {
		return goajwt.ErrJWTError("you are not uncle ben's")
	}
	return nil
}
