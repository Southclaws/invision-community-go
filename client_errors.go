package ips

import (
	"fmt"
)

// APIError represents the return payload of a failed API request
type APIError struct {
	Code string `json:"errorCode"`
	Name string `json:"errorMessage"`
}

func (err APIError) Error() string {
	return fmt.Sprintf("(%s) %s", err.Code, err.Name)
}
