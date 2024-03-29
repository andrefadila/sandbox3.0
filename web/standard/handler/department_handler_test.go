package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"sandbox3.0/persistence"
	"sandbox3.0/persistence/model"
	"sandbox3.0/repository"
)

func setupWebStandardSuite(t *testing.T) (*httptest.Server, func(t *testing.T)) {
	// Initiate service.
	db, _ := persistence.OpenMySqlConn()
	db.MigrateAndSeed()
	rs := repository.NewService(db.MysqlDB)

	// Initiate web handler.
	wh := NewWebHandler(rs)

	// Create a new test server.
	srv := httptest.NewServer(wh.RouteHandler())
	require.NotNilf(t, srv, "srv must be not nil")

	// Return a function to teardown the test.
	return srv, func(t *testing.T) {
		fmt.Println("Teardown suite")
		defer db.Close()
		defer srv.Close()
	}
}

func TestStandardGetDepartments(t *testing.T) {
	// Suite starter
	srv, teardownSuite := setupWebStandardSuite(t)
	defer teardownSuite(t)

	// Create a new request
	resp, err := http.Get(srv.URL + "/departments")
	require.NoErrorf(t, err, "failed to send request: %v", err)

	// Check the response status code
	require.Equalf(t, http.StatusOK, resp.StatusCode, "expected status code 200, but got %d", resp.StatusCode)

	// Read the response body
	respBody, bErr := io.ReadAll(resp.Body)
	require.NoErrorf(t, bErr, "failed to read the response body: %v", bErr)

	// Parse the response body
	var response map[string]interface{}
	err = json.Unmarshal(respBody, &response)
	defer resp.Body.Close()
	require.NoErrorf(t, err, "failed to parse the response body: %v", err)

	// Check the response fields
	require.Equalf(t, true, response["success"], "expected success to be true, but got %v", response["success"])
	require.NotNilf(t, response["departments"], "expected departments to be not nil, but got nil")
}

func TestStandardGetDepartment(t *testing.T) {
	// Suite starter
	srv, teardownSuite := setupWebStandardSuite(t)
	defer teardownSuite(t)

	// Create a new request
	resp, err := http.Get(srv.URL + "/departments/1")
	require.NoErrorf(t, err, "failed to send request: %v", err)

	// Check the response status code
	require.Equalf(t, http.StatusOK, resp.StatusCode, "expected status code 200, but got %d", resp.StatusCode)

	// Read the response body
	respBody, bErr := io.ReadAll(resp.Body)
	require.NoErrorf(t, bErr, "failed to read the response body: %v", bErr)

	// Parse the response body
	var response map[string]interface{}
	err = json.Unmarshal(respBody, &response)
	defer resp.Body.Close()
	require.NoErrorf(t, err, "failed to parse the response body: %v", err)

	// Check the response fields
	require.Equalf(t, true, response["success"], "expected success to be true, but got %v", response["success"])
	require.NotNilf(t, response["department"], "expected department to be not nil, but got nil")
}

func TestStandardCreateDepartment(t *testing.T) {
	// Suite starter
	srv, teardownSuite := setupWebStandardSuite(t)
	defer teardownSuite(t)

	// Send a POST request to create the department
	var jsonStr = []byte(`{"name":"Dept Test"}`)
	resp, err := http.Post(srv.URL+"/departments", "application/json", bytes.NewBuffer(jsonStr))
	require.NoErrorf(t, err, "failed to send request: %v", err)

	// Check the response status code
	require.Equalf(t, http.StatusOK, resp.StatusCode, "expected status code 200, but got %d", resp.StatusCode)

	// Read the response body
	respBody, bErr := io.ReadAll(resp.Body)
	require.NoErrorf(t, bErr, "failed to read the response body: %v", bErr)

	// Parse the response body
	var response map[string]interface{}
	err = json.Unmarshal(respBody, &response)
	defer resp.Body.Close()
	require.NoErrorf(t, err, "failed to parse the response body: %v", err)

	// Check the response fields
	require.Equalf(t, true, response["success"], "expected success to be true, but got %v", response["success"])
	require.NotNilf(t, response["department"], "expected department to be not nil, but got nil")

	jsonData, err := json.Marshal(response["department"])
	require.NoErrorf(t, err, "failed to marshal department: %v", err)

	var dept model.Department
	err = json.Unmarshal(jsonData, &dept)
	require.NoErrorf(t, err, "failed to parse the department: %v", err)
	require.Equalf(t, "Dept Test", dept.Name, "expected department name to be Dept Test, but got %s", dept.Name)
}

