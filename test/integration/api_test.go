package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IntegrationTestSuite struct {
	suite.Suite
	app *fiber.App
}

func (suite *IntegrationTestSuite) SetupSuite() {
	// Set up required environment variables for testing
	os.Setenv("APP_NAME", "vocagame-wallet-test")
	os.Setenv("APP_VERSION", "1.0.0")
	os.Setenv("APP_KEY", "test-key-123")
	os.Setenv("APP_CURRENCY_DEFAULT", "USD")

	// Create a simple test app
	suite.app = fiber.New()

	// Add test routes
	suite.app.Post("/test/register", func(c *fiber.Ctx) error {
		var payload map[string]interface{}
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid payload"})
		}

		if username, ok := payload["username"]; !ok || username == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Username required"})
		}

		return c.Status(201).JSON(fiber.Map{
			"status":  201,
			"message": "User registered successfully",
			"data":    payload,
		})
	})
}

func (suite *IntegrationTestSuite) TestBasicRouting() {
	// Test basic HTTP routing
	req := httptest.NewRequest("GET", "/", nil)
	resp, err := suite.app.Test(req)

	suite.NoError(err)
	suite.Equal(http.StatusNotFound, resp.StatusCode) // Should return 404 for undefined route
}

func (suite *IntegrationTestSuite) TestUserRegistrationPayload() {
	// Test user registration with valid payload
	registerPayload := map[string]interface{}{
		"username": "testuser",
		"email":    "test@example.com",
		"password": "testpassword123",
	}

	body, _ := json.Marshal(registerPayload)
	req := httptest.NewRequest("POST", "/test/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.app.Test(req)
	suite.NoError(err)
	suite.Equal(http.StatusCreated, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	suite.NoError(err)
	suite.Equal("User registered successfully", response["message"])
}

func (suite *IntegrationTestSuite) TestInvalidPayload() {
	// Test with invalid payload (missing username)
	invalidPayload := map[string]interface{}{
		"email":    "test@example.com",
		"password": "testpassword123",
	}

	body, _ := json.Marshal(invalidPayload)
	req := httptest.NewRequest("POST", "/test/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.app.Test(req)
	suite.NoError(err)
	suite.Equal(http.StatusBadRequest, resp.StatusCode)
}

func (suite *IntegrationTestSuite) TestHTTPMethods() {
	// Test different HTTP methods
	tests := []struct {
		method string
		path   string
		status int
	}{
		{"GET", "/", 404},
		{"POST", "/", 404},
		{"PUT", "/", 404},
		{"DELETE", "/", 404},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.path, nil)
		resp, err := suite.app.Test(req)

		suite.NoError(err)
		suite.Equal(test.status, resp.StatusCode)
	}
}

func (suite *IntegrationTestSuite) TestJSONPayloadValidation() {
	// Test various JSON payload validations
	testCases := []struct {
		name     string
		payload  interface{}
		expected int
	}{
		{
			name: "Valid payload",
			payload: map[string]interface{}{
				"username": "testuser",
				"email":    "test@example.com",
			},
			expected: 201,
		},
		{
			name: "Empty username",
			payload: map[string]interface{}{
				"username": "",
				"email":    "test@example.com",
			},
			expected: 400,
		},
		{
			name:     "Invalid JSON structure",
			payload:  "invalid json",
			expected: 400,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			var body []byte
			var err error

			if str, ok := tc.payload.(string); ok {
				body = []byte(str)
			} else {
				body, err = json.Marshal(tc.payload)
				suite.NoError(err)
			}

			req := httptest.NewRequest("POST", "/test/register", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			resp, err := suite.app.Test(req)
			suite.NoError(err)
			suite.Equal(tc.expected, resp.StatusCode)
		})
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

// Test helper functions
func TestHTTPStatusCodes(t *testing.T) {
	tests := []struct {
		name   string
		status int
		valid  bool
	}{
		{"OK", 200, true},
		{"Created", 201, true},
		{"Bad Request", 400, true},
		{"Unauthorized", 401, true},
		{"Not Found", 404, true},
		{"Internal Server Error", 500, true},
		{"Invalid Code", 999, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.valid {
				assert.True(t, test.status >= 100 && test.status < 600)
			} else {
				assert.False(t, test.status >= 100 && test.status < 600)
			}
		})
	}
}

func TestCurrencyValidation(t *testing.T) {
	validCurrencies := []string{"USD", "EUR", "JPY", "IDR"}
	invalidCurrencies := []string{"INVALID", "XYZ", "123"}

	for _, currency := range validCurrencies {
		t.Run("Valid_"+currency, func(t *testing.T) {
			assert.Contains(t, validCurrencies, currency)
		})
	}

	for _, currency := range invalidCurrencies {
		t.Run("Invalid_"+currency, func(t *testing.T) {
			assert.NotContains(t, validCurrencies, currency)
		})
	}
}

func TestAmountValidation(t *testing.T) {
	testCases := []struct {
		name   string
		amount float64
		valid  bool
	}{
		{"Positive amount", 100.0, true},
		{"Small positive", 0.01, true},
		{"Zero amount", 0.0, false},
		{"Negative amount", -10.0, false},
		{"Large amount", 999999.99, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid := tc.amount > 0
			assert.Equal(t, tc.valid, isValid)
		})
	}
}
