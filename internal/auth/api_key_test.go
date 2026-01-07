package auth

import (
    "net/http"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    tests := []struct {
        name          string
        headers       http.Header
        expectedKey   string
        expectedError error
    }{
        {
            name: "Valid API Key",
            headers: http.Header{
                "Authorization": []string{"ApiKey valid_api_key_123"},
            },
            expectedKey:   "valid_api_key_123",
            expectedError: nil,
        },
        // you can add more cases here
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            apiKey, err := GetAPIKey(tt.headers)

            if apiKey != tt.expectedKey {
                t.Errorf("expected key %q, got %q", tt.expectedKey, apiKey)
            }

            if err != tt.expectedError {
                t.Errorf("expected error %v, got %v", tt.expectedError, err)
            }
        })
    }
}