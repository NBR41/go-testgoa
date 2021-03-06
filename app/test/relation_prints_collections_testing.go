// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": relationPrintsCollections TestHelpers
//
// Command:
// $ goagen
// --design=github.com/NBR41/go-testgoa/design
// --out=$(GOPATH)/src/github.com/NBR41/go-testgoa
// --version=v1.3.1

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// ListBooksRelationPrintsCollectionsInternalServerError runs the method ListBooks of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListBooksRelationPrintsCollectionsInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/books", printID, collectionID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listBooksCtx, _err := app.NewListBooksRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.ListBooks(listBooksCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}

	// Return results
	return rw
}

// ListBooksRelationPrintsCollectionsNotFound runs the method ListBooks of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListBooksRelationPrintsCollectionsNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/books", printID, collectionID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listBooksCtx, _err := app.NewListBooksRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.ListBooks(listBooksCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// ListBooksRelationPrintsCollectionsOK runs the method ListBooks of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListBooksRelationPrintsCollectionsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int) (http.ResponseWriter, app.BookCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/books", printID, collectionID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listBooksCtx, _err := app.NewListBooksRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.ListBooks(listBooksCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.BookCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.BookCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.BookCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ListBooksRelationPrintsCollectionsOKLink runs the method ListBooks of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListBooksRelationPrintsCollectionsOKLink(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int) (http.ResponseWriter, app.BookLinkCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/books", printID, collectionID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listBooksCtx, _err := app.NewListBooksRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.ListBooks(listBooksCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.BookLinkCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.BookLinkCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.BookLinkCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ListBooksRelationPrintsCollectionsServiceUnavailable runs the method ListBooks of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListBooksRelationPrintsCollectionsServiceUnavailable(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/books", printID, collectionID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listBooksCtx, _err := app.NewListBooksRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.ListBooks(listBooksCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 503 {
		t.Errorf("invalid response status code: got %+v, expected 503", rw.Code)
	}

	// Return results
	return rw
}

// ListBooksBySeriesRelationPrintsCollectionsInternalServerError runs the method ListBooksBySeries of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListBooksBySeriesRelationPrintsCollectionsInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int, seriesID int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/series/%v/books", printID, collectionID, seriesID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	prms["series_id"] = []string{fmt.Sprintf("%v", seriesID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listBooksBySeriesCtx, _err := app.NewListBooksBySeriesRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.ListBooksBySeries(listBooksBySeriesCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}

	// Return results
	return rw
}

// ListBooksBySeriesRelationPrintsCollectionsNotFound runs the method ListBooksBySeries of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListBooksBySeriesRelationPrintsCollectionsNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int, seriesID int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/series/%v/books", printID, collectionID, seriesID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	prms["series_id"] = []string{fmt.Sprintf("%v", seriesID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listBooksBySeriesCtx, _err := app.NewListBooksBySeriesRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.ListBooksBySeries(listBooksBySeriesCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// ListBooksBySeriesRelationPrintsCollectionsOK runs the method ListBooksBySeries of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListBooksBySeriesRelationPrintsCollectionsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int, seriesID int) (http.ResponseWriter, app.BookCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/series/%v/books", printID, collectionID, seriesID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	prms["series_id"] = []string{fmt.Sprintf("%v", seriesID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listBooksBySeriesCtx, _err := app.NewListBooksBySeriesRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.ListBooksBySeries(listBooksBySeriesCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.BookCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.BookCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.BookCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ListBooksBySeriesRelationPrintsCollectionsOKLink runs the method ListBooksBySeries of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListBooksBySeriesRelationPrintsCollectionsOKLink(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int, seriesID int) (http.ResponseWriter, app.BookLinkCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/series/%v/books", printID, collectionID, seriesID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	prms["series_id"] = []string{fmt.Sprintf("%v", seriesID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listBooksBySeriesCtx, _err := app.NewListBooksBySeriesRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.ListBooksBySeries(listBooksBySeriesCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.BookLinkCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.BookLinkCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.BookLinkCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ListBooksBySeriesRelationPrintsCollectionsServiceUnavailable runs the method ListBooksBySeries of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListBooksBySeriesRelationPrintsCollectionsServiceUnavailable(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int, seriesID int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/series/%v/books", printID, collectionID, seriesID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	prms["series_id"] = []string{fmt.Sprintf("%v", seriesID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listBooksBySeriesCtx, _err := app.NewListBooksBySeriesRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.ListBooksBySeries(listBooksBySeriesCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 503 {
		t.Errorf("invalid response status code: got %+v, expected 503", rw.Code)
	}

	// Return results
	return rw
}

// ListSeriesRelationPrintsCollectionsInternalServerError runs the method ListSeries of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListSeriesRelationPrintsCollectionsInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/series", printID, collectionID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listSeriesCtx, _err := app.NewListSeriesRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.ListSeries(listSeriesCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}

	// Return results
	return rw
}

// ListSeriesRelationPrintsCollectionsNotFound runs the method ListSeries of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListSeriesRelationPrintsCollectionsNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/series", printID, collectionID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listSeriesCtx, _err := app.NewListSeriesRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.ListSeries(listSeriesCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// ListSeriesRelationPrintsCollectionsOK runs the method ListSeries of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListSeriesRelationPrintsCollectionsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int) (http.ResponseWriter, app.SeriesCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/series", printID, collectionID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listSeriesCtx, _err := app.NewListSeriesRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.ListSeries(listSeriesCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.SeriesCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.SeriesCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.SeriesCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ListSeriesRelationPrintsCollectionsOKLink runs the method ListSeries of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListSeriesRelationPrintsCollectionsOKLink(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int) (http.ResponseWriter, app.SeriesLinkCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/series", printID, collectionID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listSeriesCtx, _err := app.NewListSeriesRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.ListSeries(listSeriesCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.SeriesLinkCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.SeriesLinkCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.SeriesLinkCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ListSeriesRelationPrintsCollectionsServiceUnavailable runs the method ListSeries of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListSeriesRelationPrintsCollectionsServiceUnavailable(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RelationPrintsCollectionsController, printID int, collectionID int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/prints/%v/collections/%v/series", printID, collectionID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["print_id"] = []string{fmt.Sprintf("%v", printID)}
	prms["collection_id"] = []string{fmt.Sprintf("%v", collectionID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RelationPrintsCollectionsTest"), rw, req, prms)
	listSeriesCtx, _err := app.NewListSeriesRelationPrintsCollectionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.ListSeries(listSeriesCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 503 {
		t.Errorf("invalid response status code: got %+v, expected 503", rw.Code)
	}

	// Return results
	return rw
}
