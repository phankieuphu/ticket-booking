package main

import (
	"context"
	"user-service/internal/application"
)

func main() {
	ctx := context.Background()
	application.AccountApplication(ctx)

}
