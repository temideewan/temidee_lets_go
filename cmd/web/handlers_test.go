package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"temidee_lets_go.temideewan.net/internal/assert"
)

func TestPing(t *testing.T) {
	// initialize a new httptest.ResponseRecorder
	rr := httptest.NewRecorder()

	// initialize a new dummy http.Request.
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// call the ping handler function, passing in the httptest.ResponseRecorder and http.Request.
	ping(rr, r)

	rs := rr.Result()

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}
