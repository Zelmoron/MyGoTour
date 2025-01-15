package tests

import (
	"Tour/internal/endpoints"
	"Tour/internal/repository"
	"Tour/internal/requests"
	"Tour/internal/services"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type MockServices struct{}

func (m *MockServices) Compilator() {

}
func (m *MockServices) Registration(user requests.RegistrationRequest) error {
	return nil
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

func TestRegistration(t *testing.T) {
	app := fiber.New()
	db := repository.TestConnect()
	repository := repository.New(db)
	service := services.New(repository)
	endpoints := endpoints.New(service)
	app.Post("/registration", endpoints.Registration)

	tests := []struct {
		name             string
		expectedStatus   int
		body             requests.RegistrationRequest
		expectedResponse map[string]interface{}
	}{
		{
			name:           "Success registration",
			expectedStatus: http.StatusOK,
			body: requests.RegistrationRequest{
				Name:     "John Doe",
				Password: "securepassword123",
			},
			expectedResponse: map[string]interface{}{
				"status": "OK - Registration success",
			},
		},
		{
			name:           "Bad registration",
			expectedStatus: http.StatusBadRequest,
			body: requests.RegistrationRequest{
				Name:     "John Doe",
				Password: "securepassword124",
			},
			expectedResponse: map[string]interface{}{
				"status": "BadRequest - Registration error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest("POST", "/registration", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)

			if resp.Body != nil {
				defer resp.Body.Close()

				var responseBody map[string]interface{}
				err = json.NewDecoder(resp.Body).Decode(&responseBody)
				if err != nil {
					t.Fatalf("Ошибка при декодировании ответа: %v", err)
				}

				logrus.Info("Тело ответа: ", responseBody)

				assert.Equal(t, tt.expectedStatus, resp.StatusCode)
				assert.Equal(t, tt.expectedResponse, responseBody)
			}
		})
	}
}
