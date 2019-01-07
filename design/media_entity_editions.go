package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	editionPath   = "/editions"
	editionIDPath = "/:edition_id"
	attrEditionID = func() { Attribute("edition_id", Integer, "Unique Edition ID", defIDConstraint) }
)

// EditionMedia defines the media type used to render editions.
var EditionMedia = MediaType("application/vnd.edition+json", func() {
	Description("An edition")

	Attributes(func() {
		attrEditionID()
		attrBookID()
		Attribute("book", BookMedia, "book struct")
		attrCollectionID()
		Attribute("collection", CollectionMedia, "collection struct")
		attrPrintID()
		Attribute("print", PrintMedia, "print struct")
		attrHref()
		Required("edition_id", "book_id", "collection_id", "print_id", "href")
	})

	View("default", func() {
		Attribute("edition_id")
		Attribute("book_id")
		Attribute("book")
		Attribute("collection_id")
		Attribute("collection")
		Attribute("print_id")
		Attribute("print")
		Attribute("href")
	})
})

var _ = Resource("editions", func() {
	BasePath(editionPath)
	DefaultMedia(EditionMedia)

	Action("list", func() {
		Description("List editions")
		Routing(GET(""))
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(OK, CollectionOf(EditionMedia))
		// user NotFound
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new edition")
		Routing(POST(""))
		Payload(func() {
			attrBookID()
			attrCollectionID()
			attrPrintID()
			Required("book_id", "collection_id", "print_id")
		})
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/editions/[0-9]+")
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
		Description("Get book edition by id")
		Routing(GET(editionIDPath))
		Params(func() {
			attrEditionID()
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
		Description("delete book edition by id")
		Routing(DELETE(editionIDPath))
		Params(func() {
			attrEditionID()
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
