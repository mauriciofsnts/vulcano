package shorten_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mauriciofsnts/vulcano/internal/providers/shorten"
	"github.com/stretchr/testify/assert"
)

func TestShortner(t *testing.T) {
	// Create a test server to mock the HTTP endpoint
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"link": "https://short.link/abc123"}`))
	}))
	defer testServer.Close()

	// Call the Shortner function with the test server URL
	shortLink, err := shorten.Shortner(testServer.URL, nil)
	assert.NoError(t, err)
	assert.Contains(t, shortLink, "https://url.db.cafe/")
}
