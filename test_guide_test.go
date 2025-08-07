package gotestguide

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Prepares the mux and client for testing.
// Tests should register handlers on mux which provide mock responses for the API method being tested.
func setup(t *testing.T) (*http.ServeMux, *Client) {
	// mux is the HTTP request multiplexer used with the test server.
	mux := http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	// client is the Gitlab client being tested.
	client, err := NewClient(server.URL, "token")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	return mux, client
}

// Verifies that the HTTP method of the request matches the expected method.
func verifyHttpMethod(assert *assert.Assertions, r *http.Request, want string) {
	assert.Equal(r.Method, want, "Request method should match expected value")
}

func verifyHttpQueryParameter(assert *assert.Assertions, r *http.Request, parameter string, value string) {
	assert.Equal(r.URL.Query().Get(parameter), value, "Query parameter '%s' should match expected value", parameter)
}

// Writes the corresponding mock response to the HTTP response writer.
func writeMockResponse(t *testing.T, w http.ResponseWriter, suffix string) {
	mockFilePath := getMockFilePath(t, suffix)
	fileBytes, err := os.ReadFile(mockFilePath)
	if err != nil {
		t.Fatalf("Failed to read mock response file %s: %v", mockFilePath, err)
	}
	w.Write(fileBytes)
}

// Unmarshals the mock response from the file into the provided variable.
func getObjectFromMockResponse(t *testing.T, suffix string, v any) {
	mockFilePath := getMockFilePath(t, suffix)
	fileBytes, err := os.ReadFile(mockFilePath)
	if err != nil {
		t.Fatalf("Failed to read mock response file %s: %v", mockFilePath, err)
	}
	if v != nil {
		if err := json.Unmarshal(fileBytes, v); err != nil {
			t.Fatalf("Failed to unmarshal object from response: %v", err)
		}
	}
}

// Constructs the file path for the mock response based on the test name and optional suffix.
func getMockFilePath(t *testing.T, suffix string) string {
	testName := strings.ToLower(t.Name())
	if suffix != "" {
		suffix = "_" + strings.ToLower(suffix)
	}
	return filepath.Join("test-resources", "mock-responses", testName+suffix+".json")
}
