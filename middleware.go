package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/NBR41/go-testgoa/appsec"
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

		// check we have required claims fields
		if _, ok := claims["is_admin"]; !ok {
			return goajwt.ErrJWTError("unauthorized")
		}
		if _, ok := claims["user_id"]; !ok {
			return goajwt.ErrJWTError("unauthorized")
		}
		if _, ok := claims["scope"]; !ok {
			return goajwt.ErrJWTError("unauthorized")
		}

		// check scopes
		if req.URL != nil && (req.URL.Path == "/token/auth" || req.URL.Path == "/token/access") {
			if claims["scope"] != appsec.ScopeRefresh {
				return goajwt.ErrJWTError("unauthorized")
			}
		} else {
			if claims["scope"] != appsec.ScopeAccess {
				return goajwt.ErrJWTError("unauthorized")
			}
		}

		// store values from claims
		ctx = context.WithValue(ctx, ctxKey("is_admin"), claims["is_admin"])
		ctx = context.WithValue(ctx, ctxKey("user_id"), int64(claims["user_id"].(float64)))

		if claims["is_admin"] == true {
			return h(ctx, rw, req)
		}

		// extra check on user if there is an user id in path
		r := goa.ContextRequest(ctx)
		if len(r.Params["user_id"]) == 1 && r.Params["user_id"][0] == strconv.FormatFloat(claims["user_id"].(float64), 'f', 0, 64) {
			return h(ctx, rw, req)
		}
		return goajwt.ErrJWTError("unauthorized")
	}
}
