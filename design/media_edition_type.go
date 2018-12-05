package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	editionTypeIDPath   = "/:edition_type_id"
	attrEditionTypeID   = func() { Attribute("edition_type_id", Integer, "Unique Edition type ID", defIDConstraint) }
	attrEditionTypeName = func() {
		Attribute("edition_type_name", String, "Editor Name (Deluxe/Ultimate/Pocket)", defStringConstraint)
	}
)

//EditionTypeMedia defines the media type used to render edition types.
var EditionTypeMedia = MediaType("application/vnd.editiontype+json", func() {
	Description("An Edition Type")

	Attributes(func() {
		attrEditionTypeID()
		attrEditionTypeName()
		attrHref()
		Required("edition_type_id", "edition_type_name", "href")
	})

	View("default", func() {
		Attribute("edition_type_id")
		Attribute("edition_type_name")
		Attribute("href")
	})

	View("link", func() {
		Attribute("edition_type_id")
		Attribute("edition_type_name")
		Attribute("href")
	})
})

var _ = Resource("edition_types", func() {
	BasePath("/edition_types")
	DefaultMedia(EditionTypeMedia)

	Action("list", func() {
		Description("Get edition types")
		Routing(GET(""))
		// ok
		Response(OK, CollectionOf(EditionTypeMedia))
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	})

	Action("show", func() {
		Description("Get edition type by id")
		Routing(GET(editionTypeIDPath))
		Params(attrEditionTypeID)
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
			Member("name")
			Required("name")
		})
		Security(JWTAuth)
		// unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/edition_types/[0-9]+")
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update edition type by id")
		Routing(PUT(editionTypeIDPath))
		Params(attrEditionTypeID)
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
		Description("delete edition type by id")
		Routing(DELETE(editionTypeIDPath))
		Params(attrEditionTypeID)
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
