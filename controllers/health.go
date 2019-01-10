package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// HealthController implements the health resource.
type HealthController struct {
	*goa.Controller
	fm Fmodeler
}

// NewHealthController creates a health controller.
func NewHealthController(service *goa.Service, fm Fmodeler) *HealthController {
	return &HealthController{
		Controller: service.NewController("HealthController"),
		fm:         fm,
	}
}

// Health runs the health action.
func (c *HealthController) Health(ctx *app.HealthHealthContext) error {
	// HealthController_Health: start_implement

	// Put your logic here
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.InternalServerError()
	}
	defer func() { m.Close() }()
	return ctx.OK([]byte("ok"))
	// HealthController_Health: end_implement
}
