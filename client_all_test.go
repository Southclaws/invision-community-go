package ips

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

var client *Client

func TestMain(m *testing.M) {
	var raw map[string]interface{}

	bytes, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &raw)
	if err != nil {
		panic(err)
	}

	endpoint := raw["endpoint"].(string)
	key := raw["key"].(string)

	client, err = NewClient(endpoint, key)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
