package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("relationPrints", func() {
	Parent("prints")

	Action("listBooks", func() {
		Description("List books by print")
		Routing(GET(bookPath))
		fBookList()
	})

	Action("listCollections", func() {
		Description("List collections by print")
		Routing(GET(collectionPath))
		fCollectionList()
	})

	Action("listEditors", func() {
		Description("List editors by print")
		Routing(GET(editorPath))
		fEditorList()
	})

	Action("listSeries", func() {
		Description("List series by print")
		Routing(GET(seriesPath))
		fSeriesList()
	})
})

var _ = Resource("relationPrintsCollections", func() {
	BasePath(printPath + printIDPath + collectionPath + collectionIDPath)
	Params(func() {
		attrPrintID()
		attrCollectionID()
	})

	// /prints/[0-9]+/collections/[0-9]+/books
	Action("listBooks", func() {
		Description("List books by print-collection")
		Routing(GET(bookPath))
		fBookList()
	})

	// /prints/[0-9]+/collections/[0-9]+/series
	Action("listSeries", func() {
		Description("List series by print-collection")
		Routing(GET(seriesPath))
		fSeriesList()
	})

	// /prints/[0-9]+/collections/[0-9]+/series/[0-9]+/books
	Action("listBooksBySeries", func() {
		Description("List books by print-collection-series")
		Routing(GET(seriesPath + seriesIDPath + bookPath))
		Params(attrSeriesID)
		fBookList()
	})
})

var _ = Resource("relationPrintsEditors", func() {
	BasePath(printPath + printIDPath + editorPath + editorIDPath)
	Params(func() {
		attrEditorID()
		attrPrintID()
	})

	// /prints/[0-9]+/editors/[0-9]+/books
	Action("listBooks", func() {
		Description("List books by print-editor")
		Routing(GET(bookPath))
		fBookList()
	})

	// /prints/[0-9]+/editors/[0-9]+/collections
	Action("listCollections", func() {
		Description("List collections by print-editor")
		Routing(GET(collectionPath))
		fCollectionList()
	})

	// /prints/[0-9]+/editors/[0-9]+/series
	Action("listSeries", func() {
		Description("List series by print-editor")
		Routing(GET(seriesPath))
		fSeriesList()
	})

	// /prints/[0-9]+/editors/[0-9]+/collections/[0-9]+/books
	Action("listBooksByCollection", func() {
		Description("List books by print-editor-collection")
		Routing(GET(collectionPath + collectionIDPath + bookPath))
		Params(attrCollectionID)
		fBookList()
	})

	// /prints/[0-9]+/editors/[0-9]+/collections/[0-9]+/series
	Action("listSeriesByCollection", func() {
		Description("List series by print-editor-collection")
		Routing(GET(collectionPath + collectionIDPath + seriesPath))
		Params(attrCollectionID)
		fSeriesList()
	})

	// /prints/[0-9]+/editors/[0-9]+/series/[0-9]+/books
	Action("listBooksBySeries", func() {
		Description("List books by print-editor-series")
		Routing(GET(seriesPath + seriesIDPath + bookPath))
		Params(attrSeriesID)
		fBookList()
	})

	// /prints/[0-9]+/editors/[0-9]+/series/[0-9]+/collections
	Action("listCollectionsBySeries", func() {
		Description("List collections by print-editor-series")
		Routing(GET(seriesPath + seriesIDPath + collectionPath))
		Params(attrSeriesID)
		fCollectionList()
	})

	// /prints/[0-9]+/editors/[0-9]+/collections/[0-9]+/series/[0-9]+/books
	Action("listBooksByCollectionSeries", func() {
		Description("List books by print-editor-collection-series")
		Routing(GET(collectionPath + collectionIDPath + seriesPath + seriesIDPath + bookPath))
		Params(func() {
			attrCollectionID()
			attrSeriesID()
		})
		fBookList()
	})
	// /prints/[0-9]+/editors/[0-9]+/series/[0-9]+/collections/[0-9]+/books
	Action("listBooksBySeriesCollection", func() {
		Description("List books by print-editor-series-collection")
		Routing(GET(seriesPath + seriesIDPath + collectionPath + collectionIDPath + bookPath))
		Params(func() {
			attrCollectionID()
			attrSeriesID()
		})
		fBookList()
	})
})

var _ = Resource("relationPrintsSeries", func() {
	BasePath(printPath + printIDPath + seriesPath + seriesIDPath)
	Params(func() {
		attrSeriesID()
		attrPrintID()
	})

	// /prints/[0-9]+/series/[0-9]+/books
	Action("listBooks", func() {
		Description("List books by print-series")
		Routing(GET(bookPath))
		fBookList()
	})

	// /prints/[0-9]+/series/[0-9]+/collections
	Action("listCollections", func() {
		Description("List collections by print-series")
		Routing(GET(collectionPath))
		fCollectionList()
	})

	// /prints/[0-9]+/series/[0-9]+/editors
	Action("listEditors", func() {
		Description("List editors by print-series")
		Routing(GET(editorPath))
		fEditorList()
	})

	// /prints/[0-9]+/series/[0-9]+/collections/[0-9]+/books
	Action("listBooksByCollection", func() {
		Description("List books by print-series-collection")
		Routing(GET(collectionPath + collectionIDPath + bookPath))
		Params(attrCollectionID)
		fBookList()
	})

	// /prints/[0-9]+/series/[0-9]+/collections/[0-9]+/editors
	Action("listEditorsByCollection", func() {
		Description("List editors by print-series-collection")
		Routing(GET(collectionPath + collectionIDPath + editorPath))
		Params(attrCollectionID)
		fEditorList()
	})

	// /prints/[0-9]+/series/[0-9]+/editors/[0-9]+/books
	Action("listBooksByEditor", func() {
		Description("List books by print-series-editor")
		Routing(GET(editorPath + editorIDPath + bookPath))
		Params(attrEditorID)
		fBookList()
	})

	// /prints/[0-9]+/series/[0-9]+/editors/[0-9]+/collections
	Action("listCollectionsByEditor", func() {
		Description("List collections by print-series-editor")
		Routing(GET(editorPath + editorIDPath + collectionPath))
		Params(attrEditorID)
		fCollectionList()
	})

	// /prints/[0-9]+/series/[0-9]+/collections/[0-9]+/editors/[0-9]+/books
	Action("listBooksByCollectionEditor", func() {
		Description("List books by print-series-collection-editor")
		Routing(GET(collectionPath + collectionIDPath + editorPath + editorIDPath + bookPath))
		Params(func() {
			attrCollectionID()
			attrEditorID()
		})
		fBookList()
	})

	// /prints/[0-9]+/series/[0-9]+/editors/[0-9]+/collections/[0-9]+/books
	Action("listBooksByEditorCollection", func() {
		Description("List books by print-series-editor-collection")
		Routing(GET(editorPath + editorIDPath + collectionPath + collectionIDPath + bookPath))
		Params(func() {
			attrCollectionID()
			attrEditorID()
		})
		fBookList()
	})
})
