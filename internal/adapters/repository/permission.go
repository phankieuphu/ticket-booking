package repository

import (
	"booking-service/internal/adapters/database/models"
	"booking-service/internal/domain/entity"
	"booking-service/internal/domain/ports"
	"context"

	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

// CreatePermission implements [ports.PermissionRepository].
func (p PermissionRepository) CreatePermission(ctx context.Context, permission entity.Permission) error {
	model := p.toModel(permission)
	return p.db.Create(&model).Error
}

// GetAllPermissions implements [ports.PermissionRepository].
func (p PermissionRepository) GetAllPermissions(ctx context.Context) []entity.Permission {
	var models []models.Permission
	p.db.Find(&models)

	permissions := make([]entity.Permission, len(models))
	for i, model := range models {
		permissions[i] = p.toEntity(model)
	}
	return permissions
}

// GetPermissionByID implements [ports.PermissionRepository].
func (p PermissionRepository) GetPermissionByID(ctx context.Context, id int) entity.Permission {
	var model models.Permission
	p.db.First(&model, id)
	return p.toEntity(model)
}

// GetPermissionsByIDs implements [ports.PermissionRepository].
func (p PermissionRepository) GetPermissionsByIDs(ctx context.Context, ids []int) []entity.Permission {
	var models []models.Permission
	p.db.Where("id IN ?", ids).Find(&models)

	permissions := make([]entity.Permission, len(models))
	for i, model := range models {
		permissions[i] = p.toEntity(model)
	}
	return permissions
}

// UpdatePermission implements [ports.PermissionRepository].
func (p PermissionRepository) UpdatePermission(ctx context.Context, permission entity.Permission) error {
	model := p.toModel(permission)
	return p.db.Save(&model).Error
}

// DeletePermission implements [ports.PermissionRepository].
func (p PermissionRepository) DeletePermission(ctx context.Context, id int) error {
	return p.db.Delete(&models.Permission{}, id).Error
}

func (p PermissionRepository) toModel(permission entity.Permission) models.Permission {
	return models.Permission{
		ID:             permission.ID,
		Description:    permission.Description,
		PermissionName: permission.Name,
		Resource:       permission.Resource,
		Action:         permission.Action,
	}
}

func (p PermissionRepository) toEntity(model models.Permission) entity.Permission {
	return entity.Permission{
		ID:          model.ID,
		Name:        model.PermissionName,
		Description: model.Description,
		Resource:    model.Resource,
		Action:      model.Action,
	}
}

func NewPermissionRepository(db *gorm.DB) ports.PermissionRepository {
	return PermissionRepository{
		db: db,
	}
}
