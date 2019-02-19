package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// TokenController implements the token resource.
type TokenController struct {
	*goa.Controller
	fm  Fmodeler
	tok TokenHelper
}

// NewTokenController creates a token controller.
func NewTokenController(service *goa.Service, fm Fmodeler, tok TokenHelper) *TokenController {
	return &TokenController{
		Controller: service.NewController("TokenController"),
		fm:         fm,
		tok:        tok,
	}
}

// Access runs the access action.
func (c *TokenController) Access(ctx *app.AccessTokenContext) error {
	// TokenController_Access: start_implement
	authToken, err := c.tok.GetAccessToken(ctx.Value(CtxKey("user_id")).(int64), ctx.Value(CtxKey("is_admin")).(bool))
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get access token`, `error`, err.Error())
		return ctx.InternalServerError()
	}
	return ctx.OK(&app.Token{Token: authToken})
	// TokenController_Access: end_implement
}

// Auth runs the auth action.
func (c *TokenController) Auth(ctx *app.AuthTokenContext) error {
	// TokenController_Auth: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.GetUserByID(int(ctx.Value(CtxKey("user_id")).(int64)))
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get user`, `error`, err.Error())
		if err == model.ErrNotFound {
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
	// TokenController_Auth: end_implement
}
