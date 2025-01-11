package tests

import (
	"Tour/internal/endpoints"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type MockServices struct{}

func (m *MockServices) Compilator() {

}
func TestTest(t *testing.T) {
	app := fiber.New()
	mock := &MockServices{}
	endpoints := endpoints.New(mock)
	app.Get("/test", endpoints.TestHadler)

	test := struct {
		name             string
		expectedStatus   int
		expectedResponse map[string]interface{}
	}{
		name:           "First test",
		expectedStatus: http.StatusOK,

		expectedResponse: map[string]interface{}{
			"test": "test",
		},
	}

	t.Run(test.name, func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test", bytes.NewBuffer(nil))
		resp, err := app.Test(req)

		assert.NoError(t, err)

		assert.Equal(t, test.expectedStatus, resp.StatusCode)

		var responseBody map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&responseBody)
		assert.Equal(t, test.expectedResponse, responseBody)

	})

}
