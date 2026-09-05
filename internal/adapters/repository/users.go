package repository

import (
	"booking-service/internal/adapters/database/models"
	"booking-service/internal/domain/entity"
	"booking-service/internal/domain/ports"
	"context"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// Create implements [ports.UserRepository].
func (a UserRepository) Create(ctx context.Context, user entity.User) error {
	panic("unimplemented")
}

// Delete implements [ports.UserRepository].
func (a UserRepository) Delete(ctx context.Context, userID int) error {
	panic("unimplemented")
}

// GetByEmail implements [ports.UserRepository].
func (a UserRepository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	panic("unimplemented")
}

// GetByID implements [ports.UserRepository].
func (a UserRepository) GetByID(ctx context.Context, userID int) (entity.User, error) {
	panic("unimplemented")
}

// Update implements [ports.UserRepository].
func (a UserRepository) Update(ctx context.Context, user entity.User) error {
	panic("unimplemented")
}

// // Create implements ports.UserRepository.
// func (a UserRepository) Create(ctx context.Context, account entity.User) {
// 	models := a.toModels(account)
// 	a.db.Save(models)
// 	panic("unimplemented")
// }

func (a UserRepository) toModels(account entity.User) models.User {
	panic("unimplemented")
}

func (a UserRepository) toDomain(model models.User) entity.User {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return UserRepository{
		db: db,
	}
}
