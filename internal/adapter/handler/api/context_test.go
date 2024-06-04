package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockHTTPContext struct {
	*httptest.ResponseRecorder
	t            *testing.T
	req          *http.Request
	statusCode   int
	responseBody []byte
}

func (m *MockHTTPContext) SetRequest(r *http.Request) {
	m.req = r
}
func (m *MockHTTPContext) JSON(code int, body interface{}) {
	m.statusCode = code
	jsonData, err := json.Marshal(body)
	// json.NewEncoder(m.Body).Encode(body)
	if err != nil {
		m.t.Fatal(err)
	}
	m.responseBody = jsonData
}

func (m *MockHTTPContext) Request() *http.Request {
	return m.req
}
