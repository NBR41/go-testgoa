// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": health Resource Client
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
)

// HealthHealthPath computes a request path to the health action of health.
func HealthHealthPath() string {

	return fmt.Sprintf("/_ah/health")
}

// Perform health check.
func (c *Client) HealthHealth(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewHealthHealthRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewHealthHealthRequest create the request corresponding to the health action endpoint of the health resource.
func (c *Client) NewHealthHealthRequest(ctx context.Context, path string) (*http.Request, error) {
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
