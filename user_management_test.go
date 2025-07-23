package gotestguide

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserManagement_Whoami(t *testing.T) {
	// Prepare
	assert := assert.New(t)
	mux, client := setup(t)

	// Register a mock handler for the API endpoint
	mux.HandleFunc("/api/userManagement/whoami", func(w http.ResponseWriter, r *http.Request) {
		verifyHttpMethod(assert, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		writeMockResponse(t, w, "")
	})

	// Prepare the expected object
	var expectedObject User
	getObjectFromMockResponse(t, "", &expectedObject)

	// Execute
	effectiveObject, resp, err := client.UserManagement.Whoami()

	// Verify
	assert.NoError(err, "Should not return an error")
	assert.NotNil(resp, "Response should not be nil")
	assert.Equal(http.StatusOK, resp.StatusCode, "Expected status code to be OK")
	assert.NotNil(effectiveObject, "Returned object should not be nil")

	assert.Equal(expectedObject.ID, effectiveObject.ID, "User ID should match expected value")
	assert.Equal(expectedObject.UserName, effectiveObject.UserName, "User Name should match expected value")
	assert.Equal(expectedObject.AssociatedProjects, effectiveObject.AssociatedProjects, "Associated Projects should match expected value")
}

func TestUserManagement_GetUsers(t *testing.T) {
	// Prepare
	assert := assert.New(t)
	mux, client := setup(t)

	// Register a mock handler for the API endpoint
	mux.HandleFunc("/api/userManagement/users", func(w http.ResponseWriter, r *http.Request) {
		verifyHttpMethod(assert, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		writeMockResponse(t, w, "")
	})

	// Prepare the expected object
	var expectedObjects []*User
	getObjectFromMockResponse(t, "", &expectedObjects)

	// Execute
	effectiveObjects, resp, err := client.UserManagement.GetUsers()

	// Verify
	assert.NoError(err, "Should not return an error")
	assert.NotNil(resp, "Response should not be nil")
	assert.Equal(http.StatusOK, resp.StatusCode, "Expected status code to be OK")
	assert.NotNil(effectiveObjects, "Returned objects should not be nil")

	assert.Equal(len(expectedObjects), len(effectiveObjects), "Number of users should match expected value")
	for i := range expectedObjects {
		assert.Equal(expectedObjects[i].ID, effectiveObjects[i].ID, "User ID should match expected value")
		assert.Equal(expectedObjects[i].UserName, effectiveObjects[i].UserName, "User Name should match expected value")
		assert.Equal(expectedObjects[i].Email, effectiveObjects[i].Email, "Email should match expected value")
		assert.Equal(expectedObjects[i].UserType, effectiveObjects[i].UserType, "User Type should match expected value")
	}
}

func TestUserManagement_GetRoles(t *testing.T) {
	// Prepare
	assert := assert.New(t)
	mux, client := setup(t)

	// Register a mock handler for the API endpoint
	mux.HandleFunc("/api/userManagement/roles", func(w http.ResponseWriter, r *http.Request) {
		verifyHttpMethod(assert, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		writeMockResponse(t, w, "")
	})

	// Prepare the expected object
	var expectedObjects []*ProjectRole
	getObjectFromMockResponse(t, "", &expectedObjects)

	// Execute
	effectiveObjects, resp, err := client.UserManagement.GetRoles(1)

	// Verify
	assert.NoError(err, "Should not return an error")
	assert.NotNil(resp, "Response should not be nil")
	assert.Equal(http.StatusOK, resp.StatusCode, "Expected status code to be OK")
	assert.NotNil(effectiveObjects, "Returned objects should not be nil")

	assert.Equal(len(expectedObjects), len(effectiveObjects), "Number of roles should match expected value")
	for i := range expectedObjects {
		assert.Equal(expectedObjects[i].ID, effectiveObjects[i].ID, "Role ID should match expected value")
		assert.Equal(expectedObjects[i].Name, effectiveObjects[i].Name, "Role Name should match expected value")
	}
}
