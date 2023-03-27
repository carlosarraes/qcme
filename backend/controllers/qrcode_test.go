package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestGetUser(t *testing.T) {
	r := chi.NewRouter()
	r.Get("/{name}", app.GetUser)

	tests := []struct {
		name         string
		path         string
		expectedCode int
		expectedBody []byte
	}{
		{
			name:         "testuser",
			path:         "/testuser",
			expectedCode: 200,
			expectedBody: []byte(`{"name":"testuser","linkedIn":"https://www.linkedin.com/in/testuser","gitHub":"https://github.com/testuser"}`),
		},
		{
			name:         "",
			path:         "/unknownuser",
			expectedCode: 404,
			expectedBody: []byte(`{"message":"user not found"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", test.path, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			wantCode := rr.Code

			if wantCode != test.expectedCode {
				t.Errorf("expected status code %d, got %d", test.expectedCode, wantCode)
			}

			wantBody := strings.TrimSpace(rr.Body.String())
			if wantBody != string(test.expectedBody) {
				t.Errorf("expected body %s, got %s", test.expectedBody, wantBody)
			}
		})
	}
}

func TestGenerateQRCode(t *testing.T) {
	r := chi.NewRouter()
	r.Post("/qrcodeme", app.GenerateQRCode)

	tests := []struct {
		name         string
		body         string
		expectedCode int
		expectedBody []byte
	}{
		{
			name:         "newtestuser",
			body:         `{"name":"newtestuser","linkedIn":"https://www.linkedin.com/in/testuser","gitHub":"https://github.com/testuser"}`,
			expectedCode: 200,
		},
		{
			name:         "testuser",
			body:         `{"name":"testuser","linkedIn":"https://www.linkedin.com/in/testuser","gitHub":"https://github.com/testuser"}`,
			expectedCode: 409,
		},
		{
			name:         "invalid",
			body:         `{"name":"testuser","linkedIn":"https://www.linkedin.com/in/testuser"}`,
			expectedCode: 400,
		},
		{
			name:         "fail",
			body:         `{"name":"fail","linkedIn":"https://www.linkedin.com/in/failuser","gitHub":"https://github.com/failuser"}`,
			expectedCode: 500,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/qrcodeme", strings.NewReader(test.body))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			wantCode := rr.Code

			if wantCode != test.expectedCode {
				t.Errorf("expected status code %d, got %d", test.expectedCode, wantCode)
			}
		})
	}
}
