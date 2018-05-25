package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"

	serviceMocks "github.com/rodrigodiez/zorro/lib/mocks/service"
)

func TestServerMaskHandler(t *testing.T) {
	t.Parallel()

	zorro := &serviceMocks.Zorro{}
	server := &server{
		zorro: zorro,
	}

	req, _ := http.NewRequest("POST", "/", nil)
	rr := httptest.NewRecorder()

	vars := make(map[string]string)
	vars["key"] = "foo"
	req = mux.SetURLVars(req, vars)

	zorro.On("Mask", "foo").Return("bar").Once()

	server.maskHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "bar", rr.Body.String())
	assert.Equal(t, rr.Header().Get("Content-Type"), "text/plain")
	zorro.AssertExpectations(t)
}

func TestServerunmaskHandler(t *testing.T) {
	t.Parallel()

	tt := []struct {
		value  string
		key    string
		ok     bool
		status int
		body   string
	}{
		{value: "bar", key: "foo", ok: true, status: http.StatusOK, body: "foo"},
		{value: "bar", key: "", ok: false, status: http.StatusNotFound, body: "Not Found"},
	}

	zorro := &serviceMocks.Zorro{}
	server := &server{
		zorro: zorro,
	}

	for _, tc := range tt {
		req, _ := http.NewRequest("POST", "/", nil)
		rr := httptest.NewRecorder()

		vars := make(map[string]string)
		vars["value"] = "bar"
		req = mux.SetURLVars(req, vars)

		zorro.On("Unmask", tc.value).Return(tc.key, tc.ok).Once()

		server.unmaskHandler(rr, req)

		assert.Equal(t, tc.status, rr.Code)
		assert.Equal(t, tc.body, rr.Body.String())
		assert.Equal(t, rr.Header().Get("Content-Type"), "text/plain")
		zorro.AssertExpectations(t)
	}

}
