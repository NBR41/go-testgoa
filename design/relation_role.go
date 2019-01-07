package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("relationRole", func() {
	Parent("roles")

	Action("listAuthors", func() {
		Description("List authors by role")
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

	Action("listSeriesByAuthors", func() {
		Description("List series by role and authors")
		Routing(GET(authorPath + authorIDPath + seriesPath))
		Params(attrAuthorID)
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
