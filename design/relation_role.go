package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("relationRole", func() {
	Parent("roles")

	Action("listAuthors", func() {
		Description("List authors by role")
		Routing(GET(authorPath))
		fAuthorList()
	})

	Action("listSeriesByAuthor", func() {
		Description("List series by role and author")
		Routing(GET(authorPath + authorIDPath + seriesPath))
		Params(attrAuthorID)
		fSeriesList()
	})
})
