package main

import (
	"embed"
	"leetcode-spaced-repetition/internal"

	goose "github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedMigrations embed.FS

func main() {
	goose.SetBaseFS(embedMigrations)
	config, err := internal.GetConfig()
	if err != nil {
		panic(err)
	}
	db, err := internal.GetDBConnFromConfig(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "."); err != nil {
		panic(err)
	}
}

// RunMigrations uses Goose's provider API to apply migrations.
// func RunMigrations(ctx context.Context, db *sql.DB, migrationsDir string) error {
// 	// Create a Goose provider for Postgres
// 	provider, err := goose.NewProvider("postgres", db, goose.WithNoVersioning(false))
// 	if err != nil {
// 		return fmt.Errorf("failed to create goose provider: %w", err)
// 	}

// 	// Apply all migrations up to the latest
// 	if err := provider.Up(ctx, migrationsDir); err != nil {
// 		return fmt.Errorf("failed to apply migrations: %w", err)
// 	}

// 	return nil
// }
