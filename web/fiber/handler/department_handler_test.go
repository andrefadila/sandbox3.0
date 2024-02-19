package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"sandbox3.0/persistence"
	"sandbox3.0/repository"
)

func TestGetDepartments(t *testing.T) {
	// Initiate service
	db, _ := persistence.OpenMySqlConn()
	defer db.Close()
	db.MigrateAndSeed()
	rs := repository.NewService(db.MysqlDB)

	// Initiate web handler
	app := fiber.New(fiber.Config{
		AppName: "Sandbox 3.0",
	})
	wh := NewWebHandler(rs, app)
	wh.Init()

	// Login jwt
	reqLogin := httptest.NewRequest("POST", "/login", strings.NewReader("user=admin&password=12345"))
	reqLogin.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	respLogin, respLoginErr := app.Test(reqLogin, 1)
	require.NoErrorf(t, respLoginErr, "failed to send request: %v", respLoginErr)

	// Get the token
	respLoginBody, _ := io.ReadAll(respLogin.Body)
	var loginResult map[string]interface{}
	_ = json.Unmarshal(respLoginBody, &loginResult)
	defer respLogin.Body.Close()
	require.NotNilf(t, loginResult["token"], "expected token to be not nil, but got nil")
	bearerToken := "Bearer " + loginResult["token"].(string)

	// Get departments
	req := httptest.NewRequest("GET", "/departments", nil)
	req.Header.Add("Authorization", bearerToken)
	resp, respErr := app.Test(req, 1)
	require.NoErrorf(t, respErr, "failed to send request: %v", respErr)
	require.Equalf(t, http.StatusOK, resp.StatusCode, "expected status code 200, but got %d", resp.StatusCode)

	// Read the response body
	respBody, bErr := io.ReadAll(resp.Body)
	require.NoErrorf(t, bErr, "failed to read the response body: %v", bErr)

	// Parse the response body
	var response map[string]interface{}
	err := json.Unmarshal(respBody, &response)
	defer resp.Body.Close()
	require.NoErrorf(t, err, "failed to parse the response body: %v", err)

	// Check the response fields
	require.Equalf(t, true, response["success"], "expected success to be true, but got %v", response["success"])
	require.NotNilf(t, response["departments"], "expected departments to be not nil, but got nil")
}
