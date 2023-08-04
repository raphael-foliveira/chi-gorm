package res

import (
	"net/http"
	"strings"
	"testing"
)

type testResponseWriter struct {
	status int
	body   string
}

func (trw *testResponseWriter) Header() http.Header {
	return map[string][]string{}
}

func (trw *testResponseWriter) WriteHeader(status int) {
	trw.status = status
}

func (trw *testResponseWriter) Write(b []byte) (int, error) {
	trw.body = string(b)
	return 0, nil
}

func TestError(t *testing.T) {
	t.Run("should write a header and return an error", func(t *testing.T) {
		w := &testResponseWriter{}
		err := Error(w, 400, "test")
		if err != nil {
			t.Errorf("Expected error to be nil, got %v", err)
		}
		if w.status != 400 {
			t.Errorf("Expected status code %v, got %v", 400, w.status)
		}
		if strings.TrimSpace(w.body) != "{\"error\":\"test\"}" {
			t.Errorf("Expected body %v, got %v", "{\"error\":\"test\"}", w.body)
		}
	})
}
