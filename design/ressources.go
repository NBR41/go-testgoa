package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	userIDPath = "/:user_id"
	bookIDPath = "/:book_id"
)

var JWTAuth = JWTSecurity("JWTSec", func() {
	Description("Use JWT t oauthenticate")
})

var _ = Resource("authenticate", func() {
	BasePath("/authenticate")
	Action("auth", func() {
		Description("Get users")
		Routing(POST(""))
		Payload(AuthenticatePayload)
		Response(OK, TokenMedia)
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
		Response(NoContent)
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
	Action("update", func() {
		Description("Update user password")
		Routing(POST(""))
		Payload(PasswordPayload)
		Security(JWTAuth)
		Response(Unauthorized)
		Response(NoContent)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("validation", func() {
	BasePath("/validation")
	Parent("users")
	Action("get", func() {
		Description("Get validation token")
		Routing(GET(""))
		Security(JWTAuth)
		Response(Unauthorized)
		Response(NoContent)
		Response(BadRequest, ErrorMedia)
	})
	Action("validate", func() {
		Description("Validate user")
		Routing(POST(""))
		Payload(func() {
			Attribute("token", String, "validation token", func() {
				MinLength(1)
				MaxLength(32)
			})
		},
		)
		Response(NoContent)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("users", func() {
	BasePath("/users")
	DefaultMedia(UserMedia)
	Action("list", func() {
		Description("Get users")
		Routing(GET(""))
		Response(OK, CollectionOf(UserMedia))
	})
	Action("show", func() {
		Description("Get user by id")
		Routing(GET(userIDPath))
		Params(UserIDParam)
		Response(OK)
		Response(NotFound)
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
	Action("create", func() {
		Description("Create new user")
		Routing(POST(""))
		Payload(func() {
			Member("email")
			Required("email")
			Member("nickname")
			Required("nickname")
		})
		Response(Created, "/users/[0-9]+")
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
	Action("update", func() {
		Description("Update user by id")
		Routing(PUT(userIDPath))
		Params(UserIDParam)
		Payload(func() {
			Member("nickname")
			Required("nickname")
		})
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
	Action("delete", func() {
		Description("delete user by id")
		Routing(DELETE(userIDPath))
		Params(UserIDParam)
		Security(JWTAuth)
		Response(Unauthorized)
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("books", func() {
	BasePath("/books")
	DefaultMedia(BookMedia)
	Action("list", func() {
		Description("Get books")
		Routing(GET(""))
		Response(OK, CollectionOf(BookMedia))
	})
	Action("show", func() {
		Description("Get book by id")
		Routing(GET(bookIDPath))
		Params(BookIDParam)
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("create", func() {
		Description("Create new book")
		Routing(POST(""))
		Payload(func() {
			Member("name")
			Required("name")
		})
		Security(JWTAuth)
		Response(Unauthorized)
		Response(Created, "/books/[0-9]+")

		Response(BadRequest, ErrorMedia)
	})
	Action("update", func() {
		Description("Update book by id")
		Routing(PUT(bookIDPath))
		Params(BookIDParam)
		Payload(func() {
			Member("name")
			Required("name")
		})
		Security(JWTAuth)
		Response(Unauthorized)
		Response(NoContent)
		Response(NotFound)

		Response(BadRequest, ErrorMedia)
	})
	Action("delete", func() {
		Description("delete book by id")
		Routing(DELETE(bookIDPath))
		Params(BookIDParam)
		Security(JWTAuth)
		Response(Unauthorized)
		Response(NoContent)
		Response(NotFound)

		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("ownerships", func() {
	BasePath("/ownerships")
	Parent("users")
	DefaultMedia(OwnershipMedia)
	Action("list", func() {
		Description("Get ownerships")
		Routing(GET(""))
		Security(JWTAuth)
		Response(Unauthorized)
		Response(OK, CollectionOf(BookMedia))
		Response(NotFound)
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
	Action("create", func() {
		Description("Create new ownership")
		Routing(POST(""))
		Payload(func() {
			Member("book_id", Integer, func() {
				Minimum(1)
			})
			Required("book_id")
		})
		Security(JWTAuth)
		Response(Unauthorized)
		Response(Created, "/books/[0-9]+")
		Response(NotFound)
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
	Action("show", func() {
		Description("Get ownerships by ids")
		Routing(GET(bookIDPath))
		Params(BookIDParam)
		Security(JWTAuth)
		Response(Unauthorized)
		Response(OK)
		Response(NotFound)
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
	Action("delete", func() {
		Description("delete ownerships by ids")
		Routing(DELETE(bookIDPath))
		Params(BookIDParam)
		Security(JWTAuth)
		Response(Unauthorized)
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
})
