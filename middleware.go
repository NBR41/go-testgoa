package main

import (
	"context"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	goajwt "github.com/goadesign/goa/middleware/security/jwt"
)

type ctxKey string

func jwtUserValiadtion(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		token := goajwt.ContextJWT(ctx)
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return goajwt.ErrJWTError("wrong type of claims")
		}

		if claims["is_admin"] == true {
			ctx = context.WithValue(ctx, ctxKey("is_admin"), true)
			return h(ctx, rw, req)
		}
		r := goa.ContextRequest(ctx)
		if len(r.Params["user_id"]) == 1 && r.Params["user_id"][0] == strconv.FormatFloat(claims["user_id"].(float64), 'f', 0, 64) {
			return h(ctx, rw, req)
		}
		return goajwt.ErrJWTError("unauthorized")
	}
}
