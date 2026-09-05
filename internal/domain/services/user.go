package services

import (
	"booking-service/config"
	"booking-service/internal/domain/entity"
	"booking-service/internal/domain/ports"
	"context"
)

type UserService struct {
	config      config.Config
	repository  ports.UserRepository
	roleService ports.RoleService
}

// Save implements ports.UserService		.
func (e *UserService) Save(ctx context.Context, entity entity.User) error {
	// Handle business here:

	e.repository.Create(ctx, entity)

	return nil
}

func NewUserService(cfg config.Config, repository ports.UserRepository, roleService ports.RoleService) ports.UserService {
	return &UserService{
		repository:  repository,
		config:      cfg,
		roleService: roleService,
	}
}
