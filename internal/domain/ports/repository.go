package ports

import (
	"context"
	"user-service/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) error
	Update(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, userID int) error
	GetByID(ctx context.Context, userID int) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	GetByUsername(ctx context.Context, username string) (entity.User, error)
	GetUserRoleIDs(ctx context.Context, userID int) ([]int, error)
	// GetAll(ctx context.Context) ([]entity.User, error)
}

type RoleRepository interface {
	GetRoleByID(ctx context.Context, id int) entity.Role
	GetRolesByIDs(ctx context.Context, ids []int) []entity.Role
	CreateRole(ctx context.Context, role entity.Role) error
	GetAllRoles(ctx context.Context) []entity.Role
	UpdateRole(ctx context.Context, role entity.Role)
	DeleteRole(ctx context.Context, id int) error
	GetRolePermissions(ctx context.Context, roleID int) []entity.Permission
}

type PermissionRepository interface {
	GetPermissionByID(ctx context.Context, id int) entity.Permission
	GetPermissionsByIDs(ctx context.Context, ids []int) []entity.Permission
	CreatePermission(ctx context.Context, permission entity.Permission) error
	GetAllPermissions(ctx context.Context) []entity.Permission
	UpdatePermission(ctx context.Context, permission entity.Permission) error
	DeletePermission(ctx context.Context, id int) error
}
