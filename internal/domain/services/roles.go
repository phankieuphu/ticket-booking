package services

import (
	"booking-service/config"
	"booking-service/internal/domain/entity"
	"booking-service/internal/domain/ports"
	"context"
	"fmt"
)

type RoleService struct {
	config            config.Config
	roleRepository    ports.RoleRepository
	permissionService ports.PermissionService
}

// Delete implements [ports.RoleService].
func (r *RoleService) Delete(ctx context.Context, id int) error {
	return r.roleRepository.DeleteRole(ctx, id)
}

// Save implements [ports.RoleService].
func (r *RoleService) Save(ctx context.Context, req entity.CreateRole) error {
	if err := r.validatePermissions(ctx, req.Permission); err != nil {
		return err
	}

	role := entity.Role{
		RoleName:       req.Name,
		RolePermission: req.Permission,
	}
	return r.roleRepository.CreateRole(ctx, role)
}

// Update implements [ports.RoleService].
func (r *RoleService) Update(ctx context.Context, req entity.UpdateRole) error {
	if err := r.validatePermissions(ctx, req.Permission); err != nil {
		return err
	}

	role := entity.Role{
		ID:             req.ID,
		RoleName:       req.Name,
		RolePermission: req.Permission,
	}
	r.roleRepository.UpdateRole(ctx, role)
	return nil
}

// GetPermissionsByRoleIDs implements [ports.RoleService].
func (r *RoleService) GetPermissionsByRoleIDs(ctx context.Context, roleIDs []int) ([]entity.Permission, error) {
	seen := make(map[int]bool)
	var permissions []entity.Permission
	for _, roleID := range roleIDs {
		for _, permission := range r.roleRepository.GetRolePermissions(ctx, roleID) {
			if seen[permission.ID] {
				continue
			}
			seen[permission.ID] = true
			permissions = append(permissions, permission)
		}
	}
	return permissions, nil
}

func (r *RoleService) validatePermissions(ctx context.Context, permissionIDs []int) error {
	if len(permissionIDs) == 0 {
		return nil
	}

	permissions, err := r.permissionService.GetByIDs(ctx, permissionIDs)
	if err != nil {
		return err
	}
	if len(permissions) != len(permissionIDs) {
		return fmt.Errorf("one or more permissions do not exist")
	}
	return nil
}

func NewRoleService(config config.Config, repository ports.RoleRepository, permissionService ports.PermissionService) ports.RoleService {
	return &RoleService{
		config:            config,
		roleRepository:    repository,
		permissionService: permissionService,
	}
}
