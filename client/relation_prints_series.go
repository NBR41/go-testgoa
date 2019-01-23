// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": relationPrintsSeries Resource Client
//
// Command:
// $ goagen
// --design=github.com/NBR41/go-testgoa/design
// --out=$(GOPATH)/src/github.com/NBR41/go-testgoa
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ListBooksRelationPrintsSeriesPath computes a request path to the listBooks action of relationPrintsSeries.
func ListBooksRelationPrintsSeriesPath(printID int, seriesID int) string {
	param0 := strconv.Itoa(printID)
	param1 := strconv.Itoa(seriesID)

	return fmt.Sprintf("/prints/%s/series/%s/books", param0, param1)
}

// List books by print-series
func (c *Client) ListBooksRelationPrintsSeries(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksRelationPrintsSeriesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksRelationPrintsSeriesRequest create the request corresponding to the listBooks action endpoint of the relationPrintsSeries resource.
func (c *Client) NewListBooksRelationPrintsSeriesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListBooksByCollectionRelationPrintsSeriesPath computes a request path to the listBooksByCollection action of relationPrintsSeries.
func ListBooksByCollectionRelationPrintsSeriesPath(printID int, seriesID int, collectionID int) string {
	param0 := strconv.Itoa(printID)
	param1 := strconv.Itoa(seriesID)
	param2 := strconv.Itoa(collectionID)

	return fmt.Sprintf("/prints/%s/series/%s/collections/%s/books", param0, param1, param2)
}

// List books by print-series-collection
func (c *Client) ListBooksByCollectionRelationPrintsSeries(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksByCollectionRelationPrintsSeriesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksByCollectionRelationPrintsSeriesRequest create the request corresponding to the listBooksByCollection action endpoint of the relationPrintsSeries resource.
func (c *Client) NewListBooksByCollectionRelationPrintsSeriesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListBooksByCollectionEditorRelationPrintsSeriesPath computes a request path to the listBooksByCollectionEditor action of relationPrintsSeries.
func ListBooksByCollectionEditorRelationPrintsSeriesPath(printID int, seriesID int, collectionID int, editorID int) string {
	param0 := strconv.Itoa(printID)
	param1 := strconv.Itoa(seriesID)
	param2 := strconv.Itoa(collectionID)
	param3 := strconv.Itoa(editorID)

	return fmt.Sprintf("/prints/%s/series/%s/collections/%s/editors/%s/books", param0, param1, param2, param3)
}

// List books by print-series-collection-editor
func (c *Client) ListBooksByCollectionEditorRelationPrintsSeries(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksByCollectionEditorRelationPrintsSeriesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksByCollectionEditorRelationPrintsSeriesRequest create the request corresponding to the listBooksByCollectionEditor action endpoint of the relationPrintsSeries resource.
func (c *Client) NewListBooksByCollectionEditorRelationPrintsSeriesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListBooksByEditorRelationPrintsSeriesPath computes a request path to the listBooksByEditor action of relationPrintsSeries.
func ListBooksByEditorRelationPrintsSeriesPath(printID int, seriesID int, editorID int) string {
	param0 := strconv.Itoa(printID)
	param1 := strconv.Itoa(seriesID)
	param2 := strconv.Itoa(editorID)

	return fmt.Sprintf("/prints/%s/series/%s/editors/%s/books", param0, param1, param2)
}

// List books by print-series-editor
func (c *Client) ListBooksByEditorRelationPrintsSeries(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksByEditorRelationPrintsSeriesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksByEditorRelationPrintsSeriesRequest create the request corresponding to the listBooksByEditor action endpoint of the relationPrintsSeries resource.
func (c *Client) NewListBooksByEditorRelationPrintsSeriesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListBooksByEditorCollectionRelationPrintsSeriesPath computes a request path to the listBooksByEditorCollection action of relationPrintsSeries.
func ListBooksByEditorCollectionRelationPrintsSeriesPath(printID int, seriesID int, editorID int, collectionID int) string {
	param0 := strconv.Itoa(printID)
	param1 := strconv.Itoa(seriesID)
	param2 := strconv.Itoa(editorID)
	param3 := strconv.Itoa(collectionID)

	return fmt.Sprintf("/prints/%s/series/%s/editors/%s/collections/%s/books", param0, param1, param2, param3)
}

// List books by print-series-editor-collection
func (c *Client) ListBooksByEditorCollectionRelationPrintsSeries(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksByEditorCollectionRelationPrintsSeriesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksByEditorCollectionRelationPrintsSeriesRequest create the request corresponding to the listBooksByEditorCollection action endpoint of the relationPrintsSeries resource.
func (c *Client) NewListBooksByEditorCollectionRelationPrintsSeriesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListCollectionsRelationPrintsSeriesPath computes a request path to the listCollections action of relationPrintsSeries.
func ListCollectionsRelationPrintsSeriesPath(printID int, seriesID int) string {
	param0 := strconv.Itoa(printID)
	param1 := strconv.Itoa(seriesID)

	return fmt.Sprintf("/prints/%s/series/%s/collections", param0, param1)
}

// List collections by print-series
func (c *Client) ListCollectionsRelationPrintsSeries(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListCollectionsRelationPrintsSeriesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListCollectionsRelationPrintsSeriesRequest create the request corresponding to the listCollections action endpoint of the relationPrintsSeries resource.
func (c *Client) NewListCollectionsRelationPrintsSeriesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListCollectionsByEditorRelationPrintsSeriesPath computes a request path to the listCollectionsByEditor action of relationPrintsSeries.
func ListCollectionsByEditorRelationPrintsSeriesPath(printID int, seriesID int, editorID int) string {
	param0 := strconv.Itoa(printID)
	param1 := strconv.Itoa(seriesID)
	param2 := strconv.Itoa(editorID)

	return fmt.Sprintf("/prints/%s/series/%s/editors/%s/collections", param0, param1, param2)
}

// List collections by print-series-editor
func (c *Client) ListCollectionsByEditorRelationPrintsSeries(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListCollectionsByEditorRelationPrintsSeriesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListCollectionsByEditorRelationPrintsSeriesRequest create the request corresponding to the listCollectionsByEditor action endpoint of the relationPrintsSeries resource.
func (c *Client) NewListCollectionsByEditorRelationPrintsSeriesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListEditorsRelationPrintsSeriesPath computes a request path to the listEditors action of relationPrintsSeries.
func ListEditorsRelationPrintsSeriesPath(printID int, seriesID int) string {
	param0 := strconv.Itoa(printID)
	param1 := strconv.Itoa(seriesID)

	return fmt.Sprintf("/prints/%s/series/%s/editors", param0, param1)
}

// List editors by print-series
func (c *Client) ListEditorsRelationPrintsSeries(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListEditorsRelationPrintsSeriesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListEditorsRelationPrintsSeriesRequest create the request corresponding to the listEditors action endpoint of the relationPrintsSeries resource.
func (c *Client) NewListEditorsRelationPrintsSeriesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListEditorsByCollectionRelationPrintsSeriesPath computes a request path to the listEditorsByCollection action of relationPrintsSeries.
func ListEditorsByCollectionRelationPrintsSeriesPath(printID int, seriesID int, collectionID int) string {
	param0 := strconv.Itoa(printID)
	param1 := strconv.Itoa(seriesID)
	param2 := strconv.Itoa(collectionID)

	return fmt.Sprintf("/prints/%s/series/%s/collections/%s/editors", param0, param1, param2)
}

// List editors by print-series-collection
func (c *Client) ListEditorsByCollectionRelationPrintsSeries(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListEditorsByCollectionRelationPrintsSeriesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListEditorsByCollectionRelationPrintsSeriesRequest create the request corresponding to the listEditorsByCollection action endpoint of the relationPrintsSeries resource.
func (c *Client) NewListEditorsByCollectionRelationPrintsSeriesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
