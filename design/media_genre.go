package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	genreIDPath   = "/:genre_id"
	attrGenreID   = func() { Attribute("genre_id", Integer, "Unique Genre ID", defIDConstraint) }
	attrGenreName = func() { Attribute("genre_name", String, "Genre Name (Shonen/Shojo/Seinen)", defStringConstraint) }
)

//GenreMedia defines the media type used to render genres.
var GenreMedia = MediaType("application/vnd.genre+json", func() {
	Description("A Genre")

	Attributes(func() {
		attrGenreID()
		attrGenreName()
		attrHref()
		Required("genre_id", "genre_name", "href")
	})

	View("default", func() {
		Attribute("genre_id")
		Attribute("genre_name")
		Attribute("href")
	})

	View("link", func() {
		Attribute("genre_id")
		Attribute("genre_name")
		Attribute("href")
	})
})

var _ = Resource("genres", func() {
	BasePath("/genres")
	DefaultMedia(GenreMedia)

	Action("list", func() {
		Description("Get genres")
		Routing(GET(""))
		// ok
		Response(OK, CollectionOf(GenreMedia))
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	})

	Action("show", func() {
		Description("Get genre by id")
		Routing(GET(genreIDPath))
		Params(attrGenreID)
		// ok
		Response(OK)
		// genre not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new genre")
		Routing(POST(""))
		Payload(func() {
			Member("name")
			Required("name")
		})
		Security(JWTAuth)
		// unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/genres/[0-9]+")
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update genre by id")
		Routing(PUT(genreIDPath))
		Params(attrGenreID)
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
		Description("delete genre by id")
		Routing(DELETE(genreIDPath))
		Params(attrGenreID)
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
