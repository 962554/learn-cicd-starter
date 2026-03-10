package auth

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   http.Header
		want    string
		wantErr error
	}{
		{
			name: "Empty auth",
			input: http.Header{
				"Authorization": []string{""},
			},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Missing ApiKey in header",
			input: http.Header{
				"Authorization": []string{"f"},
			},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name: "Found wrong ApiKey in header",
			input: http.Header{
				"Authorization": []string{"ApiKey fred"},
			},
			want:    "freds",
			wantErr: nil,
		},
		{
			name: "Found ApiKey in header",
			input: http.Header{
				"Authorization": []string{"ApiKey fred"},
			},
			want:    "fred",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := GetAPIKey(tt.input)
			if err != nil {
				if tt.wantErr.Error() != err.Error() {
					t.Fatalf("expected error: %q, got: %q", tt.wantErr.Error(), err.Error())
				}
			}

			if got != tt.want {
				t.Fatalf("expected: %q, got: %q", tt.want, got)
			}

			fmt.Println(tt)
		})
	}
}
