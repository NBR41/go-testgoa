package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	userPath      = "/users"
	userIDPath    = "/:user_id"
	attrUserID    = func() { Attribute("user_id", Integer, "Unique User ID", defIDConstraint) }
	attrUserEmail = func() {
		Attribute("email", String, "user email", func() {
			Format("email")
			Example("me@foo.bar")
		})
	}
	attrUserNickname = func() {
		Attribute("nickname", String, "user nickname", func() {
			MinLength(1)
			MaxLength(32)
		})
	}
	attrUserLogin    = func() { Attribute("login", String, "email or nickname", func() { MinLength(1) }) }
	attrUserPassword = func() {
		Attribute("password", String, "user password", func() {
			MinLength(1)
			MaxLength(32)
		})
	}
	attrUserIsAdmin     = func() { Attribute("is_admin", Boolean) }
	attrUserIsValidated = func() { Attribute("is_validated", Boolean) }
)

// UserMedia defines the media type used to render users.
var UserMedia = MediaType("application/vnd.user+json", func() {
	Description("A User")

	Attributes(func() {
		attrUserID()
		attrUserEmail()
		attrUserNickname()
		attrHref()
		attrUserIsAdmin()
		attrUserIsValidated()
		Required("user_id", "email", "nickname", "is_admin", "is_validated", "href")
	})

	View("default", func() {
		Attribute("user_id")
		Attribute("email")
		Attribute("nickname")
		Attribute("is_admin")
		Attribute("is_validated")
		Attribute("href") // have a "default" view.
	})

	View("tiny", func() {
		Attribute("user_id")
		Attribute("nickname")
		Attribute("href") // have a "default" view.
	})
})

var _ = Resource("users", func() {
	BasePath(userPath)
	DefaultMedia(UserMedia)

	Action("list", func() {
		Description("Get users, optionnaly filtering by nickname or email ")
		Routing(GET(""))
		Params(func() {
			attrUserNickname()
			attrUserEmail()
		})
		Security(JWTAuth)
		Response(OK, CollectionOf(UserMedia, func() {
			View("default")
			View("tiny")
		}))
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	})

	Action("show", func() {
		Description("Get user by id")
		Routing(GET(userIDPath))
		Params(attrUserID)
		Security(JWTAuth)
		// OK
		Response(OK)
		// Not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new user")
		Routing(POST(""))
		Payload(UserCreatePayload)
		// OK
		Response(Created, "/users/[0-9]+")
		// App error
		Response(UnprocessableEntity, ErrorMedia)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update user by id")
		Routing(PUT(userIDPath))
		Params(attrUserID)
		Payload(func() {
			Member("nickname")
			Required("nickname")
		})
		Security(JWTAuth)
		// unauthorized
		Response(Unauthorized)
		// OK
		Response(NoContent)
		// user not found
		Response(NotFound)
		// App error
		Response(UnprocessableEntity, ErrorMedia)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Description("delete user by id")
		Routing(DELETE(userIDPath))
		Params(attrUserID)
		Security(JWTAuth)
		// unauthorized
		Response(Unauthorized)
		// OK
		Response(NoContent)
		// user not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
})
