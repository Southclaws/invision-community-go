package ips

import (
	"github.com/go-resty/resty"
	"github.com/pkg/errors"
)

// Client represents a RESTful client for IPS 4 API
type Client struct {
	http *resty.Client
}

// NewClient creates a client for connecting to the server and pings the /core/hello endpoint in
// order to check if the key is valid and the endpoint is online
func NewClient(endpoint, key string) (*Client, error) {
	cl := &Client{
		http: resty.New().
			SetHostURL(endpoint).
			SetQueryParam("key", key).
			SetRESTMode(),
	}

	resp, err := cl.http.R().Get("/api/core/hello")
	if err != nil {
		return nil, errors.Wrap(err, "failed to check aliveness for endpoint")
	}

	if resp.StatusCode() != 200 {
		return nil, errors.Errorf("aliveness returned non-200: %s", resp.Status())
	}

	return cl, nil
}
