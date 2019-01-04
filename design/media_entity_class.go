package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	classIDPath   = "/:class_id"
	attrClassID   = func() { Attribute("class_id", Integer, "Unique Class ID", defIDConstraint) }
	attrClassName = func() { Attribute("class_name", String, "Class Name (Shonen/Shojo/Seinen)", defStringConstraint) }
)

//ClassMedia defines the media type used to render classes.
var ClassMedia = MediaType("application/vnd.class+json", func() {
	Description("A Class")

	Attributes(func() {
		attrClassID()
		attrClassName()
		attrHref()
		Required("class_id", "class_name", "href")
	})

	View("default", func() {
		Attribute("class_id")
		Attribute("class_name")
		Attribute("href")
	})

	View("link", func() {
		Attribute("class_id")
		Attribute("class_name")
		Attribute("href")
	})
})

var _ = Resource("classes", func() {
	BasePath("/classes")
	DefaultMedia(ClassMedia)

	Action("list", func() {
		Description("List classes")
		Routing(GET(""))
		// ok
		Response(OK, CollectionOf(ClassMedia))
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	})

	Action("show", func() {
		Description("Get class by id")
		Routing(GET(classIDPath))
		Params(attrClassID)
		// ok
		Response(OK)
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new class")
		Routing(POST(""))
		Payload(func() {
			attrClassName()
			Required("class_name")
		})
		Security(JWTAuth)
		// unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/classes/[0-9]+")
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update class by id")
		Routing(PUT(classIDPath))
		Params(attrClassID)
		Payload(func() {
			attrClassName()
			Required("class_name")
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
		Description("delete class by id")
		Routing(DELETE(classIDPath))
		Params(attrClassID)
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
