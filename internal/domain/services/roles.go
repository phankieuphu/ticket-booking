package services

import (
	"booking-service/config"
	"booking-service/internal/domain/entity"
	"booking-service/internal/domain/ports"
	"context"
)

type RoleService struct {
	config config.Config
}

// Delete implements [ports.RoleService].
func (r *RoleService) Delete(context.Context, int) error {
	panic("unimplemented")
}

// Save implements [ports.RoleService].
func (r *RoleService) Save(context.Context, entity.CreateRole) error {
	panic("unimplemented")
}

// Update implements [ports.RoleService].
func (r *RoleService) Update(context.Context, entity.UpdateRole) error {
	panic("unimplemented")
}

func NewRoleService(config config.Config) ports.RoleService {
	return &RoleService{
		config: config,
	}
}
