package main

import (
	"github.com/NBR41/go-testgoa/app"
	appjwt "github.com/NBR41/go-testgoa/jwt"
	"github.com/NBR41/go-testgoa/store"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"time"
)

// AuthenticateController implements the authenticate resource.
type AuthenticateController struct {
	*goa.Controller
}

// NewAuthenticateController creates a authenticate controller.
func NewAuthenticateController(service *goa.Service) *AuthenticateController {
	return &AuthenticateController{Controller: service.NewController("AuthenticateController")}
}

// Auth runs the auth action.
func (c *AuthenticateController) Auth(ctx *app.AuthAuthenticateContext) error {
	// AuthenticateController_Auth: start_implement

	// Put your logic here
	m, err := store.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.GetAuthenticatedUser(ctx.Payload.Login, ctx.Payload.Password)
	if err != nil {
		if err == store.ErrNotFound || err == store.ErrInvalidCredentials {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	// Create the Claims
	claims := appjwt.AppClaims{
		UserID:  u.ID,
		IsAdmin: &u.IsAdmin,
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "test",
		},
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	ss, err := token.SignedString("AllYourBase")
	if err != nil {
		return ctx.InternalServerError()
	}

	return ctx.OK(&app.Token{Token: ss})
	// AuthenticateController_Auth: end_implement
}
