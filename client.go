package ips

import (
	"github.com/go-resty/resty"
	"github.com/pkg/errors"
)

// Client represents a RESTful client for IPS 4 API
type Client struct {
	http *resty.Client
	Info ClientInfo
}

// ClientInfo represents the response from the /core/hello endpoint which includes basic information
// about the running instance of Invision Community.
type ClientInfo struct {
	CommunityName string `json:"communityName"` // The name of the community
	CommunityURL  string `json:"communityUrl"`  // The community URL
	IPSVersion    string `json:"ipsVersion"`    // The Invision Community version number
}

// NewClient creates a client for connecting to the server and pings the /core/hello endpoint in
// order to check if the key is valid and the endpoint is online
func NewClient(endpoint, key string) (*Client, error) {
	cl := &Client{
		http: resty.New().
			SetHostURL(endpoint).
			SetQueryParam("key", key).
			SetRESTMode(),
		Info: ClientInfo{},
	}

	resp, err := cl.http.R().SetResult(cl.Info).Get("/api/core/hello")
	if err != nil {
		return nil, errors.Wrap(err, "failed to check aliveness for endpoint")
	}

	if resp.StatusCode() != 200 {
		return nil, errors.Errorf("aliveness returned non-200: %s", resp.Status())
	}

	return cl, nil
}
