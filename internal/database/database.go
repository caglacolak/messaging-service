package database

import (
	"context"
	"fmt"
	"project/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	DB *pgxpool.Pool
)

// InitializeDB sets up the PostgreSQL connection
func InitializeDB(cfg *config.Config) error {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.Postgres.Environment.User,
		cfg.Postgres.Environment.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Environment.DB,
	) // Db connection string
	var err error
	DB, err = pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	return nil
}

// CloseDB closes the PostgreSQL connection
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

// return the db connection pool
func GetDB() *pgxpool.Pool {
	return DB
}
