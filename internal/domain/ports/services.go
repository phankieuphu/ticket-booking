package ports

import (
	"booking-service/internal/application/command"
	"booking-service/internal/domain/entity"
	"context"
)

type AccountService interface {
	Save(context.Context, entity.Account) error
}

type TicketService interface {
	Save(context.Context, command.CreateTicketCommand) error
}
