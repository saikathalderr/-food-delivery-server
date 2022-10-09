package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/health-check", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	mockResponse := `{
		"status": "ok"
	}`

	assert.JSONEq(t, mockResponse, w.Body.String())
}
