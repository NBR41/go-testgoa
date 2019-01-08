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
func (c *Client) ListAuthorsRelationRole(ctx context.Context, path string, authorID *int) (*http.Response, error) {
	req, err := c.NewListAuthorsRelationRoleRequest(ctx, path, authorID)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListAuthorsRelationRoleRequest create the request corresponding to the listAuthors action endpoint of the relationRole resource.
func (c *Client) NewListAuthorsRelationRoleRequest(ctx context.Context, path string, authorID *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if authorID != nil {
		tmp102 := strconv.Itoa(*authorID)
		values.Set("author_id", tmp102)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListSeriesByAuthorsRelationRolePath computes a request path to the listSeriesByAuthors action of relationRole.
func ListSeriesByAuthorsRelationRolePath(roleID int) string {
	param0 := strconv.Itoa(roleID)

	return fmt.Sprintf("/roles/%s/authors", param0)
}

// List series by role and authors
func (c *Client) ListSeriesByAuthorsRelationRole(ctx context.Context, path string, authorID *int) (*http.Response, error) {
	req, err := c.NewListSeriesByAuthorsRelationRoleRequest(ctx, path, authorID)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSeriesByAuthorsRelationRoleRequest create the request corresponding to the listSeriesByAuthors action endpoint of the relationRole resource.
func (c *Client) NewListSeriesByAuthorsRelationRoleRequest(ctx context.Context, path string, authorID *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if authorID != nil {
		tmp103 := strconv.Itoa(*authorID)
		values.Set("author_id", tmp103)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}