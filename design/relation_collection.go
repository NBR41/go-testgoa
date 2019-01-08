package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("relationCollection", func() {
	Parent("collections")

	Action("listBooks", func() {
		Description("List books by collection")
		Routing(GET(bookPath))
		// ok
		Response(OK, CollectionOf(BookMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("listPrints", func() {
		Description("List prints by collection")
		Routing(GET(printPath))
		// ok
		Response(OK, CollectionOf(PrintMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("listSeries", func() {
		Description("List series by collection")
		Routing(GET(seriesPath))
		// ok
		Response(OK, CollectionOf(SeriesMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("listBooksByPrint", func() {
		Description("List books by collection and print")
		Routing(GET(printPath + printIDPath + bookPath))
		// ok
		Response(OK, CollectionOf(BookMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("listSeriesByPrint", func() {
		Description("List series by collection and print")
		Routing(GET(printPath + printIDPath + seriesPath))
		// ok
		Response(OK, CollectionOf(SeriesMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("listBooksBySeries", func() {
		Description("List books by collection and series")
		Routing(GET(seriesPath + seriesIDPath + bookPath))
		// ok
		Response(OK, CollectionOf(BookMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("listPrintsBySeries", func() {
		Description("List prints by collection and series")
		Routing(GET(seriesPath + seriesIDPath + printPath))
		// ok
		Response(OK, CollectionOf(PrintMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("listBooksByPrintsSeries", func() {
		Description("List books by collection, prints and series")
		Routing(GET(printPath + printIDPath + seriesPath + seriesIDPath + bookPath))
		// ok
		Response(OK, CollectionOf(BookMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("listBooksBySeriesPrints", func() {
		Description("List books by collection, series and prints")
		Routing(GET(seriesPath + seriesIDPath + printPath + printIDPath + bookPath))
		// ok
		Response(OK, CollectionOf(BookMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
})