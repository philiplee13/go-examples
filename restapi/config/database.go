package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Dbconfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

// this should really be a pooled connection - but for this ex is fine
func Connect(ctx context.Context, conf Dbconfig) (*pgx.Conn, error) {
	// local connection
	psqlUrl := "postgres://postgres:postgres@localhost:5433/postgres"
	conn, err := pgx.Connect(ctx, psqlUrl)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	return conn, nil
}
