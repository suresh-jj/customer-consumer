// healthcheck_test.go
package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthMessage(t *testing.T) {
	healthMessage := HealthMessage{Message: "I'm alive!"}
	output := healthMessage.Message
	expected := "I'm alive!"
	assert.Equal(t, expected, output, "Return message")
}

func TestHealthCheckFunc(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Printf(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckFunc)

	handler.ServeHTTP(rr, req)

	// Check the status code.
	status := rr.Code
	assert.Equal(t, http.StatusOK, status, "Returned wrong status code: got %v want %v")

	// Check the response body.
	expected := `{"message":"I'm alive!"}`
	assert.Equal(t, expected, rr.Body.String(), "Returned unexpected body: got %v want %v")
}
