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
	Header("Authorization")
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
			Param("email", String, "user email", func() {
				Format("email")
				Example("me@foo.bar")
			})
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
	Parent("users")

	Action("get", func() {
		Description("Get validation token")
		Routing(GET(""))
		// only admins
		Security(JWTAuth)
		Response(Unauthorized)
		// OK
		Response(NoContent)
		// Not Found
		Response(NotFound)
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("validate", func() {
		Description("Validate user")
		Routing(POST(""))
		Payload(
			func() {
				Member("token", String, "token", func() {
					MinLength(1)
				})
				Required("token")
			},
		)
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

var _ = Resource("users", func() {
	BasePath("/users")
	DefaultMedia(UserMedia)

	Action("list", func() {
		Description("Get users, optionnaly filtering by nickname or email ")
		Routing(GET(""))
		Params(func() {
			Param("nickname")
			Param("email")
		})
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
		Params(func() {
			Param("user_id", Integer, "Unique User ID", func() {
				Minimum(1)
			})
		})
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
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update user by id")
		Routing(PUT(userIDPath))
		Params(func() {
			Param("user_id", Integer, "Unique User ID", func() {
				Minimum(1)
			})
		})
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
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Description("delete user by id")
		Routing(DELETE(userIDPath))
		Params(func() {
			Param("user_id", Integer, "Unique User ID", func() {
				Minimum(1)
			})
		})
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

var _ = Resource("books", func() {
	BasePath("/books")
	DefaultMedia(BookMedia)

	Action("list", func() {
		Description("Get books")
		Routing(GET(""))
		// ok
		Response(OK, CollectionOf(BookMedia))
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	})

	Action("show", func() {
		Description("Get book by id")
		Routing(GET(bookIDPath))
		Params(func() {
			Param("book_id", Integer, "Unique Book ID", func() {
				Minimum(1)
			})
		})
		// ok
		Response(OK)
		// book not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new book")
		Routing(POST(""))
		Payload(func() {
			Member("isbn")
			Member("name")
			Required("isbn", "name")
		})
		Security(JWTAuth)
		// unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/books/[0-9]+")
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update book by id")
		Routing(PUT(bookIDPath))
		Params(func() {
			Param("book_id", Integer, "Unique Book ID", func() {
				Minimum(1)
			})
		})
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
		Description("delete book by id")
		Routing(DELETE(bookIDPath))
		Params(func() {
			Param("book_id", Integer, "Unique Book ID", func() {
				Minimum(1)
			})
		})
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

var _ = Resource("ownerships", func() {
	BasePath("/ownerships")
	Parent("users")
	DefaultMedia(OwnershipMedia)

	Action("list", func() {
		Description("Get ownerships")
		Routing(GET(""))
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(OK, CollectionOf(OwnershipMedia))
		// user NotFound
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new ownership")
		Routing(POST(""))
		Payload(func() {
			Member("book_id", Integer, "Unique Book ID", func() {
				Minimum(1)
			})
			Required("book_id")
		})
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/books/[0-9]+")
		// user NotFound
		Response(NotFound)
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("add", func() {
		Description("Create new book and ownership by isbn")
		Routing(POST(""))
		Payload(func() {
			Member("isbn", String, "Unique ISBN ID", func() {
				MinLength(1)
				MaxLength(128)
			})
			Required("isbn")
		})
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/books/[0-9]+")
		// user NotFound
		Response(NotFound)
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("show", func() {
		Description("Get ownerships by ids")
		Routing(GET(bookIDPath))
		Params(func() {
			Param("book_id", Integer, "Unique Book ID", func() {
				Minimum(1)
			})
		})
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(OK)
		// user or book NotFound
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Description("delete ownerships by ids")
		Routing(DELETE(bookIDPath))
		Params(func() {
			Param("book_id", Integer, "Unique Book ID", func() {
				Minimum(1)
			})
		})
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(NoContent)
		// user or book NotFound
		Response(NotFound)
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
