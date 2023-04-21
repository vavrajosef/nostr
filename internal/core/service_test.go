package core_test

import (
	"testing"

	"github.com/go-nostr/go-nostr/internal/core"
	"github.com/go-nostr/go-nostr/internetidentifier"
	"github.com/stretchr/testify/assert"
)

func TestService_GetHealth(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			req *core.GetHealthRequest
		}
		expect *core.GetHealthResponse
	}{
		{
			name: "GetHealth",
			args: struct {
				req *core.GetHealthRequest
			}{
				req: &core.GetHealthRequest{},
			},
			expect: &core.GetHealthResponse{
				Status: []byte("OK"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := core.NewService()

			resp, err := service.GetHealth(tt.args.req)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, tt.expect.Status, resp.Status)
			assert.NotZero(t, resp.Timestamp)
		})
	}
}

func TestService_GetInternetIdentifier(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			req *core.GetInternetIdentifierRequest
		}
		expect *core.GetInternetIdentifierResponse
	}{
		{
			name: "GetInternetIdentifier",
			args: struct {
				req *core.GetInternetIdentifierRequest
			}{
				req: &core.GetInternetIdentifierRequest{},
			},
			expect: &core.GetInternetIdentifierResponse{
				&internetidentifier.InternetIdentifier{
					Names: map[string]string{
						"bob": "b0635d6a9851d3aed0cd6c495b282167acf761729078d975fc341b22650b07b9",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := core.NewService()

			resp, err := service.GetInternetIdentifier(tt.args.req)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotNil(t, resp.InternetIdentifier)
			assert.Equal(t, tt.expect.InternetIdentifier.Names["bob"], resp.InternetIdentifier.Names["bob"])
		})
	}
}
