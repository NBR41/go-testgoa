package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// AuthenticateController implements the authenticate resource.
type AuthenticateController struct {
	*goa.Controller
	fm  Fmodeler
	tok TokenHelper
}

// NewAuthenticateController creates a authenticate controller.
func NewAuthenticateController(service *goa.Service, fm Fmodeler, tok TokenHelper) *AuthenticateController {
	return &AuthenticateController{
		Controller: service.NewController("AuthenticateController"),
		fm:         fm,
		tok:        tok,
	}
}

// Auth runs the auth action.
func (c *AuthenticateController) Auth(ctx *app.AuthAuthenticateContext) error {
	// AuthenticateController_Auth: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.GetAuthenticatedUser(ctx.Payload.Login, ctx.Payload.Password)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to auth`, `error`, err.Error())
		if err == model.ErrNotFound || err == model.ErrInvalidCredentials {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	accToken, err := c.tok.GetAccessToken(u.ID, u.IsAdmin)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get access token`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	refToken, err := c.tok.GetRefreshToken(u.ID, u.IsAdmin)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get refresh token`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToAuthTokenMedia(u, accToken, refToken))
	// AuthenticateController_Auth: end_implement
}
