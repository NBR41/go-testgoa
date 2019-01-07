package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationAuthorController implements the relationAuthor resource.
type RelationAuthorController struct {
	*goa.Controller
}

// NewRelationAuthorController creates a relationAuthor controller.
func NewRelationAuthorController(service *goa.Service) *RelationAuthorController {
	return &RelationAuthorController{Controller: service.NewController("RelationAuthorController")}
}

// ListCategories runs the listCategories action.
func (c *RelationAuthorController) ListCategories(ctx *app.ListCategoriesRelationAuthorContext) error {
	// RelationAuthorController_ListCategories: start_implement

	// Put your logic here

	res := app.CategoryCollection{}
	return ctx.OK(res)
	// RelationAuthorController_ListCategories: end_implement
}

// ListClasses runs the listClasses action.
func (c *RelationAuthorController) ListClasses(ctx *app.ListClassesRelationAuthorContext) error {
	// RelationAuthorController_ListClasses: start_implement

	// Put your logic here

	res := app.ClassCollection{}
	return ctx.OK(res)
	// RelationAuthorController_ListClasses: end_implement
}

// ListRoles runs the listRoles action.
func (c *RelationAuthorController) ListRoles(ctx *app.ListRolesRelationAuthorContext) error {
	// RelationAuthorController_ListRoles: start_implement

	// Put your logic here

	res := app.RoleCollection{}
	return ctx.OK(res)
	// RelationAuthorController_ListRoles: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationAuthorController) ListSeries(ctx *app.ListSeriesRelationAuthorContext) error {
	// RelationAuthorController_ListSeries: start_implement

	// Put your logic here

	res := app.SeriesCollection{}
	return ctx.OK(res)
	// RelationAuthorController_ListSeries: end_implement
}

// ListSeriesByCategory runs the listSeriesByCategory action.
func (c *RelationAuthorController) ListSeriesByCategory(ctx *app.ListSeriesByCategoryRelationAuthorContext) error {
	// RelationAuthorController_ListSeriesByCategory: start_implement

	// Put your logic here

	res := app.SeriesCollection{}
	return ctx.OK(res)
	// RelationAuthorController_ListSeriesByCategory: end_implement
}

// ListSeriesByClass runs the listSeriesByClass action.
func (c *RelationAuthorController) ListSeriesByClass(ctx *app.ListSeriesByClassRelationAuthorContext) error {
	// RelationAuthorController_ListSeriesByClass: start_implement

	// Put your logic here

	res := app.SeriesCollection{}
	return ctx.OK(res)
	// RelationAuthorController_ListSeriesByClass: end_implement
}

// ListSeriesByRole runs the listSeriesByRole action.
func (c *RelationAuthorController) ListSeriesByRole(ctx *app.ListSeriesByRoleRelationAuthorContext) error {
	// RelationAuthorController_ListSeriesByRole: start_implement

	// Put your logic here

	res := app.SeriesCollection{}
	return ctx.OK(res)
	// RelationAuthorController_ListSeriesByRole: end_implement
}
