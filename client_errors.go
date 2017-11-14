package ips

// APIError represents the return payload of a failed API request
type APIError struct {
	Code string `json:"errorCode"`
	Name string `json:"errorMessage"`
}

// Errors stores a list of possible API errors. IPS errors often use the same codes for different
// error types (for some unknown reason) so here, the code and name are used as a key in order to
// uniquely identify each error.
var Errors = map[string]string{
	"1S290/A:IP_ADDRESS_BANNED":              "The IP address that is sending the request has been banned from the community.",
	"1S290/A:TOO_MANY_REQUESTS_WITH_BAD_KEY": "The IP address that is sending the request has sent multiple requests with an invalid API key and so is prevented from sending any more requests for several minutes.",
	"2S290/6:NO_API_KEY":                     "No API key was sent in the request.",
	"2S290/8:IP_ADDRESS_NOT_ALLOWED":         "The API key was valid, but is configured to only be valid for requests coming from certain IP addresses and IP address the request was sent from is not in the allowed list.",
	"2S290/B:CANNOT_USE_KEY_AS_URL_PARAM":    "The API key was valid, but it is not configured to be used as URL authentication and must be used in AUTHORIZATION headers.",
	"3S290/7:INVALID_API_KEY":                "The API key sent in the request is not valid.",
	"2S290/9:INVALID_LANGUAGE":               "An X-IPS-Language header was sent in the request (which can be used to specify a language ID for the response), but its value was not valid.",
	"3S290/3:INVALID_APP":                    "The endpoint the request was sent to does not exist (the first level contains an invalid character, only alphanumerics are acceptable).",
	"3S290/4:INVALID_CONTROLLER":             "The endpoint the request was sent to does not exist (the second level contains an invalid character, only alphanumerics are acceptable).",
	"2S290/1:INVALID_APP":                    "The endpoint the request was sent to does not exist (the first level does not exist).",
	"1S290/2:APP_DISABLED":                   "The application which controls the endpoint the request was sent to is currently disabled.",
	"2S290/5:INVALID_CONTROLLER":             "The endpoint the request was sent to does not exist (the second level does not exist).",
	"2S291/1:NO_ENDPOINT":                    "The endpoint the request was sent to does not exist (the URL contains too many levels).",
	"2S291/3:NO_PERMISSION":                  "The API key does not have permission to access the requested endpoint.",
	"3S291/2:BAD_METHOD":                     "The endpoint the request was sent to does not exist - the HTTP request method may be incorrect (for example, sending a GET rather than a POST).",
}

// Description returns a human-friendly description of an error
func (e APIError) Description() string {
	return Errors[e.Code+":"+e.Name]
}
