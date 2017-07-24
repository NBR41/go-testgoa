package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// UserMedia defines the media type used to render users.
var UserMedia = MediaType("application/vnd.user+json", func() {
	Description("A User")

	Attributes(func() {
		Attribute("id", Integer, "Unique User ID", func() {
			Minimum(1)
		})
		Attribute("email", String, "user email", func() {
			Format("email")
			Example("me@foo.bar")
		})
		Attribute("nickname", String, "user nickname", func() {
			MinLength(1)
			MaxLength(32)
		})
		Attribute("href", String, "API href for making requests on the user")
		Attribute("is_admin", Boolean)
		Attribute("is_verified", Boolean)
		Required("id", "email", "nickname", "href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("email")
		Attribute("nickname")
		Attribute("href") // have a "default" view.
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("nickname")
		Attribute("href") // have a "default" view.
	})
})

// TokenMedia defines the media type used to render users.
var TokenMedia = MediaType("application/vnd.token+json", func() {
	Description("A token")

	Attributes(func() {
		Attribute("token", String, "Unique user ID", func() {
			MinLength(1)
		})
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
		Attribute("token", String, "Unique user ID", func() {
			MinLength(1)
		})
		Required("user", "token")
	})

	View("default", func() {
		Attribute("user")
		Attribute("token")
	})
})

// BookMedia defines the media type used to render books.
var BookMedia = MediaType("application/vnd.book+json", func() {
	Description("A Book")

	Attributes(func() {
		Attribute("id", Integer, "Unique Book ID", func() {
			Minimum(1)
		})
		Attribute("name", String, "Book Name", func() {
			MinLength(1)
			MaxLength(128)
		})
		Attribute("href", String, "API href for making requests on the book")

		Required("id", "name", "href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("href")
	})

	View("link", func() {
		Attribute("id")
		Attribute("name")
		Attribute("href")
	})
})

// OwnershipMedia defines the media type used to render users.
var OwnershipMedia = MediaType("application/vnd.ownership+json", func() {
	Description("A User ownership")

	Attributes(func() {
		Attribute("user_id", Integer, "Unique User ID", func() {
			Minimum(1)
		})
		Attribute("book_id", Integer, "Unique Book ID", func() {
			Minimum(1)
		})
		Attribute("book", BookMedia, "book struct")
		Attribute("href", String, "API href for making requests on the ownership")

		Required("user_id", "book_id", "href")
	})

	View("default", func() {
		Attribute("user_id")
		Attribute("book_id")
		Attribute("book")
		Attribute("href")
	})
})
