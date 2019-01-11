// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": relationAuthor Resource Client
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

// ListCategoriesRelationAuthorPath computes a request path to the listCategories action of relationAuthor.
func ListCategoriesRelationAuthorPath(authorID int) string {
	param0 := strconv.Itoa(authorID)

	return fmt.Sprintf("/authors/%s/categories", param0)
}

// List categories by author
func (c *Client) ListCategoriesRelationAuthor(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListCategoriesRelationAuthorRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListCategoriesRelationAuthorRequest create the request corresponding to the listCategories action endpoint of the relationAuthor resource.
func (c *Client) NewListCategoriesRelationAuthorRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListClassesRelationAuthorPath computes a request path to the listClasses action of relationAuthor.
func ListClassesRelationAuthorPath(authorID int) string {
	param0 := strconv.Itoa(authorID)

	return fmt.Sprintf("/authors/%s/classes", param0)
}

// List classes by author
func (c *Client) ListClassesRelationAuthor(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListClassesRelationAuthorRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListClassesRelationAuthorRequest create the request corresponding to the listClasses action endpoint of the relationAuthor resource.
func (c *Client) NewListClassesRelationAuthorRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListRolesRelationAuthorPath computes a request path to the listRoles action of relationAuthor.
func ListRolesRelationAuthorPath(authorID int) string {
	param0 := strconv.Itoa(authorID)

	return fmt.Sprintf("/authors/%s/roles", param0)
}

// List roles by author
func (c *Client) ListRolesRelationAuthor(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListRolesRelationAuthorRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListRolesRelationAuthorRequest create the request corresponding to the listRoles action endpoint of the relationAuthor resource.
func (c *Client) NewListRolesRelationAuthorRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListSeriesRelationAuthorPath computes a request path to the listSeries action of relationAuthor.
func ListSeriesRelationAuthorPath(authorID int) string {
	param0 := strconv.Itoa(authorID)

	return fmt.Sprintf("/authors/%s/series", param0)
}

// List series by author
func (c *Client) ListSeriesRelationAuthor(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListSeriesRelationAuthorRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSeriesRelationAuthorRequest create the request corresponding to the listSeries action endpoint of the relationAuthor resource.
func (c *Client) NewListSeriesRelationAuthorRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListSeriesByCategoryRelationAuthorPath computes a request path to the listSeriesByCategory action of relationAuthor.
func ListSeriesByCategoryRelationAuthorPath(authorID int, categoryID int) string {
	param0 := strconv.Itoa(authorID)
	param1 := strconv.Itoa(categoryID)

	return fmt.Sprintf("/authors/%s/categories/%s/series", param0, param1)
}

// List series by author and category
func (c *Client) ListSeriesByCategoryRelationAuthor(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListSeriesByCategoryRelationAuthorRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSeriesByCategoryRelationAuthorRequest create the request corresponding to the listSeriesByCategory action endpoint of the relationAuthor resource.
func (c *Client) NewListSeriesByCategoryRelationAuthorRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListSeriesByClassRelationAuthorPath computes a request path to the listSeriesByClass action of relationAuthor.
func ListSeriesByClassRelationAuthorPath(authorID int, classID int) string {
	param0 := strconv.Itoa(authorID)
	param1 := strconv.Itoa(classID)

	return fmt.Sprintf("/authors/%s/classes/%s/series", param0, param1)
}

// List series by author and class
func (c *Client) ListSeriesByClassRelationAuthor(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListSeriesByClassRelationAuthorRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSeriesByClassRelationAuthorRequest create the request corresponding to the listSeriesByClass action endpoint of the relationAuthor resource.
func (c *Client) NewListSeriesByClassRelationAuthorRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListSeriesByRoleRelationAuthorPath computes a request path to the listSeriesByRole action of relationAuthor.
func ListSeriesByRoleRelationAuthorPath(authorID int, roleID int) string {
	param0 := strconv.Itoa(authorID)
	param1 := strconv.Itoa(roleID)

	return fmt.Sprintf("/authors/%s/roles/%s/series", param0, param1)
}

// List series by author and role
func (c *Client) ListSeriesByRoleRelationAuthor(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListSeriesByRoleRelationAuthorRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSeriesByRoleRelationAuthorRequest create the request corresponding to the listSeriesByRole action endpoint of the relationAuthor resource.
func (c *Client) NewListSeriesByRoleRelationAuthorRequest(ctx context.Context, path string) (*http.Request, error) {
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
