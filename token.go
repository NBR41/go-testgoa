package main

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/appmodel"
	"github.com/NBR41/go-testgoa/appsec"
	"github.com/goadesign/goa"
)

// TokenController implements the token resource.
type TokenController struct {
	*goa.Controller
}

// NewTokenController creates a token controller.
func NewTokenController(service *goa.Service) *TokenController {
	return &TokenController{Controller: service.NewController("TokenController")}
}

// Access runs the access action.
func (c *TokenController) Access(ctx *app.AccessTokenContext) error {
	// TokenController_Access: start_implement
	authToken, err := appsec.GetAuthToken(ctx.Value(ctxKey("user_id")).(int64), ctx.Value(ctxKey("is_admin")).(bool))
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get password token`, `error`, err)
		return ctx.InternalServerError()
	}
	return ctx.OK(&app.Token{Token: authToken})
	// TokenController_Access: end_implement
}

// Auth runs the auth action.
func (c *TokenController) Auth(ctx *app.AuthTokenContext) error {
	// TokenController_Auth: start_implement
	m, err := appmodel.GetModeler()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.GetUserByID(int(ctx.Value(ctxKey("user_id")).(int64)))
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get user`, `error`, err)
		if err == appmodel.ErrNotFound {
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
	// TokenController_Auth: end_implement
}
