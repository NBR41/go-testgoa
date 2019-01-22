package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("relationClass", func() {
	Parent("classes")

	Action("listCategories", func() {
		Description("List categories by class")
		Routing(GET(categoryPath))
		fCategoryList()
	})

	Action("listSeries", func() {
		Description("List series by class")
		Routing(GET(seriesPath))
		fSeriesList()
	})

	Action("listSeriesByCategory", func() {
		Description("List series by class and category")
		Routing(GET(categoryPath + categoryIDPath + seriesPath))
		Params(attrCategoryID)
		fSeriesList()
	})
})
