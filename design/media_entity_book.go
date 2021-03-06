package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	bookPath     = "/books"
	bookIDPath   = "/:book_id"
	attrBookID   = func() { Attribute("book_id", Integer, "Unique Book ID", defIDConstraint) }
	attrBookISBN = func() { Attribute("book_isbn", String, "Book ISBN", defStringConstraint) }
	attrBookName = func() { Attribute("book_name", String, "Book Name", defStringConstraint) }

	fBookList = func() {
		// ok
		Response(OK, CollectionOf(BookMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	}
)

// BookMedia defines the media type used to render books.
var BookMedia = MediaType("application/vnd.book+json", func() {
	Description("A Book")

	Attributes(func() {
		attrBookID()
		attrBookISBN()
		attrBookName()
		attrSeriesID()
		Attribute("series", SeriesMedia, "series struct")
		attrHref()
		Required("book_id", "book_isbn", "book_name", "href")
	})

	View("default", func() {
		Attribute("book_id")
		Attribute("book_isbn")
		Attribute("book_name")
		Attribute("series_id")
		Attribute("series")
		Attribute("href")
	})

	View("link", func() {
		Attribute("book_id")
		Attribute("book_isbn")
		Attribute("series_id")
		Attribute("href")
	})
})

// BookDetailMedia defines the media type used to render book details.
var BookDetailMedia = MediaType("application/vnd.book_detail+json", func() {
	Description("A Book Detail")

	Attributes(func() {
		Attribute("book", BookMedia, "book struct")
		Attribute("edition", EditionMedia, "edition struct")
		Attribute("authors", CollectionOf(AuthorshipMedia), "authors struct")
		Attribute("classes", CollectionOf(ClassificationMedia), "classification struct")
		Required("book", "edition", "authors", "classes")
	})

	View("default", func() {
		Attribute("book")
		Attribute("edition")
		Attribute("authors")
		Attribute("classes")
	})
})

var _ = Resource("books", func() {
	BasePath(bookPath)
	DefaultMedia(BookMedia)

	Action("isbnSearch", func() {
		Description("search a book by ISBN")
		Routing(GET("/add"))
		Params(func() {
			attrBookISBN()
			Required("book_isbn")
		})
		Response(NotFound)
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
		Response(OK, BookDetailMedia)
	})

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
			Member("book_isbn")
			Member("book_name")
			Member("series_id")
			Required("book_isbn", "book_name", "series_id")
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
			Member("book_name")
			Member("series_id")
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
