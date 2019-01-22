package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("relationCategory", func() {
	Parent("categories")

	Action("listAuthors", func() {
		Description("List authors by category")
		Routing(GET(authorPath))
		fAuthorList()
	})

	Action("listClasses", func() {
		Description("List classes by category")
		Routing(GET(classPath))
		fClassList()
	})

	Action("listSeries", func() {
		Description("List series by category")
		Routing(GET(seriesPath))
		fSeriesList()
	})

	Action("listSeriesByClass", func() {
		Description("List series by category and class")
		Routing(GET(classPath + classIDPath + seriesPath))
		Params(attrClassID)
		fSeriesList()
	})
})
