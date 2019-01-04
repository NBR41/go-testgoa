package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("relationCategory", func() {
	Parent("categories")

	Action("listAuthors", func() {
		Description("List authors by category")
		Routing(GET(authorPath))
		// ok
		Response(OK, CollectionOf(AuthorMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("listClasses", func() {
		Description("List classes by category")
		Routing(GET(classPath))
		// ok
		Response(OK, CollectionOf(ClassMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("listSeries", func() {
		Description("List series by category")
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

	Action("listSeriesByClass", func() {
		Description("List series by category and class")
		Routing(GET(classPath + classIDPath + seriesPath))
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
