package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("relationSeries", func() {
	Parent("series")

	// /series/[0-9]+/books
	Action("listBooks", func() {
		Description("List books by series")
		Routing(GET(bookPath))
		fBookList()
	})

	// /series/[0-9]+/collections
	Action("listCollections", func() {
		Description("List collections by series")
		Routing(GET(collectionPath))
		fCollectionList()
	})

	// /series/[0-9]+/editors
	Action("listEditors", func() {
		Description("List editors by series")
		Routing(GET(editorPath))
		fEditorList()
	})

	// /series/[0-9]+/prints
	Action("listPrints", func() {
		Description("List prints by series")
		Routing(GET(printPath))
		fPrintList()
	})
})

var _ = Resource("relationSeriesCollections", func() {
	BasePath(seriesPath + seriesIDPath + collectionPath + collectionIDPath)
	Params(func() {
		attrSeriesID()
		attrCollectionID()
	})

	// /series/[0-9]+/collections/[0-9]+/books
	Action("listBooks", func() {
		Description("List books by series-collection")
		Routing(GET(bookPath))
		fBookList()
	})

	// /series/[0-9]+/collections/[0-9]+/prints
	Action("listPrints", func() {
		Description("List prints by series-collection")
		Routing(GET(printPath))
		fPrintList()
	})

	// /series/[0-9]+/collections/[0-9]+/prints/[0-9]+/books
	Action("listBooksByPrint", func() {
		Description("List books by series-collection-print")
		Routing(GET(printPath + printIDPath + bookPath))
		Params(attrPrintID)
		fBookList()
	})
})

var _ = Resource("relationSeriesEditors", func() {
	BasePath(seriesPath + seriesIDPath + editorPath + editorIDPath)
	Params(func() {
		attrSeriesID()
		attrEditorID()
	})

	// /series/[0-9]+/editors/[0-9]+/books
	Action("listBooks", func() {
		Description("List books by series-editor")
		Routing(GET(bookPath))
		fBookList()
	})

	// /series/[0-9]+/editors/[0-9]+/collections
	Action("listCollections", func() {
		Description("List collections by series-editor")
		Routing(GET(collectionPath))
		fCollectionList()
	})

	// /series/[0-9]+/editors/[0-9]+/prints
	Action("listPrints", func() {
		Description("List prints by series-editor")
		Routing(GET(printPath))
		fPrintList()
	})

	// /series/[0-9]+/editors/[0-9]+/collections/[0-9]+/books
	Action("listBooksByCollection", func() {
		Description("List books by series-editor-collection")
		Routing(GET(collectionPath + collectionIDPath + bookPath))
		Params(attrCollectionID)
		fBookList()
	})

	// /series/[0-9]+/editors/[0-9]+/collections/[0-9]+/prints
	Action("listPrintsByCollection", func() {
		Description("List prints by series-editor-collection")
		Routing(GET(collectionPath + collectionIDPath + printPath))
		Params(attrCollectionID)
		fPrintList()
	})

	// /series/[0-9]+/editors/[0-9]+/prints/[0-9]+/books
	Action("listBooksByPrint", func() {
		Description("List books by series-editor-print")
		Routing(GET(printPath + printIDPath + bookPath))
		Params(attrPrintID)
		fBookList()
	})

	// /series/[0-9]+/editors/[0-9]+/prints/[0-9]+/collections
	Action("listCollectionsByPrint", func() {
		Description("List collections by series-editor-print")
		Routing(GET(printPath + printIDPath + collectionPath))
		Params(attrPrintID)
		fCollectionList()
	})

	// /series/[0-9]+/editors/[0-9]+/collections/[0-9]+/prints/[0-9]+/books
	Action("listBooksByCollectionPrint", func() {
		Description("List books by series-editor-collection-print")
		Routing(GET(collectionPath + collectionIDPath + printPath + printIDPath + bookPath))
		Params(func() {
			attrCollectionID()
			attrPrintID()
		})
		fBookList()
	})

	// /series/[0-9]+/editors/[0-9]+/prints/[0-9]+/collections/[0-9]+/books
	Action("listBooksByPrintCollection", func() {
		Description("List books by series-editor-print-collection")
		Routing(GET(printPath + printIDPath + collectionPath + collectionIDPath + bookPath))
		Params(func() {
			attrCollectionID()
			attrPrintID()
		})
		fBookList()
	})
})

var _ = Resource("relationSeriesPrints", func() {
	BasePath(seriesPath + seriesIDPath + printPath + printIDPath)
	Params(func() {
		attrSeriesID()
		attrPrintID()
	})

	// /series/[0-9]+/prints/[0-9]+/books
	Action("listBooks", func() {
		Description("List books by series-print")
		Routing(GET(bookPath))
		fBookList()
	})

	// /series/[0-9]+/prints/[0-9]+/collections
	Action("listCollections", func() {
		Description("List collections by series-print")
		Routing(GET(collectionPath))
		fCollectionList()
	})

	// /series/[0-9]+/prints/[0-9]+/editors
	Action("listEditors", func() {
		Description("List editors by series-print")
		Routing(GET(editorPath))
		fEditorList()
	})

	// /series/[0-9]+/prints/[0-9]+/editors/[0-9]+/books
	Action("listBooksByEditor", func() {
		Description("List books by series-print-editor")
		Routing(GET(editorPath + editorIDPath + bookPath))
		Params(attrEditorID)
		fBookList()
	})

	// /series/[0-9]+/prints/[0-9]+/editors/[0-9]+/collections
	Action("listCollectionsByEditor", func() {
		Description("List collections by series-print-editor")
		Routing(GET(editorPath + editorIDPath + collectionPath))
		Params(attrEditorID)
		fCollectionList()
	})

	// /series/[0-9]+/prints/[0-9]+/editors/[0-9]+/collections/[0-9]+/books
	Action("listBooksByEditorCollection", func() {
		Description("List books by series-print-editor-collection")
		Routing(GET(editorPath + editorIDPath + collectionPath + collectionIDPath + bookPath))
		Params(func() {
			attrEditorID()
			attrCollectionID()
		})
		fBookList()
	})
})
