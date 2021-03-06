// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": relationPrintsCollections Resource Client
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

// ListBooksRelationPrintsCollectionsPath computes a request path to the listBooks action of relationPrintsCollections.
func ListBooksRelationPrintsCollectionsPath(printID int, collectionID int) string {
	param0 := strconv.Itoa(printID)
	param1 := strconv.Itoa(collectionID)

	return fmt.Sprintf("/prints/%s/collections/%s/books", param0, param1)
}

// List books by print-collection
func (c *Client) ListBooksRelationPrintsCollections(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksRelationPrintsCollectionsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksRelationPrintsCollectionsRequest create the request corresponding to the listBooks action endpoint of the relationPrintsCollections resource.
func (c *Client) NewListBooksRelationPrintsCollectionsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListBooksBySeriesRelationPrintsCollectionsPath computes a request path to the listBooksBySeries action of relationPrintsCollections.
func ListBooksBySeriesRelationPrintsCollectionsPath(printID int, collectionID int, seriesID int) string {
	param0 := strconv.Itoa(printID)
	param1 := strconv.Itoa(collectionID)
	param2 := strconv.Itoa(seriesID)

	return fmt.Sprintf("/prints/%s/collections/%s/series/%s/books", param0, param1, param2)
}

// List books by print-collection-series
func (c *Client) ListBooksBySeriesRelationPrintsCollections(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksBySeriesRelationPrintsCollectionsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksBySeriesRelationPrintsCollectionsRequest create the request corresponding to the listBooksBySeries action endpoint of the relationPrintsCollections resource.
func (c *Client) NewListBooksBySeriesRelationPrintsCollectionsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListSeriesRelationPrintsCollectionsPath computes a request path to the listSeries action of relationPrintsCollections.
func ListSeriesRelationPrintsCollectionsPath(printID int, collectionID int) string {
	param0 := strconv.Itoa(printID)
	param1 := strconv.Itoa(collectionID)

	return fmt.Sprintf("/prints/%s/collections/%s/series", param0, param1)
}

// List series by print-collection
func (c *Client) ListSeriesRelationPrintsCollections(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListSeriesRelationPrintsCollectionsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSeriesRelationPrintsCollectionsRequest create the request corresponding to the listSeries action endpoint of the relationPrintsCollections resource.
func (c *Client) NewListSeriesRelationPrintsCollectionsRequest(ctx context.Context, path string) (*http.Request, error) {
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