func TestStandardUpdateDepartment(t *testing.T) {
	// Suite starter
	srv, teardownSuite := setupWebStandardSuite(t)
	defer teardownSuite(t)

	// Create a new request
	reqBody := []byte(`{"name":"Updated Dept"}`)
	req, err := http.NewRequest("PUT", srv.URL+"/departments/1", bytes.NewBuffer(reqBody))
	require.NoErrorf(t, err, "failed to create request: %v", err)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	require.NoErrorf(t, err, "failed to send request: %v", err)

	// Check the response status code
	require.Equalf(t, http.StatusOK, resp.StatusCode, "expected status code 200, but got %d", resp.StatusCode)

	// Read the response body
	respBody, bErr := io.ReadAll(resp.Body)
	require.NoErrorf(t, bErr, "failed to read the response body: %v", bErr)

	// Parse the response body
	var response map[string]interface{}
	err = json.Unmarshal(respBody, &response)
	defer resp.Body.Close()
	require.NoErrorf(t, err, "failed to parse the response body: %v", err)

	// Check the response fields
	require.Equalf(t, true, response["success"], "expected success to be true, but got %v", response["success"])
	require.NotNilf(t, response["department"], "expected department to be not nil, but got nil")

	jsonData, err := json.Marshal(response["department"])
	require.NoErrorf(t, err, "failed to marshal department: %v", err)

	var dept model.Department
	err = json.Unmarshal(jsonData, &dept)
	require.NoErrorf(t, err, "failed to parse the department: %v", err)
	require.Equalf(t, "Updated Dept", dept.Name, "expected department name to be Updated Dept, but got %s", dept.Name)
}

func TestStandardDeleteDepartment(t *testing.T) {
	// Initiate service
	db, _ := persistence.OpenMySqlConn()
	defer db.Close()
	db.MigrateAndSeed()
	rs := repository.NewService(db.MysqlDB)

	// Initiate web handler
	wh := NewWebHandler(rs)

	// Create a new test server
	srv := httptest.NewServer(wh.RouteHandler())
	defer srv.Close()

	// Send a POST request to create the department
	var jsonStr = []byte(`{"name":"Dept Test"}`)
	respCreate, err := http.Post(srv.URL+"/departments", "application/json", bytes.NewBuffer(jsonStr))
	require.NoErrorf(t, err, "failed to send request: %v", err)

	// Parse the response body create
	respBodyCreate, bErr := io.ReadAll(respCreate.Body)
	require.NoErrorf(t, bErr, "failed to read the response body: %v", bErr)
	var responseCreate map[string]interface{}
	_ = json.Unmarshal(respBodyCreate, &responseCreate)
	defer respCreate.Body.Close()
	jsonData, err := json.Marshal(responseCreate["department"])
	require.NoErrorf(t, err, "failed to marshal department: %v", err)
	var dept model.Department
	err = json.Unmarshal(jsonData, &dept)
	require.NoErrorf(t, err, "failed to parse the department: %v", err)

	// Create a new delete request
	req, err := http.NewRequest("DELETE", fmt.Sprintf(srv.URL+"/departments/%d", dept.ID), nil)
	require.NoErrorf(t, err, "failed to create request: %v", err)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	require.NoErrorf(t, err, "failed to send request: %v", err)

	// Check the response status code
	require.Equalf(t, http.StatusOK, resp.StatusCode, "expected status code 200, but got %d", resp.StatusCode)

	// Read the response body
	respBody, bErr := io.ReadAll(resp.Body)
	require.NoErrorf(t, bErr, "failed to read the response body: %v", bErr)

	// Parse the response body
	var response map[string]interface{}
	err = json.Unmarshal(respBody, &response)
	defer resp.Body.Close()
	require.NoErrorf(t, err, "failed to parse the response body: %v", err)

	// Check the response fields
	require.Equalf(t, true, response["success"], "expected success to be true, but got %v", response["success"])
}
