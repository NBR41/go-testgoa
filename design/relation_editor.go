package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("relationEditors", func() {
	Parent("editors")

	Action("listBooks", func() {
		Description("List books by editor")
		Routing(GET(bookPath))
		fBookList()
	})

	Action("listCollections", func() {
		Description("List collections by editor")
		Routing(GET(collectionPath))
		fCollectionList()
	})

	Action("listPrints", func() {
		Description("List prints by editor")
		Routing(GET(printPath))
		fPrintList()
	})

	Action("listSeries", func() {
		Description("List series by editor")
		Routing(GET(seriesPath))
		fSeriesList()
	})
})

var _ = Resource("relationEditorsPrints", func() {
	BasePath(editorPath + editorIDPath + printPath + printIDPath)
	Params(func() {
		attrEditorID()
		attrPrintID()
	})

	Action("listBooks", func() {
		Description("List books by editor-print")
		Routing(GET(bookPath))
		fBookList()
	})

	Action("listCollections", func() {
		Description("List collections by editor-print")
		Routing(GET(collectionPath))
		fCollectionList()
	})

	Action("listSeries", func() {
		Description("List series by editor-print")
		Routing(GET(seriesPath))
		fSeriesList()
	})

	Action("listBooksByCollection", func() {
		Description("List books by editor-print-collection")
		Routing(GET(collectionPath + collectionIDPath + bookPath))
		Params(attrCollectionID)
		fBookList()
	})

	Action("listSeriesByCollection", func() {
		Description("List series by editor-print-collection")
		Routing(GET(collectionPath + collectionIDPath + seriesPath))
		Params(attrCollectionID)
		fSeriesList()
	})

	Action("listBooksBySeries", func() {
		Description("List books by editor-print-series")
		Routing(GET(seriesPath + seriesIDPath + bookPath))
		Params(attrSeriesID)
		fBookList()
	})

	Action("listCollectionsBySeries", func() {
		Description("List collections by editor-print-series")
		Routing(GET(seriesPath + seriesIDPath + collectionPath))
		Params(attrSeriesID)
		fCollectionList()
	})

	Action("listBooksByCollectionSeries", func() {
		Description("List books by editor-print-collection-series")
		Routing(GET(collectionPath + collectionIDPath + seriesPath + seriesIDPath + bookPath))
		Params(func() {
			attrCollectionID()
			attrSeriesID()
		})
		fBookList()
	})

	Action("listBooksBySeriesCollection", func() {
		Description("List books by editor-print-series-collection")
		Routing(GET(seriesPath + seriesIDPath + collectionPath + collectionIDPath + bookPath))
		Params(func() {
			attrCollectionID()
			attrSeriesID()
		})
		fBookList()
	})
})

var _ = Resource("relationEditorsCollections", func() {
	BasePath(editorPath + editorIDPath + collectionPath + collectionIDPath)
	Params(func() {
		attrEditorID()
		attrCollectionID()
	})

	// /editors/[0-9]+/collections/[0-9]+/books
	Action("listBooks", func() {
		Description("List books by editor-collection")
		Routing(GET(bookPath))
		fBookList()
	})

	// /editors/[0-9]+/collections/[0-9]+/prints
	Action("listPrints", func() {
		Description("List prints by editor-collection")
		Routing(GET(collectionPath))
		fPrintList()
	})

	// /editors/[0-9]+/collections/[0-9]+/series
	Action("listSeries", func() {
		Description("List series by editor-collection")
		Routing(GET(seriesPath))
		fSeriesList()
	})

	// /editors/[0-9]+/collections/[0-9]+/prints/[0-9]+/books
	Action("listBooksByPrint", func() {
		Description("List books by editor-collection-print")
		Routing(GET(printPath + printIDPath + bookPath))
		Params(attrPrintID)
		fBookList()
	})
	// // /editors/[0-9]+/collections/[0-9]+/prints/[0-9]+/series
	Action("listSeriesByPrint", func() {
		Description("List series by editor-collection-print")
		Routing(GET(printPath + printIDPath + seriesPath))
		Params(attrPrintID)
		fSeriesList()
	})

	// /editors/[0-9]+/collections/[0-9]+/series/[0-9]+/books
	Action("listBooksBySeries", func() {
		Description("List books by editor-collection-series")
		Routing(GET(seriesPath + seriesIDPath + bookPath))
		Params(attrSeriesID)
		fBookList()
	})
	// /editors/[0-9]+/collections/[0-9]+/series/[0-9]+/prints
	Action("listPrintsBySeries", func() {
		Description("List prints by editor-collection-series")
		Routing(GET(seriesPath + seriesIDPath + printPath))
		Params(attrSeriesID)
		fPrintList()
	})

	// /editors/[0-9]+/collections/[0-9]+/prints/[0-9]+/series/[0-9]+/books
	Action("listBooksByPrintSeries", func() {
		Description("List books by editor-collection-print-series")
		Routing(GET(printPath + printIDPath + seriesPath + seriesIDPath + bookPath))
		Params(func() {
			attrSeriesID()
			attrPrintID()
		})
		fBookList()
	})

	// /editors/[0-9]+/collections/[0-9]+/series/[0-9]+/prints/[0-9]+/books
	Action("listBooksBySeriesPrint", func() {
		Description("List books by editor-collection-series-print")
		Routing(GET(seriesPath + seriesIDPath + printPath + printIDPath + bookPath))
		Params(func() {
			attrSeriesID()
			attrPrintID()
		})
		fBookList()
	})
})

