package main

import (
	"context"
	"learn-postgres/common"
	"log/slog"
)

func main() {
	ctx := context.Background()
	// TODO parse arguments
	slog.InfoContext(ctx, "Starting postgres connection")
	helper, err := common.GeneratePostgresConnection(ctx, common.DEFAULT_URL)
	if err != nil {
		panic(err)
	}
	slog.InfoContext(ctx, "Creating tables")
	err = helper.CreateTables(ctx)
	if err != nil {
		panic(err)
	}
	slog.InfoContext(ctx, "Generating demo data")
	err = helper.GenerateDefaultDemoData(ctx)
	if err != nil {
		panic(err)
	}
}
