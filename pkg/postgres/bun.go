package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/najibjodiansyah/mekari-employee/pkg/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bunotel"
)

func DB() (*bun.DB, *sql.DB) {
	ctx := context.Background()

	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Config.PgCfg.Username,
		config.Config.PgCfg.Password,
		config.Config.PgCfg.Host,
		config.Config.PgCfg.Port,
		config.Config.PgCfg.Database,
	)

	dbConn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dbURI)))

	if err := dbConn.PingContext(ctx); err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	db := bun.NewDB(dbConn, pgdialect.New())
	db.AddQueryHook(bunotel.NewQueryHook(bunotel.WithDBName(config.Config.PgCfg.Database)))

	// queryLog := bundebug.NewQueryHook(bundebug.WithVerbose(true))
	// db.AddQueryHook(queryLog)

	return db, dbConn
}
