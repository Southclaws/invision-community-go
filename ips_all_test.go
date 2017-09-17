package ips

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	endpoint string
	key      string
	client   *Client
)

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

	endpoint = raw["endpoint"].(string)
	key = raw["key"].(string)

	os.Exit(m.Run())
}

func TestConnectClient(t *testing.T) {
	var err error
	client, err = NewClient(endpoint, key)
	assert.NoError(t, err)
	assert.NotNil(t, client)
}
