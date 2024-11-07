package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header

		wantErr bool
	}{
		{
			name:    "Valid API Key",
			headers: http.Header{"Authorization": []string{"ApiKey my-api-key"}},
			wantErr: false,
		}, {
			name:    "Empty Authorization Header",
			headers: http.Header{},
			wantErr: true,
		},
		{
			name:    "No ApiKey Prefix",
			headers: http.Header{"Authorization": []string{"Bearer my-api-key"}},
			wantErr: true,
		}, {
			name:    "Missing ApiKey",
			headers: http.Header{"Authorization": []string{"ApiKey"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error:%v", err)
			}
		})
	}
}
