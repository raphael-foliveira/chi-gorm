package client

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/raphael-foliveira/chi-gorm/pkg/resp"
)

func TestParseClientFromBody(t *testing.T) {
	tests := []struct {
		name      string
		body      func() *bytes.Buffer
		assertion func(c Client, err error)
	}{
		{
			"should return an error when the body is invalid",
			func() *bytes.Buffer {
				invalidBody := new(bytes.Buffer)
				invalidBody.WriteString(`{"name":test"`)
				return invalidBody
			},
			func(c Client, err error) {
				if err == nil {
					t.Error("Expected an error, got nil")
				}
			},
		},
		{
			"should return an instance of Client when the body is valid",
			func() *bytes.Buffer {
				validBody := new(bytes.Buffer)
				bytes, _ := json.Marshal(resp.M{
					"name":  "test",
					"email": "test@email.com",
				})
				validBody.Write(bytes)
				return validBody
			},
			func(c Client, err error) {
				if err != nil {
					t.Errorf("Expected nil, got %v", err)
				}
				if c == (Client{}) {
					t.Errorf("Expected an instance of Client, got %v", c)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := parseClientFromBody(tt.body())
			tt.assertion(c, err)
		})
	}
}
