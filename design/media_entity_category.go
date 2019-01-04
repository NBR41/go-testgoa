package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	categoryPath     = "/categories"
	categoryIDPath   = "/:category_id"
	attrCategoryID   = func() { Attribute("category_id", Integer, "Unique Category ID", defIDConstraint) }
	attrCategoryName = func() {
		Attribute("category_name", String, "Category Name (Shonen/Shojo/Seinen)", defStringConstraint)
	}
)

//CategoryMedia defines the media type used to render categories.
var CategoryMedia = MediaType("application/vnd.category+json", func() {
	Description("A Category")

	Attributes(func() {
		attrCategoryID()
		attrCategoryName()
		attrHref()
		Required("category_id", "category_name", "href")
	})

	View("default", func() {
		Attribute("category_id")
		Attribute("category_name")
		Attribute("href")
	})

	View("link", func() {
		Attribute("category_id")
		Attribute("href")
	})
})

var _ = Resource("categories", func() {
	BasePath(categoryPath)
	DefaultMedia(CategoryMedia)

	Action("list", func() {
		Description("List categories")
		Routing(GET(""))
		// ok
		Response(OK, CollectionOf(CategoryMedia))
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	})

	Action("show", func() {
		Description("Get category by id")
		Routing(GET(categoryIDPath))
		Params(attrCategoryID)
		// ok
		Response(OK)
		// category not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new category")
		Routing(POST(""))
		Payload(func() {
			attrCategoryName()
			Required("category_name")
		})
		Security(JWTAuth)
		// unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/categories/[0-9]+")
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update category by id")
		Routing(PUT(categoryIDPath))
		Params(attrCategoryID)
		Payload(func() {
			attrCategoryName()
			Required("category_name")
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
		Description("delete category by id")
		Routing(DELETE(categoryIDPath))
		Params(attrCategoryID)
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
