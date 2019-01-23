// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": relationSeries Resource Client
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

// ListBooksRelationSeriesPath computes a request path to the listBooks action of relationSeries.
func ListBooksRelationSeriesPath(seriesID int) string {
	param0 := strconv.Itoa(seriesID)

	return fmt.Sprintf("/series/%s/books", param0)
}

// List books by series
func (c *Client) ListBooksRelationSeries(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksRelationSeriesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksRelationSeriesRequest create the request corresponding to the listBooks action endpoint of the relationSeries resource.
func (c *Client) NewListBooksRelationSeriesRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListCollectionsRelationSeriesPath computes a request path to the listCollections action of relationSeries.
func ListCollectionsRelationSeriesPath(seriesID int) string {
	param0 := strconv.Itoa(seriesID)

	return fmt.Sprintf("/series/%s/collections", param0)
}

// List collections by series
func (c *Client) ListCollectionsRelationSeries(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListCollectionsRelationSeriesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListCollectionsRelationSeriesRequest create the request corresponding to the listCollections action endpoint of the relationSeries resource.
func (c *Client) NewListCollectionsRelationSeriesRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListEditorsRelationSeriesPath computes a request path to the listEditors action of relationSeries.
func ListEditorsRelationSeriesPath(seriesID int) string {
	param0 := strconv.Itoa(seriesID)

	return fmt.Sprintf("/series/%s/editors", param0)
}

// List editors by series
func (c *Client) ListEditorsRelationSeries(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListEditorsRelationSeriesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListEditorsRelationSeriesRequest create the request corresponding to the listEditors action endpoint of the relationSeries resource.
func (c *Client) NewListEditorsRelationSeriesRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListPrintsRelationSeriesPath computes a request path to the listPrints action of relationSeries.
func ListPrintsRelationSeriesPath(seriesID int) string {
	param0 := strconv.Itoa(seriesID)

	return fmt.Sprintf("/series/%s/prints", param0)
}

// List prints by series
func (c *Client) ListPrintsRelationSeries(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListPrintsRelationSeriesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListPrintsRelationSeriesRequest create the request corresponding to the listPrints action endpoint of the relationSeries resource.
func (c *Client) NewListPrintsRelationSeriesRequest(ctx context.Context, path string) (*http.Request, error) {
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