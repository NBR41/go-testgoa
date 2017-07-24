package main

import (
	"fmt"
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/appmodel"
	"github.com/NBR41/go-testgoa/appsec"
	"github.com/goadesign/goa"
)

// ToAuthTokenMedia converts a user model and token into a auth token media type
func ToAuthTokenMedia(a *appmodel.User, token string) *app.Authtoken {
	return &app.Authtoken{
		User:  ToUserMedia(a),
		Token: token,
	}
}

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
	m, err := appmodel.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.GetAuthenticatedUser(ctx.Payload.Login, ctx.Payload.Password)
	if err != nil {
		if err == appmodel.ErrNotFound || err == appmodel.ErrInvalidCredentials {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	token, err := appsec.GetAuthToken(u.ID, u.IsAdmin)
	if err != nil {
		fmt.Println(err)
		return ctx.InternalServerError()
	}

	return ctx.OK(ToAuthTokenMedia(u, token))
	// AuthenticateController_Auth: end_implement
}
