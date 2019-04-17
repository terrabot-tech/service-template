package repository

import (
	"time"

	"github.com/terrabot-tech/service-template/models"
	"github.com/terrabot-tech/service-template/provider"
)

// Repository struct
type Repository struct {
	provider provider.Provider
	timeout  time.Duration
}

// NewRepository return new repository
func NewRepository(pr provider.Provider, db models.SQLDataBase) *Repository {
	return &Repository{
		provider: pr,
		timeout:  time.Second,
	}
}
