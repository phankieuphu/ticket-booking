package repository

import (
	"booking-service/internal/domain/entity"
	"booking-service/internal/domain/ports"
	"context"

	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

// CreatePermission implements [ports.PermissionRepository].
func (p PermissionRepository) CreatePermission(ctx context.Context, permission entity.Permission) {
	panic("unimplemented")
}

// GetAllPermissions implements [ports.PermissionRepository].
func (p PermissionRepository) GetAllPermissions(ctx context.Context) []entity.Permission {
	panic("unimplemented")
}

// GetPermissionByID implements [ports.PermissionRepository].
func (p PermissionRepository) GetPermissionByID(ctx context.Context, id int) entity.Permission {
	panic("unimplemented")
}

// GetPermissionsByIDs implements [ports.PermissionRepository].
func (p PermissionRepository) GetPermissionsByIDs(ctx context.Context, ids []int) []entity.Permission {
	panic("unimplemented")
}

// UpdatePermission implements [ports.PermissionRepository].
func (p PermissionRepository) UpdatePermission(ctx context.Context, permission entity.Permission) {
	panic("unimplemented")
}

func NewPermissionRepository(db *gorm.DB) ports.PermissionRepository {
	return PermissionRepository{
		db: db,
	}
}
