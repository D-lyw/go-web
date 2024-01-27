package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	router := SetAndConfigRouter()

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	router.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
