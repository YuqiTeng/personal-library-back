package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartRouterGetRoot(t *testing.T) {
	router := StartRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestStartRouterGetUsageDataWithoutParam(t *testing.T) {
	router := StartRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/usageData", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
}