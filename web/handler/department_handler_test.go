package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDepartment(t *testing.T) {
	// Create a new instance of the WebHandler
	wh := &WebHandler{}

	// Create a new request
	req, err := http.NewRequest("GET", "http://localhost:3030/departments/1", nil)
	require.NoErrorf(t, err, "failed to create a new request: %v", err)

	// Create a new response recorder
	rec := httptest.NewRecorder()

	// Call the GetDepartment handler function
	wh.GetDepartment(rec, req)

	// Check the response status code
	require.Equalf(t, http.StatusOK, rec.Code, "expected status code 200, but got %d", rec.Code)

	// Parse the response body
	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	require.NoErrorf(t, err, "failed to parse the response body: %v", err)

	// Check the response fields
	require.Equalf(t, true, response["success"], "expected success to be true, but got %v", response["success"])
	require.NotNilf(t, response["department"], "expected department to be not nil, but got nil")
}
