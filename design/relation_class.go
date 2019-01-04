package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("relationClass", func() {
	Parent("classes")

	Action("listCategories", func() {
		Description("List categories by class")
		Routing(GET(categoryPath))
		// ok
		Response(OK, CollectionOf(CategoryMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("listSeries", func() {
		Description("List series by class")
		Routing(GET(seriesPath))
		// ok
		Response(OK, CollectionOf(SeriesMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("listSeriesByCategory", func() {
		Description("List series by class and category")
		Routing(GET(categoryPath + categoryIDPath + seriesPath))
		// ok
		Response(OK, CollectionOf(SeriesMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
})
