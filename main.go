package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/najibjodiansyah/mekari-employee/controller"
	"github.com/najibjodiansyah/mekari-employee/pkg/config"
	"github.com/najibjodiansyah/mekari-employee/pkg/postgres"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bunotel"
)

func main() {
	ctx := context.Background()

	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Config.PgCfg.Username,
		config.Config.PgCfg.Password,
		config.Config.PgCfg.Host,
		config.Config.PgCfg.Port,
		config.Config.PgCfg.Database,
	)

	dbConn := postgres.NewPostgresConn(dbURI)
	defer dbConn.Close()

	if err := dbConn.PingContext(ctx); err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	v := validator.New()

	db := bun.NewDB(dbConn, pgdialect.New())
	db.AddQueryHook(bunotel.NewQueryHook(bunotel.WithDBName(config.Config.PgCfg.Database)))

	// queryLog := bundebug.NewQueryHook(bundebug.WithVerbose(true))
	// db.AddQueryHook(queryLog)
	controller.EmployeeApi(db, v)
}
