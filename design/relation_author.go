package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("relationAuthor", func() {
	Parent("authors")

	Action("listCategories", func() {
		Description("List categories by author")
		Routing(GET(categoryPath))
		fCategoryList()
	})

	Action("listClasses", func() {
		Description("List classes by author")
		Routing(GET(classPath))
		fClassList()
	})

	Action("listRoles", func() {
		Description("List roles by author")
		Routing(GET(rolePath))
		fRoleList()
	})

	Action("listSeries", func() {
		Description("List series by author")
		Routing(GET(seriesPath))
		fSeriesList()
	})

	Action("listSeriesByCategory", func() {
		Description("List series by author and category")
		Routing(GET(categoryPath + categoryIDPath + seriesPath))
		Params(attrCategoryID)
		fSeriesList()
	})

	Action("listSeriesByClass", func() {
		Description("List series by author and class")
		Routing(GET(classPath + classIDPath + seriesPath))
		Params(attrClassID)
		fSeriesList()
	})

	Action("listSeriesByRole", func() {
		Description("List series by author and role")
		Routing(GET(rolePath + roleIDPath + seriesPath))
		Params(attrRoleID)
		fSeriesList()
	})
})
