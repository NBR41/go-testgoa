package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	seriesPath     = "/series"
	seriesIDPath   = "/:series_id"
	attrSeriesID   = func() { Attribute("series_id", Integer, "Unique Series ID", defIDConstraint) }
	attrSeriesName = func() { Attribute("series_name", String, "Series Name (Akira/Dragon ball)", defStringConstraint) }
)

//SeriesMedia defines the media type used to render series.
var SeriesMedia = MediaType("application/vnd.series+json", func() {
	Description("A Serie")

	Attributes(func() {
		attrSeriesID()
		attrSeriesName()
		attrCategoryID()
		Attribute("category", CategoryMedia, "category struct")
		attrHref()
		Required("series_id", "category_id", "series_name", "href")
	})

	View("default", func() {
		Attribute("series_id")
		Attribute("series_name")
		Attribute("category_id")
		Attribute("category")
		Attribute("href")
	})

	View("link", func() {
		Attribute("series_id")
		Attribute("series_name")
		Attribute("category_id")
		Attribute("href")
	})
})

var _ = Resource("series", func() {
	BasePath(seriesPath)
	DefaultMedia(SeriesMedia)

	Action("list", func() {
		Description("List series")
		Routing(GET(""))
		// ok
		Response(OK, CollectionOf(SeriesMedia))
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	})

	Action("show", func() {
		Description("Get serie by id")
		Routing(GET(seriesIDPath))
		Params(attrSeriesID)
		// ok
		Response(OK)
		// series not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new series")
		Routing(POST(""))
		Payload(func() {
			attrSeriesName()
			attrCategoryID()
			Required("series_name", "category_id")
		})
		Security(JWTAuth)
		// unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/series/[0-9]+")
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update serie by id")
		Routing(PUT(seriesIDPath))
		Params(attrSeriesID)
		Payload(func() {
			attrSeriesName()
			attrCategoryID()
		})
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(NoContent)
		// NotFound
		Response(NotFound)
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Description("delete serie by id")
		Routing(DELETE(seriesIDPath))
		Params(attrSeriesID)
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(NoContent)
		// NotFound
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
})
