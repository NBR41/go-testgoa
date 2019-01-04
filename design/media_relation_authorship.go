package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	authorshipPath   = "/authorships"
	authorshipIDPath = "/:authorship_id"
	attrAuthorshipID = func() { Attribute("authorship_id", Integer, "Unique Authorship ID", defIDConstraint) }
)

// AuthorshipMedia defines the media type used to render authorship.
var AuthorshipMedia = MediaType("application/vnd.authorship+json", func() {
	Description("An Author authorship")

	Attributes(func() {
		attrAuthorID()
		Attribute("author", AuthorMedia, "author struct")
		attrBookID()
		Attribute("book", BookMedia, "book struct")
		attrRoleID()
		Attribute("role", RoleMedia, "role struct")
		attrHref()
		Required("author_id", "book_id", "role_id", "href")
	})

	View("default", func() {
		Attribute("author_id")
		Attribute("author")
		Attribute("book_id")
		Attribute("book")
		Attribute("role_id")
		Attribute("role")
		Attribute("href")
	})
})

var _ = Resource("authorships", func() {
	BasePath(authorshipPath)
	DefaultMedia(AuthorshipMedia)

	Action("list", func() {
		Description("List authorships")
		Routing(GET(""))
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(OK, CollectionOf(AuthorshipMedia))
		// user NotFound
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new authorship")
		Routing(POST(""))
		Payload(func() {
			attrAuthorID()
			attrBookID()
			attrRoleID()
			Required("author_id", "book_id", "role_id")
		})
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/authorships/[0-9]+")
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
		Description("Get authorships by ids")
		Routing(GET(authorshipIDPath))
		Params(func() {
			attrAuthorshipID()
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
		Description("delete authorships by ids")
		Routing(DELETE(authorshipIDPath))
		Params(func() {
			attrAuthorshipID()
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
