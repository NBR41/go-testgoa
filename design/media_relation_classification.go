package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("classifications", func() {
	BasePath("/classifications")
	Parent("series")
	DefaultMedia(ClassMedia)

	Action("list", func() {
		Description("Get series classes")
		Routing(GET(""))
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(OK, CollectionOf(ClassMedia))
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
		Response(Created, "/books/[0-9]+/roles/[0-9]+")
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
		Description("Get series class")
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
