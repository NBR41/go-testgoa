package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("relationAuthor", func() {
	Parent("authors")

	Action("listCategories", func() {
		Description("List categories by author")
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

	Action("listClasses", func() {
		Description("List classes by author")
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

	Action("listRoles", func() {
		Description("List roles by author")
		Routing(GET(rolePath))
		// ok
		Response(OK, CollectionOf(RoleMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("listSeries", func() {
		Description("List series by author")
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
		Description("List series by author and category")
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

	Action("listSeriesByClass", func() {
		Description("List series by author and class")
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

	Action("listSeriesByRole", func() {
		Description("List series by author and role")
		Routing(GET(rolePath + roleIDPath + seriesPath))
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
