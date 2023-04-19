package core

import "time"

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
