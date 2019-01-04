package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	editorPath     = "/editors"
	editorIDPath   = "/:editor_id"
	attrEditorID   = func() { Attribute("editor_id", Integer, "Unique Editor ID", defIDConstraint) }
	attrEditorName = func() { Attribute("editor_name", String, "Editor Name (Glénat/Delcourt)", defStringConstraint) }
)

//EditorMedia defines the media type used to render editors.
var EditorMedia = MediaType("application/vnd.editor+json", func() {
	Description("An Editor")

	Attributes(func() {
		attrEditorID()
		attrEditorName()
		attrHref()
		Required("editor_id", "editor_name", "href")
	})

	View("default", func() {
		Attribute("editor_id")
		Attribute("editor_name")
		Attribute("href")
	})

	View("link", func() {
		Attribute("editor_id")
		Attribute("href")
	})
})

var _ = Resource("editors", func() {
	BasePath(editorPath)
	DefaultMedia(EditorMedia)

	Action("list", func() {
		Description("List editors")
		Routing(GET(""))
		// ok
		Response(OK, CollectionOf(EditorMedia))
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	})

	Action("show", func() {
		Description("Get editor by id")
		Routing(GET(editorIDPath))
		Params(attrEditorID)
		// ok
		Response(OK)
		// editor not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new editor")
		Routing(POST(""))
		Payload(func() {
			attrEditorName()
			Required("editor_name")
		})
		Security(JWTAuth)
		// unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/editors/[0-9]+")
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update editor by id")
		Routing(PUT(editorIDPath))
		Params(attrEditorID)
		Payload(func() {
			attrEditorName()
			Required("editor_name")
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
		Description("delete editor by id")
		Routing(DELETE(editorIDPath))
		Params(attrEditorID)
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
