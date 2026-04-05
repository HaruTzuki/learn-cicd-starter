package auth

import (
	"errors"
	"net/http"
	"testing"
)

func Test_GetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name: "valid header",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-api-key"},
			},
			want:    "my-api-key",
			wantErr: nil,
		},
		{
			name: "missing header",
			headers: http.Header{
				"Authorization": []string{},
			},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed header",
			headers: http.Header{
				"Authorization": []string{"InvalidHeader"},
			},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}
			if (err != nil && tt.wantErr == nil) || (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
