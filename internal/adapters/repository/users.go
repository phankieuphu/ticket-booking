package repository

import (
	"booking-service/internal/adapters/database/models"
	"booking-service/internal/domain/entity"
	"booking-service/internal/domain/ports"
	"context"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

// Create implements ports.AccountRepository.
func (a AccountRepository) Create(ctx context.Context, account entity.User) {
	models := a.toModels(account)
	a.db.Save(models)
	panic("unimplemented")
}

func (a AccountRepository) toModels(account entity.User) models.User {
	panic("unimplemented")
}

func (a AccountRepository) toDomain(model models.User) entity.User {
	panic("unimplemented")
}

func NewAccountRepository(db *gorm.DB) ports.AccountRepository {
	return AccountRepository{
		db: db,
	}
}
