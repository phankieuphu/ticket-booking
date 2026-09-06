package services

import (
	"context"
	"fmt"
	"user-service/config"
	"user-service/internal/domain/entity"
	"user-service/internal/domain/ports"
)

type PermissionService struct {
	config     config.Config
	repository ports.PermissionRepository
}

// Delete implements [ports.PermissionService].
func (p *PermissionService) Delete(ctx context.Context, id int) error {
	return p.repository.DeletePermission(ctx, id)
}

// GetByID implements [ports.PermissionService].
func (p *PermissionService) GetByID(ctx context.Context, id int) (entity.Permission, error) {
	permission := p.repository.GetPermissionByID(ctx, id)
	if permission.ID == 0 {
		return entity.Permission{}, fmt.Errorf("permission with id %d not found", id)
	}
	return permission, nil
}

// GetByIDs implements [ports.PermissionService].
func (p *PermissionService) GetByIDs(ctx context.Context, ids []int) ([]entity.Permission, error) {
	return p.repository.GetPermissionsByIDs(ctx, ids), nil
}

// Save implements [ports.PermissionService].
func (p *PermissionService) Save(ctx context.Context, permission entity.Permission) error {
	return p.repository.CreatePermission(ctx, permission)
}

// Update implements [ports.PermissionService].
func (p *PermissionService) Update(ctx context.Context, permission entity.Permission) error {
	return p.repository.UpdatePermission(ctx, permission)
}

func NewPermissionService(cfg config.Config, repository ports.PermissionRepository) ports.PermissionService {
	return &PermissionService{
		config:     cfg,
		repository: repository,
	}
}
