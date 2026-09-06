package repository

import (
	"booking-service/internal/adapters/database/models"
	"booking-service/internal/domain/entity"
	"booking-service/internal/domain/ports"
	"context"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

// CreateRole implements [ports.RoleRepository].
func (r RoleRepository) CreateRole(ctx context.Context, role entity.Role) error {
	model := r.toModel(role)
	return r.db.Create(&model).Error
}

// GetAllRoles implements [ports.RoleRepository].
func (r RoleRepository) GetAllRoles(ctx context.Context) []entity.Role {
	var models []models.Role
	r.db.Preload("Permissions").Find(&models)

	var entities []entity.Role
	for _, i := range models {
		entities = append(entities, r.toEntity(i))
	}
	return entities
}

// GetRoleByID implements [ports.RoleRepository].
func (r RoleRepository) GetRoleByID(ctx context.Context, id int) entity.Role {
	var model models.Role
	r.db.Preload("Permissions").First(&model, id)
	return r.toEntity(model)
}

// GetRolePermissions implements [ports.RoleRepository].
func (r RoleRepository) GetRolePermissions(ctx context.Context, roleID int) []entity.Permission {
	var model models.Role
	r.db.Preload("Permissions").First(&model, roleID)

	permissions := make([]entity.Permission, len(model.Permissions))
	for i, p := range model.Permissions {
		permissions[i] = entity.Permission{
			ID:          p.ID,
			Name:        p.PermissionName,
			Description: p.Description,
			Resource:    p.Resource,
			Action:      p.Action,
		}
	}
	return permissions
}

// GetRolesByIDs implements [ports.RoleRepository].
func (r RoleRepository) GetRolesByIDs(ctx context.Context, ids []int) []entity.Role {
	var models []models.Role
	r.db.Preload("Permissions").Where("id IN ?", ids).Find(&models)

	entities := make([]entity.Role, len(models))
	for i, m := range models {
		entities[i] = r.toEntity(m)
	}
	return entities
}

// UpdateRole implements [ports.RoleRepository].
func (r RoleRepository) UpdateRole(ctx context.Context, role entity.Role) {
	model := r.toModel(role)
	r.db.Model(&models.Role{ID: role.ID}).Update("role_name", role.RoleName)
	r.db.Model(&models.Role{ID: role.ID}).Association("Permissions").Replace(model.Permissions)
}

// DeleteRole implements [ports.RoleRepository].
func (r RoleRepository) DeleteRole(ctx context.Context, id int) error {
	return r.db.Delete(&models.Role{}, id).Error
}

func (r RoleRepository) toModel(role entity.Role) models.Role {
	permissions := make([]models.Permission, len(role.RolePermission))
	for i, permissionID := range role.RolePermission {
		permissions[i] = models.Permission{ID: permissionID}
	}
	return models.Role{
		ID:          role.ID,
		RoleName:    role.RoleName,
		Permissions: permissions,
	}
}

func (r RoleRepository) toEntity(model models.Role) entity.Role {
	permissionIDs := make([]int, len(model.Permissions))
	for i, p := range model.Permissions {
		permissionIDs[i] = p.ID
	}
	return entity.Role{
		ID:             model.ID,
		RoleName:       model.RoleName,
		RolePermission: permissionIDs,
	}
}

func NewRoleRepository(db *gorm.DB) ports.RoleRepository {
	return RoleRepository{
		db: db,
	}
}
