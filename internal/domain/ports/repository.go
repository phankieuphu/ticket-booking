package ports

import (
	"booking-service/internal/domain/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) error
	Update(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, userID int) error
	GetByID(ctx context.Context, userID int) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	// GetAll(ctx context.Context) ([]entity.User, error)
}

type RoleRepository interface {
	GetRoleByID(ctx context.Context, id int) entity.Role
	GetRolesByIDs(ctx context.Context, ids []int) []entity.Role
	CreateRole(ctx context.Context, role entity.Role)
	GetAllRoles(ctx context.Context) []entity.Role
	UpdateRole(ctx context.Context, role entity.Role)
	GetRolePermissions(ctx context.Context, roleID int) []entity.Permission
}

type PermissionRepository interface {
	GetPermissionByID(ctx context.Context, id int) entity.Permission
	GetPermissionsByIDs(ctx context.Context, ids []int) []entity.Permission
	CreatePermission(ctx context.Context, permission entity.Permission)
	GetAllPermissions(ctx context.Context) []entity.Permission
	UpdatePermission(ctx context.Context, permission entity.Permission)
}
