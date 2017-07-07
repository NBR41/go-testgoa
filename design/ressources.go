package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	userIDPath = "/:userID"
	bookIDPath = "/:bookID"
)

var (
	userParam = func() {
		Param("userID", Integer, "User ID", func() {
			Minimum(1)
		})
	}
	bookParam = func() {
		Param("bookID", Integer, "Book ID", func() {
			Minimum(1)
		})
	}
)

var JWTAuth = JWTSecurity("JWTSec", func() {
	Description("Use JWT t oauthenticate")
})

var _ = Resource("authenticate", func() {
	BasePath("")
	DefaultMedia(UserMedia)
	Action("auth", func() {
		Description("Get users")
		Routing(POST(""))
		Payload(func() {
			Member("login", String, "email or nickname", func() {
				MinLength(5)
			})
			Required("login")
			passwordAttribute()
			Required("password")
		})
		Response(OK, TokenMedia)
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
		Params(userParam)
		Response(OK)
		Response(NotFound)
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
		Response(BadRequest, ErrorMedia)
	})
	Action("update", func() {
		Description("Update user by id")
		Routing(PUT(userIDPath))
		Params(userParam)
		Payload(func() {
			Member("nickname")
			Required("nickname")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("delete", func() {
		Description("delete user by id")
		Routing(DELETE(userIDPath))
		Params(userParam)
		Security(JWTAuth)
		Response(NoContent)
		Response(NotFound)
		Response(Unauthorized)
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
		Params(bookParam)
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
		Response(Created, "/books/[0-9]+")
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
	Action("update", func() {
		Description("Update book by id")
		Routing(PUT(bookIDPath))
		Params(bookParam)
		Payload(func() {
			Member("name")
			Required("name")
		})
		Security(JWTAuth)
		Response(NoContent)
		Response(NotFound)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
	Action("delete", func() {
		Description("delete book by id")
		Routing(DELETE(bookIDPath))
		Params(bookParam)
		Security(JWTAuth)
		Response(NoContent)
		Response(NotFound)
		Response(Unauthorized)
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
		Response(OK, CollectionOf(BookMedia))
		Response(Unauthorized)
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
		Response(Created, "/books/[0-9]+")
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
	Action("show", func() {
		Description("Get ownerships by ids")
		Routing(GET(bookIDPath))
		Params(bookParam)
		Security(JWTAuth)
		Response(OK)
		Response(NotFound)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
	Action("delete", func() {
		Description("delete ownerships by ids")
		Routing(DELETE(bookIDPath))
		Params(bookParam)
		Security(JWTAuth)
		Response(NoContent)
		Response(NotFound)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})
