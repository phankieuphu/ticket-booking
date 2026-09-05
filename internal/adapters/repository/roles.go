package repository

import (
	"booking-service/internal/domain/entity"
	"booking-service/internal/domain/ports"
	"context"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

// CreateRole implements [ports.RoleRepository].
func (r RoleRepository) CreateRole(ctx context.Context, role entity.Role) {
	panic("unimplemented")
}

// GetAllRoles implements [ports.RoleRepository].
func (r RoleRepository) GetAllRoles(ctx context.Context) []entity.Role {
	panic("unimplemented")
}

// GetRoleByID implements [ports.RoleRepository].
func (r RoleRepository) GetRoleByID(ctx context.Context, id int) entity.Role {
	panic("unimplemented")
}

// GetRolePermissions implements [ports.RoleRepository].
func (r RoleRepository) GetRolePermissions(ctx context.Context, roleID int) []entity.Permission {
	panic("unimplemented")
}

// GetRolesByIDs implements [ports.RoleRepository].
func (r RoleRepository) GetRolesByIDs(ctx context.Context, ids []int) []entity.Role {
	panic("unimplemented")
}

// UpdateRole implements [ports.RoleRepository].
func (r RoleRepository) UpdateRole(ctx context.Context, role entity.Role) {
	panic("unimplemented")
}

func NewRoleRepository(db *gorm.DB) ports.RoleRepository {
	return RoleRepository{
		db: db,
	}
}
