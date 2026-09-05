package ports

import (
	"booking-service/internal/application/command"
	"booking-service/internal/domain/entity"
	"context"
)

type UserService interface {
	Save(context.Context, entity.User) error
}

type RoleService interface {
	Save(context.Context, entity.CreateRole) error
	Update(context.Context, entity.UpdateRole) error
	Delete(context.Context, int) error
}

type PermissionService interface {
	Save(context.Context, entity.Permission) error
	Update(context.Context, entity.Permission) error
	Delete(context.Context, int) error
	GetByID(ctx context.Context, id int) (entity.Permission, error)
	GetByIDs(ctx context.Context, ids []int) ([]entity.Permission, error)
}

type TicketService interface {
	Save(context.Context, command.CreateTicketCommand) error
}
