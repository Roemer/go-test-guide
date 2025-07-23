package gotestguide

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlatform_GetProject(t *testing.T) {
	// Prepare
	assert := assert.New(t)
	mux, client := setup(t)

	// Register a mock handler for the API endpoint
	mux.HandleFunc("/api/platform/projects/1", func(w http.ResponseWriter, r *http.Request) {
		verifyHttpMethod(assert, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		writeMockResponse(t, w, "")
	})

	// Prepare the expected object
	var expectedObject Project
	getObjectFromMockResponse(t, "", &expectedObject)

	// Execute
	effectiveObject, resp, err := client.Platform.GetProject(1)

	// Verify
	assert.NoError(err, "Should not return an error")
	assert.NotNil(resp, "Response should not be nil")
	assert.Equal(http.StatusOK, resp.StatusCode, "Expected status code to be OK")
	assert.NotNil(effectiveObject, "Returned object should not be nil")

	assert.Equal(expectedObject.ID, effectiveObject.ID, "Project ID should match expected value")
	assert.Equal(expectedObject.Name, effectiveObject.Name, "Project Name should match expected value")
}
