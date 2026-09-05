package ports

import (
	"booking-service/internal/application/command"
	"booking-service/internal/domain/entity"
	"context"
)

type AccountService interface {
	Save(context.Context, entity.User) error
}

type RoleService interface {
	Save(context.Context, entity.CreateRole) error
	Update(context.Context, entity.UpdateRole) error
	Delete(context.Context, int) error
}

type TicketService interface {
	Save(context.Context, command.CreateTicketCommand) error
}
