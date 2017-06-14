package oceanconnect

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenewal(t *testing.T) {
	var response string
	reqCount := 0

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqCount++
		fmt.Fprintln(w, response)
	}))
	defer s.Close()

	c := Client{
		c:     &http.Client{},
		cfg:   Config{
			URL: s.URL,
			AppID: "<appid>",
		},
	}

	// Test with 300 sec expiry a immediate re-login is issued
	response = `{"accessToken":"85fe3222f362e3b6e943e483bd9c6f9b","tokenType":"bearer","refreshToken":null,"expiresIn":300,"scope":"default"}`
	assert.Nil(t, c.Login(), "expected no error for login")
	assert.Equal(t, 1, reqCount, "request-counter should be 1")

	req, err := http.NewRequest(http.MethodGet, s.URL, nil)
	assert.Nil(t, err, "couldn't get http-request")
	_, err = c.doRequest(req)
	assert.Nil(t, err, "requestFailed")
	assert.Equal(t, 3, reqCount, "request-counter should be 3")

	// Test with 305 sec expiry which shouldnt cause re-login
	response = `{"accessToken":"85fe3222f362e3b6e943e483bd9c6f9b","tokenType":"bearer","refreshToken":null,"expiresIn":305,"scope":"default"}`
	assert.Nil(t, c.Login(), "expected no error for login")
	assert.Equal(t, 4, reqCount, "request-counter should be 4")

	req2, err := http.NewRequest(http.MethodGet, s.URL, nil)
	assert.Nil(t, err, "couldn't get second http-request")
	_, err = c.doRequest(req2)
	assert.Nil(t, err, "requestFailed")
	assert.Equal(t, 5, reqCount, "request-counter should be 5")
}
