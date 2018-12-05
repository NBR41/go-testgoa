package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	attrToken        = func() { Attribute("token", String, "token", func() { MinLength(1) }) }
	attrAccessToken  = func() { Attribute("access_token", String, "Access Token", func() { MinLength(1) }) }
	attrRefreshToken = func() { Attribute("refresh_token", String, "Refresh Token", func() { MinLength(1) }) }
)

// TokenMedia defines the media type used to render users.
var TokenMedia = MediaType("application/vnd.token+json", func() {
	Description("A token")

	Attributes(func() {
		attrToken()
		Required("token")
	})

	View("default", func() {
		Attribute("token")
	})
})

// AuthTokenMedia defines the media type used to render authenticated users.
var AuthTokenMedia = MediaType("application/vnd.authtoken+json", func() {
	Description("An auth token")
	Attributes(func() {
		Attribute("user", UserMedia, "user struct")
		attrAccessToken()
		attrRefreshToken()
		Required("user", "refresh_token", "access_token")
	})

	View("default", func() {
		Attribute("user")
		Attribute("refresh_token")
		Attribute("access_token")
	})
})

var _ = Resource("token", func() {
	BasePath("/token")

	Action("access", func() {
		Description("Get users")
		Routing(GET("/access_token"))
		Security(JWTAuth)
		// OK
		Response(OK, TokenMedia)
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("auth", func() {
		Description("Get users")
		Routing(GET("/auth"))
		Security(JWTAuth)
		// OK
		Response(OK, AuthTokenMedia)
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
})
