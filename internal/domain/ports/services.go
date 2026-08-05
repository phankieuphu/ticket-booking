package ports

import (
	"booking-service/internal/domain/entity"
	"context"
)

type AccountService interface {
	Save(context.Context, entity.Account) error
}
