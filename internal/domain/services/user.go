package services

import (
	"booking-service/config"
	"booking-service/internal/domain/entity"
	"booking-service/internal/domain/ports"
	"context"
)

type AccountService struct {
	config     config.Config
	repository ports.AccountRepository
}

// Save implements ports.AccountService.
func (e *AccountService) Save(ctx context.Context, entity entity.User) error {
	// Handle business here:

	e.repository.Create(ctx, entity)

	return nil
}

func NewAccountService(cfg config.Config, repository ports.AccountRepository) ports.AccountService {
	return &AccountService{
		repository: repository,
		config:     cfg,
	}
}
