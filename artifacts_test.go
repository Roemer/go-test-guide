package gotestguide

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArtifacts_GetStorages(t *testing.T) {
	// Prepare
	assert := assert.New(t)
	mux, client := setup(t)

	// Register a mock handler for the API endpoint
	mux.HandleFunc("/api/artifact/depositories/dep1/storages", func(w http.ResponseWriter, r *http.Request) {
		verifyHttpMethod(assert, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		writeMockResponse(t, w, "")
	})

	// Prepare the expected object
	var expectedObject storages
	getObjectFromMockResponse(t, "", &expectedObject)

	// Execute
	effectiveObject, resp, err := client.Artifacts.GetStorages("dep1")

	// Verify
	assert.NoError(err, "Should not return an error")
	assert.NotNil(resp, "Response should not be nil")
	assert.Equal(http.StatusOK, resp.StatusCode, "Expected status code to be OK")
	assert.NotNil(effectiveObject, "Returned object should not be nil")

	assert.Len(effectiveObject, len(expectedObject.Storages), "Storage length should match")
	for i, storage := range effectiveObject {
		assert.Equal(storage.GetType(), expectedObject.Storages[i].GetType(), "Storage type should match expected value")
	}
}

func TestArtifacts_GetStorage(t *testing.T) {
	// Prepare
	assert := assert.New(t)
	mux, client := setup(t)

	// Register a mock handler for the API endpoint
	mux.HandleFunc("/api/artifact/depositories/dep1/storages/1", func(w http.ResponseWriter, r *http.Request) {
		verifyHttpMethod(assert, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		writeMockResponse(t, w, "")
	})

	// Prepare the expected object
	var expectedObject StorageFile
	getObjectFromMockResponse(t, "", &expectedObject)

	// Execute
	effectiveObject, resp, err := client.Artifacts.GetStorage("dep1", 1)

	// Verify
	assert.NoError(err, "Should not return an error")
	assert.NotNil(resp, "Response should not be nil")
	assert.Equal(http.StatusOK, resp.StatusCode, "Expected status code to be OK")
	assert.NotNil(effectiveObject, "Returned object should not be nil")
	assert.Equal(expectedObject.GetType(), effectiveObject.GetType(), "Storage type should match expected value")
}
