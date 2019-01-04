// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": editions Resource Client
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

// CreateEditionsPayload is the editions create action payload.
type CreateEditionsPayload struct {
	// Unique Book ID
	BookID int `form:"book_id" json:"book_id" yaml:"book_id" xml:"book_id"`
	// Unique Collection ID
	CollectionID int `form:"collection_id" json:"collection_id" yaml:"collection_id" xml:"collection_id"`
	// Unique Print ID
	PrintID int `form:"print_id" json:"print_id" yaml:"print_id" xml:"print_id"`
}

// CreateEditionsPath computes a request path to the create action of editions.
func CreateEditionsPath() string {

	return fmt.Sprintf("/editions")
}

// Create new edition
func (c *Client) CreateEditions(ctx context.Context, path string, payload *CreateEditionsPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateEditionsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateEditionsRequest create the request corresponding to the create action endpoint of the editions resource.
func (c *Client) NewCreateEditionsRequest(ctx context.Context, path string, payload *CreateEditionsPayload, contentType string) (*http.Request, error) {
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

// DeleteEditionsPath computes a request path to the delete action of editions.
func DeleteEditionsPath(editionID string) string {
	param0 := editionID

	return fmt.Sprintf("/editions/%s", param0)
}

// delete book edition by id
func (c *Client) DeleteEditions(ctx context.Context, path string, editorID *int) (*http.Response, error) {
	req, err := c.NewDeleteEditionsRequest(ctx, path, editorID)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteEditionsRequest create the request corresponding to the delete action endpoint of the editions resource.
func (c *Client) NewDeleteEditionsRequest(ctx context.Context, path string, editorID *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if editorID != nil {
		tmp101 := strconv.Itoa(*editorID)
		values.Set("editor_id", tmp101)
	}
	u.RawQuery = values.Encode()
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

// ListEditionsPath computes a request path to the list action of editions.
func ListEditionsPath() string {

	return fmt.Sprintf("/editions")
}

// List editions
func (c *Client) ListEditions(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListEditionsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListEditionsRequest create the request corresponding to the list action endpoint of the editions resource.
func (c *Client) NewListEditionsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
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

// ShowEditionsPath computes a request path to the show action of editions.
func ShowEditionsPath(editionID int) string {
	param0 := strconv.Itoa(editionID)

	return fmt.Sprintf("/editions/%s", param0)
}

// Get book edition by id
func (c *Client) ShowEditions(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowEditionsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowEditionsRequest create the request corresponding to the show action endpoint of the editions resource.
func (c *Client) NewShowEditionsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
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
