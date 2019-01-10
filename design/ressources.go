package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// JWTAuth jwt security
var JWTAuth = JWTSecurity("JWTSec", func() {
	Description("Use JWT to authenticate")
	Header("Authorization")
})

var _ = Resource("health", func() {
	BasePath("/_ah")
	Action("health", func() {
		Routing(GET("/health"))
		Description("Perform health check.")
		Response(InternalServerError)
		Response(OK, "text/plain")
	})
})

var _ = Resource("authenticate", func() {
	BasePath("/authenticate")

	Action("auth", func() {
		Description("Get users")
		Routing(POST(""))
		Payload(AuthenticatePayload)
		// OK
		Response(OK, AuthTokenMedia)
		// wrong credentials
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("password", func() {
	BasePath("/password")

	Action("get", func() {
		Description("Get password reset")
		Routing(GET(""))
		Params(func() {
			attrUserEmail()
			Required("email")
		})
		// OK
		Response(NoContent)
		// Not found (user not found)
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update user password")
		Routing(POST(""))
		Payload(PasswordChangePayload)
		// OK
		Response(NoContent)
		// App Error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("validation", func() {
	BasePath("/validation")

	Action("get", func() {
		Description("Get validation token")
		Routing(GET(userIDPath))
		Params(attrUserID)
		// only admins
		Security(JWTAuth)
		Response(Unauthorized)
		// OK
		Response(NoContent)
		// Not Found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("validate", func() {
		Description("Validate user")
		Routing(POST(""))
		Payload(func() {
			attrToken()
			Required("token")
		})
		//OK
		Response(NoContent)
		// user not found
		Response(NotFound)
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/swagger.json", "public/swagger/swagger.json")
})
