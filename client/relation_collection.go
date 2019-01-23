// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": relationCollection Resource Client
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

// ListBooksRelationCollectionPath computes a request path to the listBooks action of relationCollection.
func ListBooksRelationCollectionPath(collectionID int) string {
	param0 := strconv.Itoa(collectionID)

	return fmt.Sprintf("/collections/%s/books", param0)
}

// List books by collection
func (c *Client) ListBooksRelationCollection(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksRelationCollectionRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksRelationCollectionRequest create the request corresponding to the listBooks action endpoint of the relationCollection resource.
func (c *Client) NewListBooksRelationCollectionRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListBooksByPrintRelationCollectionPath computes a request path to the listBooksByPrint action of relationCollection.
func ListBooksByPrintRelationCollectionPath(collectionID int, printID int) string {
	param0 := strconv.Itoa(collectionID)
	param1 := strconv.Itoa(printID)

	return fmt.Sprintf("/collections/%s/prints/%s/books", param0, param1)
}

// List books by collection and print
func (c *Client) ListBooksByPrintRelationCollection(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksByPrintRelationCollectionRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksByPrintRelationCollectionRequest create the request corresponding to the listBooksByPrint action endpoint of the relationCollection resource.
func (c *Client) NewListBooksByPrintRelationCollectionRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListBooksByPrintSeriesRelationCollectionPath computes a request path to the listBooksByPrintSeries action of relationCollection.
func ListBooksByPrintSeriesRelationCollectionPath(collectionID int, printID int, seriesID int) string {
	param0 := strconv.Itoa(collectionID)
	param1 := strconv.Itoa(printID)
	param2 := strconv.Itoa(seriesID)

	return fmt.Sprintf("/collections/%s/prints/%s/series/%s/books", param0, param1, param2)
}

// List books by collection, prints and series
func (c *Client) ListBooksByPrintSeriesRelationCollection(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksByPrintSeriesRelationCollectionRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksByPrintSeriesRelationCollectionRequest create the request corresponding to the listBooksByPrintSeries action endpoint of the relationCollection resource.
func (c *Client) NewListBooksByPrintSeriesRelationCollectionRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListBooksBySeriesRelationCollectionPath computes a request path to the listBooksBySeries action of relationCollection.
func ListBooksBySeriesRelationCollectionPath(collectionID int, seriesID int) string {
	param0 := strconv.Itoa(collectionID)
	param1 := strconv.Itoa(seriesID)

	return fmt.Sprintf("/collections/%s/series/%s/books", param0, param1)
}

// List books by collection and series
func (c *Client) ListBooksBySeriesRelationCollection(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksBySeriesRelationCollectionRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksBySeriesRelationCollectionRequest create the request corresponding to the listBooksBySeries action endpoint of the relationCollection resource.
func (c *Client) NewListBooksBySeriesRelationCollectionRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListBooksBySeriesPrintRelationCollectionPath computes a request path to the listBooksBySeriesPrint action of relationCollection.
func ListBooksBySeriesPrintRelationCollectionPath(collectionID int, seriesID int, printID int) string {
	param0 := strconv.Itoa(collectionID)
	param1 := strconv.Itoa(seriesID)
	param2 := strconv.Itoa(printID)

	return fmt.Sprintf("/collections/%s/series/%s/prints/%s/books", param0, param1, param2)
}

// List books by collection, series and print
func (c *Client) ListBooksBySeriesPrintRelationCollection(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksBySeriesPrintRelationCollectionRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksBySeriesPrintRelationCollectionRequest create the request corresponding to the listBooksBySeriesPrint action endpoint of the relationCollection resource.
func (c *Client) NewListBooksBySeriesPrintRelationCollectionRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListPrintsRelationCollectionPath computes a request path to the listPrints action of relationCollection.
func ListPrintsRelationCollectionPath(collectionID int) string {
	param0 := strconv.Itoa(collectionID)

	return fmt.Sprintf("/collections/%s/prints", param0)
}

// List prints by collection
func (c *Client) ListPrintsRelationCollection(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListPrintsRelationCollectionRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListPrintsRelationCollectionRequest create the request corresponding to the listPrints action endpoint of the relationCollection resource.
func (c *Client) NewListPrintsRelationCollectionRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListPrintsBySeriesRelationCollectionPath computes a request path to the listPrintsBySeries action of relationCollection.
func ListPrintsBySeriesRelationCollectionPath(collectionID int, seriesID int) string {
	param0 := strconv.Itoa(collectionID)
	param1 := strconv.Itoa(seriesID)

	return fmt.Sprintf("/collections/%s/series/%s/prints", param0, param1)
}

// List prints by collection and series
func (c *Client) ListPrintsBySeriesRelationCollection(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListPrintsBySeriesRelationCollectionRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListPrintsBySeriesRelationCollectionRequest create the request corresponding to the listPrintsBySeries action endpoint of the relationCollection resource.
func (c *Client) NewListPrintsBySeriesRelationCollectionRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListSeriesRelationCollectionPath computes a request path to the listSeries action of relationCollection.
func ListSeriesRelationCollectionPath(collectionID int) string {
	param0 := strconv.Itoa(collectionID)

	return fmt.Sprintf("/collections/%s/series", param0)
}

// List series by collection
func (c *Client) ListSeriesRelationCollection(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListSeriesRelationCollectionRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSeriesRelationCollectionRequest create the request corresponding to the listSeries action endpoint of the relationCollection resource.
func (c *Client) NewListSeriesRelationCollectionRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListSeriesByPrintRelationCollectionPath computes a request path to the listSeriesByPrint action of relationCollection.
func ListSeriesByPrintRelationCollectionPath(collectionID int, printID int) string {
	param0 := strconv.Itoa(collectionID)
	param1 := strconv.Itoa(printID)

	return fmt.Sprintf("/collections/%s/prints/%s/series", param0, param1)
}

// List series by collection and print
func (c *Client) ListSeriesByPrintRelationCollection(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListSeriesByPrintRelationCollectionRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSeriesByPrintRelationCollectionRequest create the request corresponding to the listSeriesByPrint action endpoint of the relationCollection resource.
func (c *Client) NewListSeriesByPrintRelationCollectionRequest(ctx context.Context, path string) (*http.Request, error) {
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
