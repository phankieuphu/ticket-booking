package ports

import (
	"booking-service/internal/domain/entity"
	"context"
)

type AccountRepository interface {
	Create(ctx context.Context, account entity.User)
}

type RoleRepository interface {
	GetRoleByID(ctx context.Context, id int) entity.Role
}
