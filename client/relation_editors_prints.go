// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": relationEditorsPrints Resource Client
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

// ListBooksRelationEditorsPrintsPath computes a request path to the listBooks action of relationEditorsPrints.
func ListBooksRelationEditorsPrintsPath(editorID int, printID int) string {
	param0 := strconv.Itoa(editorID)
	param1 := strconv.Itoa(printID)

	return fmt.Sprintf("/editors/%s/prints/%s/books", param0, param1)
}

// List books by editor-print
func (c *Client) ListBooksRelationEditorsPrints(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksRelationEditorsPrintsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksRelationEditorsPrintsRequest create the request corresponding to the listBooks action endpoint of the relationEditorsPrints resource.
func (c *Client) NewListBooksRelationEditorsPrintsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListBooksByCollectionRelationEditorsPrintsPath computes a request path to the listBooksByCollection action of relationEditorsPrints.
func ListBooksByCollectionRelationEditorsPrintsPath(editorID int, printID int, collectionID int) string {
	param0 := strconv.Itoa(editorID)
	param1 := strconv.Itoa(printID)
	param2 := strconv.Itoa(collectionID)

	return fmt.Sprintf("/editors/%s/prints/%s/collections/%s/books", param0, param1, param2)
}

// List books by editor-print-collection
func (c *Client) ListBooksByCollectionRelationEditorsPrints(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksByCollectionRelationEditorsPrintsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksByCollectionRelationEditorsPrintsRequest create the request corresponding to the listBooksByCollection action endpoint of the relationEditorsPrints resource.
func (c *Client) NewListBooksByCollectionRelationEditorsPrintsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListBooksByCollectionSeriesRelationEditorsPrintsPath computes a request path to the listBooksByCollectionSeries action of relationEditorsPrints.
func ListBooksByCollectionSeriesRelationEditorsPrintsPath(editorID int, printID int, collectionID int, seriesID int) string {
	param0 := strconv.Itoa(editorID)
	param1 := strconv.Itoa(printID)
	param2 := strconv.Itoa(collectionID)
	param3 := strconv.Itoa(seriesID)

	return fmt.Sprintf("/editors/%s/prints/%s/collections/%s/series/%s/books", param0, param1, param2, param3)
}

// List books by editor-print-collection-series
func (c *Client) ListBooksByCollectionSeriesRelationEditorsPrints(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksByCollectionSeriesRelationEditorsPrintsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksByCollectionSeriesRelationEditorsPrintsRequest create the request corresponding to the listBooksByCollectionSeries action endpoint of the relationEditorsPrints resource.
func (c *Client) NewListBooksByCollectionSeriesRelationEditorsPrintsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListBooksBySeriesRelationEditorsPrintsPath computes a request path to the listBooksBySeries action of relationEditorsPrints.
func ListBooksBySeriesRelationEditorsPrintsPath(editorID int, printID int, seriesID int) string {
	param0 := strconv.Itoa(editorID)
	param1 := strconv.Itoa(printID)
	param2 := strconv.Itoa(seriesID)

	return fmt.Sprintf("/editors/%s/prints/%s/series/%s/books", param0, param1, param2)
}

// List books by editor-print-series
func (c *Client) ListBooksBySeriesRelationEditorsPrints(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksBySeriesRelationEditorsPrintsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksBySeriesRelationEditorsPrintsRequest create the request corresponding to the listBooksBySeries action endpoint of the relationEditorsPrints resource.
func (c *Client) NewListBooksBySeriesRelationEditorsPrintsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListBooksBySeriesCollectionRelationEditorsPrintsPath computes a request path to the listBooksBySeriesCollection action of relationEditorsPrints.
func ListBooksBySeriesCollectionRelationEditorsPrintsPath(editorID int, printID int, seriesID int, collectionID int) string {
	param0 := strconv.Itoa(editorID)
	param1 := strconv.Itoa(printID)
	param2 := strconv.Itoa(seriesID)
	param3 := strconv.Itoa(collectionID)

	return fmt.Sprintf("/editors/%s/prints/%s/series/%s/collections/%s/books", param0, param1, param2, param3)
}

// List books by editor-print-series-collection
func (c *Client) ListBooksBySeriesCollectionRelationEditorsPrints(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksBySeriesCollectionRelationEditorsPrintsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksBySeriesCollectionRelationEditorsPrintsRequest create the request corresponding to the listBooksBySeriesCollection action endpoint of the relationEditorsPrints resource.
func (c *Client) NewListBooksBySeriesCollectionRelationEditorsPrintsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListCollectionsRelationEditorsPrintsPath computes a request path to the listCollections action of relationEditorsPrints.
func ListCollectionsRelationEditorsPrintsPath(editorID int, printID int) string {
	param0 := strconv.Itoa(editorID)
	param1 := strconv.Itoa(printID)

	return fmt.Sprintf("/editors/%s/prints/%s/collections", param0, param1)
}

// List collections by editor-print
func (c *Client) ListCollectionsRelationEditorsPrints(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListCollectionsRelationEditorsPrintsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListCollectionsRelationEditorsPrintsRequest create the request corresponding to the listCollections action endpoint of the relationEditorsPrints resource.
func (c *Client) NewListCollectionsRelationEditorsPrintsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListCollectionsBySeriesRelationEditorsPrintsPath computes a request path to the listCollectionsBySeries action of relationEditorsPrints.
func ListCollectionsBySeriesRelationEditorsPrintsPath(editorID int, printID int, seriesID int) string {
	param0 := strconv.Itoa(editorID)
	param1 := strconv.Itoa(printID)
	param2 := strconv.Itoa(seriesID)

	return fmt.Sprintf("/editors/%s/prints/%s/series/%s/collections", param0, param1, param2)
}

// List collections by editor-print-series
func (c *Client) ListCollectionsBySeriesRelationEditorsPrints(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListCollectionsBySeriesRelationEditorsPrintsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListCollectionsBySeriesRelationEditorsPrintsRequest create the request corresponding to the listCollectionsBySeries action endpoint of the relationEditorsPrints resource.
func (c *Client) NewListCollectionsBySeriesRelationEditorsPrintsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListSeriesRelationEditorsPrintsPath computes a request path to the listSeries action of relationEditorsPrints.
func ListSeriesRelationEditorsPrintsPath(editorID int, printID int) string {
	param0 := strconv.Itoa(editorID)
	param1 := strconv.Itoa(printID)

	return fmt.Sprintf("/editors/%s/prints/%s/series", param0, param1)
}

// List series by editor-print
func (c *Client) ListSeriesRelationEditorsPrints(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListSeriesRelationEditorsPrintsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSeriesRelationEditorsPrintsRequest create the request corresponding to the listSeries action endpoint of the relationEditorsPrints resource.
func (c *Client) NewListSeriesRelationEditorsPrintsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListSeriesByCollectionRelationEditorsPrintsPath computes a request path to the listSeriesByCollection action of relationEditorsPrints.
func ListSeriesByCollectionRelationEditorsPrintsPath(editorID int, printID int, collectionID int) string {
	param0 := strconv.Itoa(editorID)
	param1 := strconv.Itoa(printID)
	param2 := strconv.Itoa(collectionID)

	return fmt.Sprintf("/editors/%s/prints/%s/collections/%s/series", param0, param1, param2)
}

// List series by editor-print-collection
func (c *Client) ListSeriesByCollectionRelationEditorsPrints(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListSeriesByCollectionRelationEditorsPrintsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSeriesByCollectionRelationEditorsPrintsRequest create the request corresponding to the listSeriesByCollection action endpoint of the relationEditorsPrints resource.
func (c *Client) NewListSeriesByCollectionRelationEditorsPrintsRequest(ctx context.Context, path string) (*http.Request, error) {
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
