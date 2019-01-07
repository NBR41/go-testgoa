package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// ClassificationMedia defines the media type used to render classification.
var ClassificationMedia = MediaType("application/vnd.classification+json", func() {
	Description("A series classification")

	Attributes(func() {
		Attribute("class", ClassMedia, "class struct")
		attrHref()
		Required("class", "href")
	})

	View("default", func() {
		Attribute("class")
		Attribute("href")
	})
})

var _ = Resource("classifications", func() {
	BasePath("/classifications")
	Parent("series")
	DefaultMedia(ClassificationMedia)

	Action("list", func() {
		Description("List series classes")
		Routing(GET(""))
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(OK, CollectionOf(ClassificationMedia))
		// user NotFound
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new series classification")
		Routing(POST(""))
		Payload(func() {
			attrClassID()
			Required("class_id")
		})
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/classes/[0-9]+")
		// user NotFound
		Response(NotFound)
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("show", func() {
		Description("Get series classification")
		Routing(GET(classIDPath))
		Params(func() {
			attrClassID()
		})
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(OK)
		// user or book NotFound
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Description("delete series classification")
		Routing(DELETE(classIDPath))
		Params(func() {
			attrClassID()
		})
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(NoContent)
		// user or book NotFound
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
})