var _ = Resource("relationEditorsSeries", func() {
	BasePath(editorPath + editorIDPath + seriesPath + seriesIDPath)
	Params(func() {
		attrEditorID()
		attrSeriesID()
	})

	// /editors/[0-9]+/series/[0-9]+/books
	Action("listBooks", func() {
		Description("List books by editor-series")
		Routing(GET(bookPath))
		fBookList()
	})

	// /editors/[0-9]+/series/[0-9]+/collections
	Action("listCollections", func() {
		Description("List collections by editor-series")
		Routing(GET(collectionPath))
		fCollectionList()
	})

	// /editors/[0-9]+/series/[0-9]+/prints
	Action("listPrints", func() {
		Description("List prints by editor-series")
		Routing(GET(printPath))
		fPrintList()
	})

	// /editors/[0-9]+/series/[0-9]+/collections/[0-9]+/books
	Action("listBooksByCollection", func() {
		Description("List books by editor-series-collection")
		Routing(GET(collectionPath + collectionIDPath + bookPath))
		Params(attrCollectionID)
		fBookList()
	})

	// /editors/[0-9]+/series/[0-9]+/collections/[0-9]+/prints
	Action("listPrintsByCollection", func() {
		Description("List prints by editor-series-collection")
		Routing(GET(collectionPath + collectionIDPath + printPath))
		Params(attrCollectionID)
		fPrintList()
	})

	// /editors/[0-9]+/series/[0-9]+/prints/[0-9]+/books
	Action("listBooksByPrint", func() {
		Description("List books by editor-series-print")
		Routing(GET(printPath + printIDPath + bookPath))
		Params(attrPrintID)
		fBookList()
	})

	// /editors/[0-9]+/series/[0-9]+/prints/[0-9]+/collections
	Action("listCollectionsByPrint", func() {
		Description("List collections by editor-series-print")
		Routing(GET(printPath + printIDPath + collectionPath))
		Params(attrPrintID)
		fCollectionList()
	})

	// /editors/[0-9]+/series/[0-9]+/collections/[0-9]+/prints/[0-9]+/books
	Action("listBooksByCollectionPrint", func() {
		Description("List books by editor-series-collection-print")
		Routing(GET(collectionPath + collectionIDPath + printPath + printIDPath + bookPath))
		Params(func() {
			attrPrintID()
			attrCollectionID()
		})
		fBookList()
	})

	// /editors/[0-9]+/series/[0-9]+/prints/[0-9]+/collections/[0-9]+/books
	Action("listBooksByPrintCollection", func() {
		Description("List books by editor-series-print-collection")
		Routing(GET(printPath + printIDPath + collectionPath + collectionIDPath + bookPath))
		Params(func() {
			attrPrintID()
			attrCollectionID()
		})
		fBookList()
	})
})
