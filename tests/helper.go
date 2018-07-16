package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

func PerformRequest(r http.Handler, method, path string, auth string, body string) *httptest.ResponseRecorder {
	reader := bytes.NewReader([]byte(body))
	req, _ := http.NewRequest(method, path, reader)
	req.Header.Add("Authorization", auth)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
