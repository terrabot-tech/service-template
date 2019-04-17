package service

import (
	"github.com/terrabot-tech/service-template/repository"
)

// Service struct
type Service struct {
	Repository *repository.Repository
}

// NewService return new service
func NewService(r *repository.Repository) *Service {
	return &Service{r}
}

//Close service
func (s *Service) Close() {}
