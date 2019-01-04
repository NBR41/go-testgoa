package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	authorPath     = "/authors"
	authorIDPath   = "/:author_id"
	attrAuthorID   = func() { Attribute("author_id", Integer, "Unique Author ID", defIDConstraint) }
	attrAuthorName = func() { Attribute("author_name", String, "Author Name", defStringConstraint) }
)

//AuthorMedia defines the media type used to render authors.
var AuthorMedia = MediaType("application/vnd.author+json", func() {
	Description("An Author")

	Attributes(func() {
		attrAuthorID()
		attrAuthorName()
		attrHref()
		Required("author_id", "author_name", "href")
	})

	View("default", func() {
		Attribute("author_id")
		Attribute("author_name")
		Attribute("href")
	})

	View("link", func() {
		Attribute("author_id")
		Attribute("href")
	})
})

var _ = Resource("authors", func() {
	BasePath(authorPath)
	DefaultMedia(AuthorMedia)

	Action("list", func() {
		Description("List authors")
		Routing(GET(""))
		// ok
		Response(OK, CollectionOf(AuthorMedia))
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	})

	Action("show", func() {
		Description("Get author by id")
		Routing(GET(authorIDPath))
		Params(attrAuthorID)
		// ok
		Response(OK)
		// author not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new author")
		Routing(POST(""))
		Payload(func() {
			attrAuthorName()
			Required("author_name")
		})
		Security(JWTAuth)
		// unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/authors/[0-9]+")
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update author by id")
		Routing(PUT(authorIDPath))
		Params(attrAuthorID)
		Payload(func() {
			attrAuthorName()
			Required("author_name")
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
		Description("delete author by id")
		Routing(DELETE(authorIDPath))
		Params(attrAuthorID)
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
