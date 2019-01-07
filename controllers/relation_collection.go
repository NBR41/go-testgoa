package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationCollectionController implements the relationCollection resource.
type RelationCollectionController struct {
	*goa.Controller
}

// NewRelationCollectionController creates a relationCollection controller.
func NewRelationCollectionController(service *goa.Service) *RelationCollectionController {
	return &RelationCollectionController{Controller: service.NewController("RelationCollectionController")}
}

// ListBooks runs the listBooks action.
func (c *RelationCollectionController) ListBooks(ctx *app.ListBooksRelationCollectionContext) error {
	// RelationCollectionController_ListBooks: start_implement

	// Put your logic here

	res := app.BookCollection{}
	return ctx.OK(res)
	// RelationCollectionController_ListBooks: end_implement
}

// ListBooksByPrint runs the listBooksByPrint action.
func (c *RelationCollectionController) ListBooksByPrint(ctx *app.ListBooksByPrintRelationCollectionContext) error {
	// RelationCollectionController_ListBooksByPrint: start_implement

	// Put your logic here

	res := app.BookCollection{}
	return ctx.OK(res)
	// RelationCollectionController_ListBooksByPrint: end_implement
}

// ListBooksByPrintsSeries runs the listBooksByPrintsSeries action.
func (c *RelationCollectionController) ListBooksByPrintsSeries(ctx *app.ListBooksByPrintsSeriesRelationCollectionContext) error {
	// RelationCollectionController_ListBooksByPrintsSeries: start_implement

	// Put your logic here

	res := app.BookCollection{}
	return ctx.OK(res)
	// RelationCollectionController_ListBooksByPrintsSeries: end_implement
}

// ListBooksBySeries runs the listBooksBySeries action.
func (c *RelationCollectionController) ListBooksBySeries(ctx *app.ListBooksBySeriesRelationCollectionContext) error {
	// RelationCollectionController_ListBooksBySeries: start_implement

	// Put your logic here

	res := app.BookCollection{}
	return ctx.OK(res)
	// RelationCollectionController_ListBooksBySeries: end_implement
}

// ListBooksBySeriesPrints runs the listBooksBySeriesPrints action.
func (c *RelationCollectionController) ListBooksBySeriesPrints(ctx *app.ListBooksBySeriesPrintsRelationCollectionContext) error {
	// RelationCollectionController_ListBooksBySeriesPrints: start_implement

	// Put your logic here

	res := app.BookCollection{}
	return ctx.OK(res)
	// RelationCollectionController_ListBooksBySeriesPrints: end_implement
}

// ListPrints runs the listPrints action.
func (c *RelationCollectionController) ListPrints(ctx *app.ListPrintsRelationCollectionContext) error {
	// RelationCollectionController_ListPrints: start_implement

	// Put your logic here

	res := app.PrintCollection{}
	return ctx.OK(res)
	// RelationCollectionController_ListPrints: end_implement
}

// ListPrintsBySeries runs the listPrintsBySeries action.
func (c *RelationCollectionController) ListPrintsBySeries(ctx *app.ListPrintsBySeriesRelationCollectionContext) error {
	// RelationCollectionController_ListPrintsBySeries: start_implement

	// Put your logic here

	res := app.PrintCollection{}
	return ctx.OK(res)
	// RelationCollectionController_ListPrintsBySeries: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationCollectionController) ListSeries(ctx *app.ListSeriesRelationCollectionContext) error {
	// RelationCollectionController_ListSeries: start_implement

	// Put your logic here

	res := app.SeriesCollection{}
	return ctx.OK(res)
	// RelationCollectionController_ListSeries: end_implement
}

// ListSeriesByPrint runs the listSeriesByPrint action.
func (c *RelationCollectionController) ListSeriesByPrint(ctx *app.ListSeriesByPrintRelationCollectionContext) error {
	// RelationCollectionController_ListSeriesByPrint: start_implement

	// Put your logic here

	res := app.SeriesCollection{}
	return ctx.OK(res)
	// RelationCollectionController_ListSeriesByPrint: end_implement
}
