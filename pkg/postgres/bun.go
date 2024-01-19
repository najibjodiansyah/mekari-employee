package postgres

import (
	"database/sql"

	"github.com/uptrace/bun/driver/pgdriver"
)

func NewPostgresConn(databaseURL string) *sql.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(databaseURL)))

	return sqldb
}
