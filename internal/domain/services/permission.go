package services

import (
	"booking-service/config"
	"booking-service/internal/domain/entity"
	"booking-service/internal/domain/ports"
	"context"
)

type PermissionService struct {
	config     config.Config
	repository ports.PermissionRepository
}

// Delete implements [ports.PermissionService].
func (p *PermissionService) Delete(context.Context, int) error {
	panic("unimplemented")
}

// GetByID implements [ports.PermissionService].
func (p *PermissionService) GetByID(ctx context.Context, id int) (entity.Permission, error) {
	panic("unimplemented")
}

// GetByIDs implements [ports.PermissionService].
func (p *PermissionService) GetByIDs(ctx context.Context, ids []int) ([]entity.Permission, error) {
	panic("unimplemented")
}

// Save implements [ports.PermissionService].
func (p *PermissionService) Save(context.Context, entity.Permission) error {
	panic("unimplemented")
}

// Update implements [ports.PermissionService].
func (p *PermissionService) Update(context.Context, entity.Permission) error {
	panic("unimplemented")
}

func NewPermissionService(cfg config.Config, repository ports.PermissionRepository) ports.PermissionService {
	return &PermissionService{
		config:     cfg,
		repository: repository,
	}
}
