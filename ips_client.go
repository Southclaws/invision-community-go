package ips

import (
	"github.com/go-resty/resty"
	"github.com/pkg/errors"
)

type Client struct {
	http *resty.Client
}

func NewClient(endpoint, key string) (*Client, error) {
	cl := &Client{
		http: resty.New().
			SetDebug(true).
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
