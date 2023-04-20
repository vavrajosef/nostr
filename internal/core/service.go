package core

import (
	"time"

	"github.com/go-nostr/go-nostr/pkg/internetidentifier"
)

func NewService() *Service {
	return &Service{}
}

type Service struct{}

type GetHealthRequest struct {
}

type GetHealthResponse struct {
	Status    []byte `json:"status,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

func (s *Service) GetHealth(req *GetHealthRequest) (*GetHealthResponse, error) {
	return &GetHealthResponse{
		Status:    []byte("OK"),
		Timestamp: time.Now().Unix(),
	}, nil
}

type GetInternetIdentifierRequest struct {
}

type GetInternetIdentifierResponse struct {
	*internetidentifier.InternetIdentifier
}

func (s *Service) GetInternetIdentifier(req *GetInternetIdentifierRequest) (*GetInternetIdentifierResponse, error) {
	return &GetInternetIdentifierResponse{
		&internetidentifier.InternetIdentifier{
			Names: map[string]string{
				"bob": "b0635d6a9851d3aed0cd6c495b282167acf761729078d975fc341b22650b07b9",
			},
		},
	}, nil
}
