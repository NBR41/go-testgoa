// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": relationCategory Resource Client
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

// ListAuthorsRelationCategoryPath computes a request path to the listAuthors action of relationCategory.
func ListAuthorsRelationCategoryPath(categoryID int) string {
	param0 := strconv.Itoa(categoryID)

	return fmt.Sprintf("/categories/%s/authors", param0)
}

// List authors by category
func (c *Client) ListAuthorsRelationCategory(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListAuthorsRelationCategoryRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListAuthorsRelationCategoryRequest create the request corresponding to the listAuthors action endpoint of the relationCategory resource.
func (c *Client) NewListAuthorsRelationCategoryRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListClassesRelationCategoryPath computes a request path to the listClasses action of relationCategory.
func ListClassesRelationCategoryPath(categoryID int) string {
	param0 := strconv.Itoa(categoryID)

	return fmt.Sprintf("/categories/%s/classes", param0)
}

// List classes by category
func (c *Client) ListClassesRelationCategory(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListClassesRelationCategoryRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListClassesRelationCategoryRequest create the request corresponding to the listClasses action endpoint of the relationCategory resource.
func (c *Client) NewListClassesRelationCategoryRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListSeriesRelationCategoryPath computes a request path to the listSeries action of relationCategory.
func ListSeriesRelationCategoryPath(categoryID int) string {
	param0 := strconv.Itoa(categoryID)

	return fmt.Sprintf("/categories/%s/series", param0)
}

// List series by category
func (c *Client) ListSeriesRelationCategory(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListSeriesRelationCategoryRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSeriesRelationCategoryRequest create the request corresponding to the listSeries action endpoint of the relationCategory resource.
func (c *Client) NewListSeriesRelationCategoryRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListSeriesByClassRelationCategoryPath computes a request path to the listSeriesByClass action of relationCategory.
func ListSeriesByClassRelationCategoryPath(categoryID int, classID string) string {
	param0 := strconv.Itoa(categoryID)
	param1 := classID

	return fmt.Sprintf("/categories/%s/classes/%s/series", param0, param1)
}

// List series by category and class
func (c *Client) ListSeriesByClassRelationCategory(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListSeriesByClassRelationCategoryRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSeriesByClassRelationCategoryRequest create the request corresponding to the listSeriesByClass action endpoint of the relationCategory resource.
func (c *Client) NewListSeriesByClassRelationCategoryRequest(ctx context.Context, path string) (*http.Request, error) {
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
