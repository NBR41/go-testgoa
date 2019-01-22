package controllers

import (
	"context"
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

type listResponse interface {
	NotFound() error
	InternalServerError() error
	ServiceUnavailable() error
}

type authorsResponse interface {
	OK(r app.AuthorCollection) error
	listResponse
}

type booksResponse interface {
	OK(r app.BookCollection) error
	listResponse
}

type categoriesResponse interface {
	OK(r app.CategoryCollection) error
	listResponse
}

type classesResponse interface {
	OK(r app.ClassCollection) error
	listResponse
}

type collectionsResponse interface {
	OK(r app.CollectionCollection) error
	listResponse
}

type editorsResponse interface {
	OK(r app.EditorCollection) error
	listResponse
}

type printsResponse interface {
	OK(r app.PrintCollection) error
	listResponse
}

type seriesResponse interface {
	OK(r app.SeriesCollection) error
	listResponse
}

type ListBuilder struct {
}

//ListAuthors process request list for authors
func (ListBuilder) ListAuthors(ctx context.Context, fm Fmodeler, rCtx authorsResponse, categoryID, roleID *int) error {
	m, err := fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return rCtx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	if categoryID != nil {
		_, err = m.GetCategoryByID(*categoryID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get category`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if roleID != nil {
		_, err = m.GetRoleByID(*roleID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get role`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	list, err := m.ListAuthorsByIDs(categoryID, roleID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get author list`, `error`, err.Error())
		return rCtx.InternalServerError()
	}
	bs := make(app.AuthorCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToAuthorMedia(bk)
	}
	return rCtx.OK(bs)
}

//ListBooks process request list for books
func (ListBuilder) ListBooks(ctx context.Context, fm Fmodeler, rCtx booksResponse, collectionID, editorID, printID, seriesID *int) error {
	m, err := fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return rCtx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	if collectionID != nil {
		_, err = m.GetCollectionByID(*collectionID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get collection`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if editorID != nil {
		_, err = m.GetEditorByID(*editorID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get editor`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if printID != nil {
		_, err = m.GetPrintByID(*printID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get print`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if seriesID != nil {
		_, err = m.GetSeriesByID(*seriesID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get series`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	list, err := m.ListBooksByIDs(collectionID, editorID, printID, seriesID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get book list`, `error`, err.Error())
		return rCtx.InternalServerError()
	}
	bs := make(app.BookCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToBookMedia(bk)
	}
	return rCtx.OK(bs)
}

func (ListBuilder) ListCategories(ctx context.Context, fm Fmodeler, rCtx categoriesResponse, authorID, classID *int) error {
	m, err := fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return rCtx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	if authorID != nil {
		_, err = m.GetAuthorByID(*authorID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get author`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if classID != nil {
		_, err = m.GetClassByID(*classID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get class`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	list, err := m.ListCategoriesByIDs(authorID, classID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get category list`, `error`, err.Error())
		return rCtx.InternalServerError()
	}
	bs := make(app.CategoryCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToCategoryMedia(bk)
	}
	return rCtx.OK(bs)
}

//ListClasses process request list for classes
func (ListBuilder) ListClasses(ctx context.Context, fm Fmodeler, rCtx classesResponse, authorID, categoryID, seriesID *int) error {
	m, err := fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return rCtx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	if authorID != nil {
		_, err = m.GetAuthorByID(*authorID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get author`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if categoryID != nil {
		_, err = m.GetCategoryByID(*categoryID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get category`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if seriesID != nil {
		_, err = m.GetSeriesByID(*seriesID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get series`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	list, err := m.ListClassesByIDs(authorID, categoryID, seriesID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get class list`, `error`, err.Error())
		return rCtx.InternalServerError()
	}
	bs := make(app.ClassCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToClassMedia(bk)
	}
	return rCtx.OK(bs)
}

//ListCollections process request list for collections
func (ListBuilder) ListCollections(ctx context.Context, fm Fmodeler, rCtx collectionsResponse, editorID, printID, seriesID *int) error {
	m, err := fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return rCtx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	if editorID != nil {
		_, err = m.GetEditorByID(*editorID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get editor`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if printID != nil {
		_, err = m.GetPrintByID(*printID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get print`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if seriesID != nil {
		_, err = m.GetSeriesByID(*seriesID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get series`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	list, err := m.ListCollectionsByIDs(editorID, printID, seriesID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get collection list`, `error`, err.Error())
		return rCtx.InternalServerError()
	}
	bs := make(app.CollectionCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToCollectionMedia(bk)
	}
	return rCtx.OK(bs)
}

//ListEditors process request list for editors
func (ListBuilder) ListEditors(ctx context.Context, fm Fmodeler, rCtx editorsResponse, printID, seriesID *int) error {
	m, err := fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return rCtx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	if printID != nil {
		_, err = m.GetPrintByID(*printID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get print`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if seriesID != nil {
		_, err = m.GetSeriesByID(*seriesID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get series`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	list, err := m.ListEditorsByIDs(printID, seriesID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get editor list`, `error`, err.Error())
		return rCtx.InternalServerError()
	}
	bs := make(app.EditorCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToEditorMedia(bk)
	}
	return rCtx.OK(bs)
}

//ListPrints process request list for prints
func (ListBuilder) ListPrints(ctx context.Context, fm Fmodeler, rCtx printsResponse, collectionID, editorID, seriesID *int) error {
	m, err := fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return rCtx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	if collectionID != nil {
		_, err = m.GetCollectionByID(*collectionID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get collection`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if editorID != nil {
		_, err = m.GetEditorByID(*editorID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get editor`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if seriesID != nil {
		_, err = m.GetSeriesByID(*seriesID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get series`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	list, err := m.ListPrintsByIDs(collectionID, editorID, seriesID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get print list`, `error`, err.Error())
		return rCtx.InternalServerError()
	}
	bs := make(app.PrintCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToPrintMedia(bk)
	}
	return rCtx.OK(bs)
}

//ListSeries process request list for series
func (ListBuilder) ListSeries(ctx context.Context, fm Fmodeler, rCtx seriesResponse, authorID, categoryID, classID, roleID *int) error {
	m, err := fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return rCtx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	if authorID != nil {
		_, err = m.GetAuthorByID(*authorID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get author`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if categoryID != nil {
		_, err = m.GetCategoryByID(*categoryID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get category`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if classID != nil {
		_, err = m.GetClassByID(*classID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get class`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if roleID != nil {
		_, err = m.GetRoleByID(*roleID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get role`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	list, err := m.ListSeriesByIDs(authorID, categoryID, classID, roleID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get series list`, `error`, err.Error())
		return rCtx.InternalServerError()
	}
	bs := make(app.SeriesCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToSeriesMedia(bk)
	}
	return rCtx.OK(bs)
}

//ListSeriesByEditionIDs process request list for series
func (ListBuilder) ListSeriesByEditionIDs(ctx context.Context, fm Fmodeler, rCtx seriesResponse, collectionID, editorID, printID *int) error {
	m, err := fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return rCtx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	if collectionID != nil {
		_, err = m.GetCollectionByID(*collectionID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get collection`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if editorID != nil {
		_, err = m.GetEditorByID(*editorID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get editor`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	if printID != nil {
		_, err = m.GetPrintByID(*printID)
		if err != nil {
			goa.ContextLogger(ctx).Error(`failed to get print`, `error`, err.Error())
			if err == model.ErrNotFound {
				return rCtx.NotFound()
			}
			return rCtx.InternalServerError()
		}
	}

	list, err := m.ListSeriesByEditionIDs(collectionID, editorID, printID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get series list`, `error`, err.Error())
		return rCtx.InternalServerError()
	}
	bs := make(app.SeriesCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToSeriesMedia(bk)
	}
	return rCtx.OK(bs)
}
