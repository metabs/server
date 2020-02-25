package main

import (
	"context"
	"github.com/unprogettosenzanomecheforseinizieremo/server/internal/db"
)

func main() {
	ctx := context.Background()
	db, err := db.New(ctx)
	if err != nil {
		panic(err)
	}

	if _, err := db.Collection("Ping").Doc("0").Set(ctx, map[string]bool{"1": true}); err != nil {
		panic(err)
	}
}
