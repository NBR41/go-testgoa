package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	collectionIDPath   = "/:collection_id"
	attrCollectionID   = func() { Attribute("collection_id", Integer, "Unique Collection ID", defIDConstraint) }
	attrCollectionName = func() {
		Attribute("collection_name", String, "Collection Name (Découverte/Shonen)", defStringConstraint)
	}
)

//CollectionMedia defines the media type used to render collections.
var CollectionMedia = MediaType("application/vnd.collection+json", func() {
	Description("A Collection")

	Attributes(func() {
		attrCollectionID()
		attrCollectionName()
		attrHref()
		Required("collection_id", "collection_name", "href")
	})

	View("default", func() {
		Attribute("collection_id")
		Attribute("collection_name")
		Attribute("href")
	})

	View("link", func() {
		Attribute("collection_id")
		Attribute("collection_name")
		Attribute("href")
	})
})

var _ = Resource("collections", func() {
	BasePath("/collections")
	DefaultMedia(CollectionMedia)

	Action("list", func() {
		Description("Get collections")
		Routing(GET(""))
		// ok
		Response(OK, CollectionOf(CollectionMedia))
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	})

	Action("show", func() {
		Description("Get collection by id")
		Routing(GET(collectionIDPath))
		Params(attrCollectionID)
		// ok
		Response(OK)
		// collection not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new collection")
		Routing(POST(""))
		Payload(func() {
			Member("name")
			Required("name")
		})
		Security(JWTAuth)
		// unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/collections/[0-9]+")
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update collection by id")
		Routing(PUT(collectionIDPath))
		Params(attrCollectionID)
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
		Description("delete collection by id")
		Routing(DELETE(collectionIDPath))
		Params(attrCollectionID)
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
