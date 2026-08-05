package ports

import (
	"booking-service/internal/domain/entity"
	"context"
)

type AccountRepository interface {
	Create(ctx context.Context, account entity.Account)
}
