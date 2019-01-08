// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": classes Resource Client
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

// CreateClassesPayload is the classes create action payload.
type CreateClassesPayload struct {
	// Class Name (Thriller/Romance/...)
	ClassName string `form:"class_name" json:"class_name" yaml:"class_name" xml:"class_name"`
}

// CreateClassesPath computes a request path to the create action of classes.
func CreateClassesPath() string {

	return fmt.Sprintf("/classes")
}

// Create new class
func (c *Client) CreateClasses(ctx context.Context, path string, payload *CreateClassesPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateClassesRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateClassesRequest create the request corresponding to the create action endpoint of the classes resource.
func (c *Client) NewCreateClassesRequest(ctx context.Context, path string, payload *CreateClassesPayload, contentType string) (*http.Request, error) {
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

// DeleteClassesPath computes a request path to the delete action of classes.
func DeleteClassesPath(classID int) string {
	param0 := strconv.Itoa(classID)

	return fmt.Sprintf("/classes/%s", param0)
}

// delete class by id
func (c *Client) DeleteClasses(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteClassesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteClassesRequest create the request corresponding to the delete action endpoint of the classes resource.
func (c *Client) NewDeleteClassesRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListClassesPath computes a request path to the list action of classes.
func ListClassesPath() string {

	return fmt.Sprintf("/classes")
}

// List classes
func (c *Client) ListClasses(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListClassesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListClassesRequest create the request corresponding to the list action endpoint of the classes resource.
func (c *Client) NewListClassesRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowClassesPath computes a request path to the show action of classes.
func ShowClassesPath(classID int) string {
	param0 := strconv.Itoa(classID)

	return fmt.Sprintf("/classes/%s", param0)
}

// Get class by id
func (c *Client) ShowClasses(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowClassesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowClassesRequest create the request corresponding to the show action endpoint of the classes resource.
func (c *Client) NewShowClassesRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateClassesPayload is the classes update action payload.
type UpdateClassesPayload struct {
	// Class Name (Thriller/Romance/...)
	ClassName string `form:"class_name" json:"class_name" yaml:"class_name" xml:"class_name"`
}

// UpdateClassesPath computes a request path to the update action of classes.
func UpdateClassesPath(classID int) string {
	param0 := strconv.Itoa(classID)

	return fmt.Sprintf("/classes/%s", param0)
}

// Update class by id
func (c *Client) UpdateClasses(ctx context.Context, path string, payload *UpdateClassesPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateClassesRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateClassesRequest create the request corresponding to the update action endpoint of the classes resource.
func (c *Client) NewUpdateClassesRequest(ctx context.Context, path string, payload *UpdateClassesPayload, contentType string) (*http.Request, error) {
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