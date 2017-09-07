package main

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/appmodel"
	"github.com/NBR41/go-testgoa/appsec"
	"github.com/goadesign/goa"
)

// ToAuthTokenMedia converts a user model and token into a auth token media type
func ToAuthTokenMedia(a *appmodel.User, accToken, refToken string) *app.Authtoken {
	return &app.Authtoken{
		User:         ToUserMedia(a),
		AccessToken:  accToken,
		RefreshToken: refToken,
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
	m, err := appmodel.GetModeler()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.GetAuthenticatedUser(ctx.Payload.Login, ctx.Payload.Password)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to auth`, `error`, err)
		if err == appmodel.ErrNotFound || err == appmodel.ErrInvalidCredentials {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	accToken, err := appsec.GetAuthToken(u.ID, u.IsAdmin)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get access token`, `error`, err)
		return ctx.InternalServerError()
	}

	refToken, err := appsec.GetRefreshToken(u.ID, u.IsAdmin)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get refresh token`, `error`, err)
		return ctx.InternalServerError()
	}

	return ctx.OK(ToAuthTokenMedia(u, accToken, refToken))
	// AuthenticateController_Auth: end_implement
}
