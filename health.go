package main

import (
	"fmt"

	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/appmodel"
	"github.com/goadesign/goa"
)

// HealthController implements the health resource.
type HealthController struct {
	*goa.Controller
}

// NewHealthController creates a health controller.
func NewHealthController(service *goa.Service) *HealthController {
	return &HealthController{Controller: service.NewController("HealthController")}
}

// Health runs the health action.
func (c *HealthController) Health(ctx *app.HealthHealthContext) error {
	// HealthController_Health: start_implement

	// Put your logic here
	m, err := appmodel.GetModeler()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return fmt.Errorf("failed to connect to DB")
	}
	defer func() { m.Close() }()
	return ctx.OK([]byte("ok"))
	// HealthController_Health: end_implement
}
