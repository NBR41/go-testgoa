package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	printIDPath   = "/:print_id"
	attrPrintID   = func() { Attribute("print_id", Integer, "Unique Print ID", defIDConstraint) }
	attrPrintName = func() {
		Attribute("print_name", String, "Print Name (Deluxe/Ultimate/Pocket)", defStringConstraint)
	}
)

//PrintMedia defines the media type used to render edition types.
var PrintMedia = MediaType("application/vnd.print+json", func() {
	Description("A Print")

	Attributes(func() {
		attrPrintID()
		attrPrintName()
		attrHref()
		Required("print_id", "print_name", "href")
	})

	View("default", func() {
		Attribute("print_id")
		Attribute("print_name")
		Attribute("href")
	})

	View("link", func() {
		Attribute("print_id")
		Attribute("print_name")
		Attribute("href")
	})
})

var _ = Resource("prints", func() {
	BasePath("/prints")
	DefaultMedia(PrintMedia)

	Action("list", func() {
		Description("List prints")
		Routing(GET(""))
		// ok
		Response(OK, CollectionOf(PrintMedia))
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	})

	Action("show", func() {
		Description("Get print by id")
		Routing(GET(printIDPath))
		Params(attrPrintID)
		// ok
		Response(OK)
		// edition type not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new edition type")
		Routing(POST(""))
		Payload(func() {
			attrPrintName()
			Required("print_name")
		})
		Security(JWTAuth)
		// unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/prints/[0-9]+")
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update print by id")
		Routing(PUT(printIDPath))
		Params(attrPrintID)
		Payload(func() {
			attrPrintName()
			Required("print_name")
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
		Description("delete print by id")
		Routing(DELETE(printIDPath))
		Params(attrPrintID)
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
