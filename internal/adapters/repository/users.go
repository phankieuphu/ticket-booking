package repository

import (
	"context"
	"strconv"
	"user-service/internal/adapters/database/models"
	"user-service/internal/domain/entity"
	"user-service/internal/domain/ports"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// Create implements [ports.UserRepository].
func (a UserRepository) Create(ctx context.Context, user entity.User) error {
	model := a.toModels(user)
	return a.db.Create(&model).Error
}

// Delete implements [ports.UserRepository].
func (a UserRepository) Delete(ctx context.Context, userID int) error {
	return a.db.Delete(&models.User{}, userID).Error
}

// GetByEmail implements [ports.UserRepository].
func (a UserRepository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var model models.User
	err := a.db.Joins("Profile").Where(`"Profile".email = ?`, email).First(&model).Error
	if err != nil {
		return entity.User{}, err
	}
	return a.toDomain(model), nil
}

// GetByID implements [ports.UserRepository].
func (a UserRepository) GetByID(ctx context.Context, userID int) (entity.User, error) {
	var model models.User
	err := a.db.First(&model, userID).Error
	if err != nil {
		return entity.User{}, err
	}
	return a.toDomain(model), nil
}

// Update implements [ports.UserRepository].
func (a UserRepository) Update(ctx context.Context, user entity.User) error {
	model := a.toModels(user)
	return a.db.Save(&model).Error
}

// GetByUsername implements [ports.UserRepository].
func (a UserRepository) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	var model models.User
	err := a.db.Where("username = ?", username).First(&model).Error
	if err != nil {
		return entity.User{}, err
	}
	return a.toDomain(model), nil
}

// GetUserRoleIDs implements [ports.UserRepository].
func (a UserRepository) GetUserRoleIDs(ctx context.Context, userID int) ([]int, error) {
	var userRoles []models.UserRole
	if err := a.db.Where("user_id = ?", userID).Find(&userRoles).Error; err != nil {
		return nil, err
	}

	roleIDs := make([]int, len(userRoles))
	for i, ur := range userRoles {
		roleIDs[i] = ur.RoleID
	}
	return roleIDs, nil
}

func (a UserRepository) toModels(user entity.User) models.User {
	id, _ := strconv.Atoi(user.ID)
	return models.User{
		ID:           id,
		Username:     user.Username,
		PasswordHash: user.Password,
	}
}

func (a UserRepository) toDomain(model models.User) entity.User {
	return entity.User{
		ID:       strconv.Itoa(model.ID),
		Username: model.Username,
		Password: model.PasswordHash,
	}
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return UserRepository{
		db: db,
	}
}
