package homepage

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "GET /",
			in:             httptest.NewRequest(http.MethodGet, "/", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   message,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			h := NewHandlers(nil)
			h.Home(test.out, test.in)
			if test.expectedStatus != test.out.Code {
				t.Logf("expected status %v, got %v", test.expectedStatus, test.out.Code)
				t.Fail()
			}

			body := test.out.Body.String()
			if body != test.expectedBody {
				t.Logf("expected body %v, got %v", test.expectedBody, body)
				t.Fail()
			}
		})
	}
}
