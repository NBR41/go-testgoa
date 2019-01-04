package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// OwnershipMedia defines the media type used to render ownership.
var OwnershipMedia = MediaType("application/vnd.ownership+json", func() {
	Description("A User ownership")

	Attributes(func() {
		attrUserID()
		attrBookID()
		Attribute("book", BookMedia, "book struct")
		attrHref()
		Required("user_id", "book_id", "href")
	})

	View("default", func() {
		Attribute("user_id")
		Attribute("book_id")
		Attribute("book")
		Attribute("href")
	})
})

var _ = Resource("ownerships", func() {
	BasePath("/ownerships")
	Parent("users")
	DefaultMedia(OwnershipMedia)

	Action("list", func() {
		Description("List user ownerships")
		Routing(GET(""))
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(OK, CollectionOf(OwnershipMedia))
		// user NotFound
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new user ownership")
		Routing(POST(""))
		Payload(func() {
			attrBookID()
			Required("book_id")
		})
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/books/[0-9]+")
		// user NotFound
		Response(NotFound)
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("add", func() {
		Description("Create new book and ownership by isbn")
		Routing(POST("/isbn"))
		Payload(func() {
			attrBookISBN()
			Required("book_isbn")
		})
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/books/[0-9]+")
		// user NotFound
		Response(NotFound)
		// App error
		Response(UnprocessableEntity, ErrorMedia)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("show", func() {
		Description("Get user ownerships by ids")
		Routing(GET(bookIDPath))
		Params(attrBookID)
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
		Description("delete ownerships by ids")
		Routing(DELETE(bookIDPath))
		Params(attrBookID)
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
