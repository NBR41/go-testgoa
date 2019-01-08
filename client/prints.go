// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": prints Resource Client
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

// CreatePrintsPayload is the prints create action payload.
type CreatePrintsPayload struct {
	// Print Name (Deluxe/Ultimate/Pocket)
	PrintName string `form:"print_name" json:"print_name" yaml:"print_name" xml:"print_name"`
}

// CreatePrintsPath computes a request path to the create action of prints.
func CreatePrintsPath() string {

	return fmt.Sprintf("/prints")
}

// Create new print edition
func (c *Client) CreatePrints(ctx context.Context, path string, payload *CreatePrintsPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreatePrintsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreatePrintsRequest create the request corresponding to the create action endpoint of the prints resource.
func (c *Client) NewCreatePrintsRequest(ctx context.Context, path string, payload *CreatePrintsPayload, contentType string) (*http.Request, error) {
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

// DeletePrintsPath computes a request path to the delete action of prints.
func DeletePrintsPath(printID int) string {
	param0 := strconv.Itoa(printID)

	return fmt.Sprintf("/prints/%s", param0)
}

// delete print by id
func (c *Client) DeletePrints(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeletePrintsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeletePrintsRequest create the request corresponding to the delete action endpoint of the prints resource.
func (c *Client) NewDeletePrintsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListPrintsPath computes a request path to the list action of prints.
func ListPrintsPath() string {

	return fmt.Sprintf("/prints")
}

// List prints
func (c *Client) ListPrints(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListPrintsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListPrintsRequest create the request corresponding to the list action endpoint of the prints resource.
func (c *Client) NewListPrintsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowPrintsPath computes a request path to the show action of prints.
func ShowPrintsPath(printID int) string {
	param0 := strconv.Itoa(printID)

	return fmt.Sprintf("/prints/%s", param0)
}

// Get print by id
func (c *Client) ShowPrints(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowPrintsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowPrintsRequest create the request corresponding to the show action endpoint of the prints resource.
func (c *Client) NewShowPrintsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdatePrintsPayload is the prints update action payload.
type UpdatePrintsPayload struct {
	// Print Name (Deluxe/Ultimate/Pocket)
	PrintName string `form:"print_name" json:"print_name" yaml:"print_name" xml:"print_name"`
}

// UpdatePrintsPath computes a request path to the update action of prints.
func UpdatePrintsPath(printID int) string {
	param0 := strconv.Itoa(printID)

	return fmt.Sprintf("/prints/%s", param0)
}

// Update print by id
func (c *Client) UpdatePrints(ctx context.Context, path string, payload *UpdatePrintsPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdatePrintsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdatePrintsRequest create the request corresponding to the update action endpoint of the prints resource.
func (c *Client) NewUpdatePrintsRequest(ctx context.Context, path string, payload *UpdatePrintsPayload, contentType string) (*http.Request, error) {
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