// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": relationRole Resource Client
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

// ListAuthorsRelationRolePath computes a request path to the listAuthors action of relationRole.
func ListAuthorsRelationRolePath(roleID int) string {
	param0 := strconv.Itoa(roleID)

	return fmt.Sprintf("/roles/%s/authors", param0)
}

// List authors by role
func (c *Client) ListAuthorsRelationRole(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListAuthorsRelationRoleRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListAuthorsRelationRoleRequest create the request corresponding to the listAuthors action endpoint of the relationRole resource.
func (c *Client) NewListAuthorsRelationRoleRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListSeriesByAuthorRelationRolePath computes a request path to the listSeriesByAuthor action of relationRole.
func ListSeriesByAuthorRelationRolePath(roleID int, authorID int) string {
	param0 := strconv.Itoa(roleID)
	param1 := strconv.Itoa(authorID)

	return fmt.Sprintf("/roles/%s/authors/%s/series", param0, param1)
}

// List series by role and author
func (c *Client) ListSeriesByAuthorRelationRole(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListSeriesByAuthorRelationRoleRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSeriesByAuthorRelationRoleRequest create the request corresponding to the listSeriesByAuthor action endpoint of the relationRole resource.
func (c *Client) NewListSeriesByAuthorRelationRoleRequest(ctx context.Context, path string) (*http.Request, error) {
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
