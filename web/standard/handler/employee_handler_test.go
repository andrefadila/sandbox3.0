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

func TestGetEmployees(t *testing.T) {
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

	// Create a new request
	resp, err := http.Get(srv.URL + "/employees")
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
	require.NotNilf(t, response["employees"], "expected employees to be not nil, but got nil")
}

func TestGetEmployee(t *testing.T) {
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

	// Create a new request
	resp, err := http.Get(srv.URL + "/employees/1")
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
	require.NotNilf(t, response["employee"], "expected employee to be not nil, but got nil")
}

func TestCreateEmployee(t *testing.T) {
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

	// Send a POST request to create the employee
	var jsonStr = []byte(`{"name":"Employee Test"}`)
	resp, err := http.Post(srv.URL+"/employees", "application/json", bytes.NewBuffer(jsonStr))
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
	require.NotNilf(t, response["employee"], "expected employee to be not nil, but got nil")

	jsonData, err := json.Marshal(response["employee"])
	require.NoErrorf(t, err, "failed to marshal employee: %v", err)

	var emp model.Employee
	err = json.Unmarshal(jsonData, &emp)
	require.NoErrorf(t, err, "failed to parse the employee: %v", err)
	require.Equalf(t, "Employee Test", emp.Name, "expected employee name to be Employee Test, but got %s", emp.Name)
}

func TestUpdateEmployee(t *testing.T) {
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

	// Create a new request
	reqBody := []byte(`{"name":"Updated Employee"}`)
	req, err := http.NewRequest("PUT", srv.URL+"/employees/1", bytes.NewBuffer(reqBody))
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
	require.NotNilf(t, response["employee"], "expected employee to be not nil, but got nil")

	jsonData, err := json.Marshal(response["employee"])
	require.NoErrorf(t, err, "failed to marshal employee: %v", err)

	var emp model.Employee
	err = json.Unmarshal(jsonData, &emp)
	require.NoErrorf(t, err, "failed to parse the employee: %v", err)
	require.Equalf(t, "Updated Employee", emp.Name, "expected employee name to be Updated Employee, but got %s", emp.Name)
}

func TestDeleteEmployee(t *testing.T) {
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

	// Send a POST request to create the employee
	var jsonStr = []byte(`{"name":"Employee Test"}`)
	respCreate, err := http.Post(srv.URL+"/employees", "application/json", bytes.NewBuffer(jsonStr))
	require.NoErrorf(t, err, "failed to send request: %v", err)

	// Parse the response body create
	respBodyCreate, bErr := io.ReadAll(respCreate.Body)
	require.NoErrorf(t, bErr, "failed to read the response body: %v", bErr)
	var responseCreate map[string]interface{}
	_ = json.Unmarshal(respBodyCreate, &responseCreate)
	defer respCreate.Body.Close()
	jsonData, err := json.Marshal(responseCreate["employee"])
	require.NoErrorf(t, err, "failed to marshal employee: %v", err)
	var emp model.Employee
	err = json.Unmarshal(jsonData, &emp)
	require.NoErrorf(t, err, "failed to parse the employee: %v", err)

	// Create a new delete request
	req, err := http.NewRequest("DELETE", fmt.Sprintf(srv.URL+"/employees/%d", emp.ID), nil)
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
