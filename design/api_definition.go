package design

import (
	// All symbols
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	defStringConstraint = func() {
		MinLength(1)
		MaxLength(128)
	}

	defIDConstraint = func() { Minimum(1) }

	attrHref = func() { Attribute("href", String, "API href for making requests") }
)

var _ = API("my-inventory", func() {
	Title("The virtual my-inventory")
	Description("An API to manage my-inventory")
	Contact(func() {
		Name("NBR41")
		Email("nbr41@foo.bar")
		URL("http://my-inventory.design")
	})
	License(func() {
		Name("MIT")
		//URL("")
	})
	Docs(func() {
		Description("api docs")
		//URL("http://goa.design/getting-started.html")
	})
	Host("localhost:8089")
	Scheme("http")

	Origin("http://localhost:4200", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		MaxAge(600)
		Headers("Authorization", "Origin", "Content-Type", "Accept")
		Credentials()
	})

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})
