package res

import (
	"errors"
	"net/http"
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
		err := Error(w, 400, "test", errors.New("test"))
		if err.Error() != "test" {
			t.Errorf("Expected error to be 'test', got %v", err)
		}
		if w.status != 400 {
			t.Errorf("Expected status code %v, got %v", 400, w.status)
		}
	})
}
