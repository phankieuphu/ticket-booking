package main

import (
	"booking-service/internal/application"
	"context"
)

func main() {
	ctx := context.Background()
	application.AccountApplication(ctx)

}
