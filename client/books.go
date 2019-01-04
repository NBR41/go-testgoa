// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": books Resource Client
//
// Command:
// $ goagen
// --design=github.com/NBR41/go-testgoa/design
// --out=$(GOPATH)/src/github.com/NBR41/go-testgoa
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CreateBooksPayload is the books create action payload.
type CreateBooksPayload struct {
	Isbn string `form:"isbn" json:"isbn" yaml:"isbn" xml:"isbn"`
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Unique Series ID
	SeriesID int `form:"series_id" json:"series_id" yaml:"series_id" xml:"series_id"`
}

// CreateBooksPath computes a request path to the create action of books.
func CreateBooksPath() string {

	return fmt.Sprintf("/books")
}

// Create new book
func (c *Client) CreateBooks(ctx context.Context, path string, payload *CreateBooksPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateBooksRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateBooksRequest create the request corresponding to the create action endpoint of the books resource.
func (c *Client) NewCreateBooksRequest(ctx context.Context, path string, payload *CreateBooksPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSecSigner != nil {
		if err := c.JWTSecSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// DeleteBooksPath computes a request path to the delete action of books.
func DeleteBooksPath(bookID int) string {
	param0 := strconv.Itoa(bookID)

	return fmt.Sprintf("/books/%s", param0)
}

// delete book by id
func (c *Client) DeleteBooks(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteBooksRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteBooksRequest create the request corresponding to the delete action endpoint of the books resource.
func (c *Client) NewDeleteBooksRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSecSigner != nil {
		if err := c.JWTSecSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ListBooksPath computes a request path to the list action of books.
func ListBooksPath() string {

	return fmt.Sprintf("/books")
}

// List books
func (c *Client) ListBooks(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBooksRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBooksRequest create the request corresponding to the list action endpoint of the books resource.
func (c *Client) NewListBooksRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowBooksPath computes a request path to the show action of books.
func ShowBooksPath(bookID int) string {
	param0 := strconv.Itoa(bookID)

	return fmt.Sprintf("/books/%s", param0)
}

// Get book by id
func (c *Client) ShowBooks(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowBooksRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowBooksRequest create the request corresponding to the show action endpoint of the books resource.
func (c *Client) NewShowBooksRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateBooksPayload is the books update action payload.
type UpdateBooksPayload struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// Unique Series ID
	SeriesID *int `form:"series_id,omitempty" json:"series_id,omitempty" yaml:"series_id,omitempty" xml:"series_id,omitempty"`
}

// UpdateBooksPath computes a request path to the update action of books.
func UpdateBooksPath(bookID int) string {
	param0 := strconv.Itoa(bookID)

	return fmt.Sprintf("/books/%s", param0)
}

// Update book by id
func (c *Client) UpdateBooks(ctx context.Context, path string, payload *UpdateBooksPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateBooksRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateBooksRequest create the request corresponding to the update action endpoint of the books resource.
func (c *Client) NewUpdateBooksRequest(ctx context.Context, path string, payload *UpdateBooksPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSecSigner != nil {
		if err := c.JWTSecSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
