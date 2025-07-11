package common

import (
	"testing"
)

func TestJsonRetorno(t *testing.T) {
	tests := []struct {
		name         string
		mensagem     string
		codigo       int
		expectedCode string
	}{
		{"BadRequest", "erro", 400, "PATRULHAAGRICOLA_CADASTROS_ERROR_BAD_REQUEST"},
		{"Unauthorized", "erro", 401, "PATRULHAAGRICOLA_CADASTROS_ERROR_UNAUTHORIZED"},
		{"Forbidden", "erro", 403, "PATRULHAAGRICOLA_CADASTROS_ERROR_FORBIDDEN"},
		{"Conflict", "erro", 409, "PATRULHAAGRICOLA_CADASTROS_ERROR_CONFLICT"},
		{"InternalServerError", "erro", 500, "PATRULHAAGRICOLA_CADASTROS_ERROR_INTERNAL_SERVER_ERROR"},
		{"Unknown", "erro", 418, "PATRULHAAGRICOLA_CADASTROS_ERROR_UNKNOWN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := JsonRetorno(tt.mensagem, tt.codigo)

			if result["message"] != tt.mensagem {
				t.Errorf("expected message %q, got %q", tt.mensagem, result["message"])
			}
			if result["errorCode"] != tt.expectedCode {
				t.Errorf("expected errorCode %q, got %q", tt.expectedCode, result["errorCode"])
			}
			if result["httpStatusCode"] != tt.codigo {
				t.Errorf("expected httpStatusCode %d, got %v", tt.codigo, result["httpStatusCode"])
			}
			if _, ok := result["debugID"]; !ok {
				t.Error("expected debugID to be present")
			}
		})
	}
}
