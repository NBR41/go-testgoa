package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("relationCollection", func() {
	Parent("collections")

	Action("listBooks", func() {
		Description("List books by collection")
		Routing(GET(bookPath))
		fBookList()
	})

	Action("listPrints", func() {
		Description("List prints by collection")
		Routing(GET(printPath))
		fPrintList()
	})

	Action("listSeries", func() {
		Description("List series by collection")
		Routing(GET(seriesPath))
		fSeriesList()
	})

	Action("listBooksByPrint", func() {
		Description("List books by collection and print")
		Routing(GET(printPath + printIDPath + bookPath))
		Params(attrPrintID)
		fBookList()
	})

	Action("listSeriesByPrint", func() {
		Description("List series by collection and print")
		Routing(GET(printPath + printIDPath + seriesPath))
		Params(attrPrintID)
		fSeriesList()
	})

	Action("listBooksBySeries", func() {
		Description("List books by collection and series")
		Routing(GET(seriesPath + seriesIDPath + bookPath))
		Params(attrSeriesID)
		fBookList()
	})

	Action("listPrintsBySeries", func() {
		Description("List prints by collection and series")
		Routing(GET(seriesPath + seriesIDPath + printPath))
		Params(attrSeriesID)
		fPrintList()
	})

	Action("listBooksByPrintSeries", func() {
		Description("List books by collection, prints and series")
		Routing(GET(printPath + printIDPath + seriesPath + seriesIDPath + bookPath))
		Params(func() {
			attrPrintID()
			attrSeriesID()
		})
		fBookList()
	})

	Action("listBooksBySeriesPrint", func() {
		Description("List books by collection, series and print")
		Routing(GET(seriesPath + seriesIDPath + printPath + printIDPath + bookPath))
		Params(func() {
			attrPrintID()
			attrSeriesID()
		})
		fBookList()
	})
})
