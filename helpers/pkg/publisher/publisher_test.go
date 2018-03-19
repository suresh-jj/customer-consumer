package publisher

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublishProduct(t *testing.T) {
	os.Setenv("GOOGLE_CLOUD_PROJECT", "google-cloud-project-name")
	os.Setenv("PULL_FOR_NETSUITE_PRODUCT_BETA", "netsuite-product-beta-subscriber-name")
	os.Setenv("NETSUITE_PRODUCT_BETA", "netsuite-product-beta-name")

	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/customer-data-publish", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PublishProduct)

	handler.ServeHTTP(rr, req)

	// Check the status code.
	status := rr.Code
	assert.Equal(t, http.StatusOK, status, "Returned wrong status code.")

	// Check the response body.
	expected := ""
	assert.Equal(t, expected, rr.Body.String(), "Returned unexpected body.")
}
