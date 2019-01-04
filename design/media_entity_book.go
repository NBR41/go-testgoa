package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	bookIDPath   = "/:book_id"
	attrBookID   = func() { Attribute("book_id", Integer, "Unique Book ID", defIDConstraint) }
	attrBookISBN = func() { Attribute("book_isbn", String, "Book ISBN", defStringConstraint) }
	attrBookName = func() { Attribute("book_name", String, "Book Name", defStringConstraint) }
)

// BookMedia defines the media type used to render books.
var BookMedia = MediaType("application/vnd.book+json", func() {
	Description("A Book")

	Attributes(func() {
		attrBookID()
		attrBookISBN()
		attrBookName()
		attrHref()
		Required("book_id", "book_isbn", "book_name", "href")
	})

	View("default", func() {
		Attribute("book_id")
		Attribute("book_isbn")
		Attribute("book_name")
		Attribute("href")
	})

	View("link", func() {
		Attribute("book_id")
		Attribute("book_isbn")
		Attribute("book_name")
		Attribute("href")
	})
})

var _ = Resource("books", func() {
	BasePath("/books")
	DefaultMedia(BookMedia)

	Action("list", func() {
		Description("List books")
		Routing(GET(""))
		// ok
		Response(OK, CollectionOf(BookMedia))
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	})

	Action("show", func() {
		Description("Get book by id")
		Routing(GET(bookIDPath))
		Params(attrBookID)
		// ok
		Response(OK)
		// book not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new book")
		Routing(POST(""))
		Payload(func() {
			Member("isbn")
			Member("name")
			Required("isbn", "name")
		})
		Security(JWTAuth)
		// unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/books/[0-9]+")
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update book by id")
		Routing(PUT(bookIDPath))
		Params(attrBookID)
		Payload(func() {
			Member("name")
			Required("name")
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
		Description("delete book by id")
		Routing(DELETE(bookIDPath))
		Params(attrBookID)
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
