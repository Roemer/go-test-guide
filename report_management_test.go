package gotestguide

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportManagement_GetFilters(t *testing.T) {
	// Prepare
	assert := assert.New(t)
	mux, client := setup(t)

	// Register a mock handler for the API endpoint
	mux.HandleFunc("/api/report/filters", func(w http.ResponseWriter, r *http.Request) {
		verifyHttpMethod(assert, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		writeMockResponse(t, w, "")
	})

	// Prepare the expected object
	var expectedObject []FilterInformation
	getObjectFromMockResponse(t, "", &expectedObject)

	// Execute
	effectiveObject, resp, err := client.ReportManagement.GetFilters(1, nil, nil)

	// Verify
	assert.NoError(err, "Should not return an error")
	assert.NotNil(resp, "Response should not be nil")
	assert.Equal(http.StatusOK, resp.StatusCode, "Expected status code to be OK")
	assert.NotNil(effectiveObject, "Returned object should not be nil")

	assert.Equal(len(expectedObject), len(effectiveObject), "Number of filters should match")
	for i := range expectedObject {
		assert.Equal(expectedObject[i].FilterId, effectiveObject[i].FilterId, "Filter ID should match expected value")
		assert.Equal(expectedObject[i].Name, effectiveObject[i].Name, "Filter Name should match expected value")
		assert.Equal(expectedObject[i].Category, effectiveObject[i].Category, "Filter Category should match expected value")
		assert.Equal(expectedObject[i].Description, effectiveObject[i].Description, "Filter Description should match expected value")
	}
}

func TestReportManagement_GetFilter(t *testing.T) {
	// Prepare
	assert := assert.New(t)
	mux, client := setup(t)

	// Register a mock handler for the API endpoint
	mux.HandleFunc("/api/report/filters/1", func(w http.ResponseWriter, r *http.Request) {
		verifyHttpMethod(assert, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		writeMockResponse(t, w, "")
	})

	// Prepare the expected object
	var expectedObject Filter
	getObjectFromMockResponse(t, "", &expectedObject)

	// Execute
	effectiveObject, resp, err := client.ReportManagement.GetFilter(1)

	// Verify
	assert.NoError(err, "Should not return an error")
	assert.NotNil(resp, "Response should not be nil")
	assert.Equal(http.StatusOK, resp.StatusCode, "Expected status code to be OK")
	assert.NotNil(effectiveObject, "Returned object should not be nil")
}

func TestReportManagement_GetTestCaseExecutionsByFilter(t *testing.T) {
	// Prepare
	assert := assert.New(t)
	mux, client := setup(t)

	// Register a mock handler for the API endpoint
	mux.HandleFunc("/api/report/testCaseExecutions/filter", func(w http.ResponseWriter, r *http.Request) {
		verifyHttpMethod(assert, r, http.MethodPost)
		verifyHttpQueryParameter(assert, r, "projectId", "1")
		w.WriteHeader(http.StatusOK)
		writeMockResponse(t, w, "")
	})

	// Prepare the expected object
	var expectedObject []TestCaseExecution
	getObjectFromMockResponse(t, "", &expectedObject)

	// Execute
	effectiveObject, resp, err := client.ReportManagement.GetTestCaseExecutionsByFilter(1, nil, nil, &FilterParameters{})

	// Verify
	assert.NoError(err, "Should not return an error")
	assert.NotNil(resp, "Response should not be nil")
	assert.Equal(http.StatusOK, resp.StatusCode, "Expected status code to be OK")
	assert.NotNil(effectiveObject, "Returned object should not be nil")
}

func TestReportManagement_GetTestCaseExecutionsByProjectFilter(t *testing.T) {
	// Prepare
	assert := assert.New(t)
	mux, client := setup(t)

	// Register a mock handler for the API endpoint
	mux.HandleFunc("/api/report/testCaseExecutions/filter/1", func(w http.ResponseWriter, r *http.Request) {
		verifyHttpMethod(assert, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		writeMockResponse(t, w, "")
	})

	// Prepare the expected object
	var expectedObject []TestCaseExecution
	getObjectFromMockResponse(t, "", &expectedObject)

	// Execute
	effectiveObject, resp, err := client.ReportManagement.GetTestCaseExecutionsByProjectFilter(1, nil, nil)

	// Verify
	assert.NoError(err, "Should not return an error")
	assert.NotNil(resp, "Response should not be nil")
	assert.Equal(http.StatusOK, resp.StatusCode, "Expected status code to be OK")
	assert.NotNil(effectiveObject, "Returned object should not be nil")
}
