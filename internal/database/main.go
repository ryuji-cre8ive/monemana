package database

import (
	"database/sql"
	"embed"
	_ "github.com/lib/pq"
	// migrate "github.com/rubenv/sql-migrate"
	"os"
)

var migrationsFS embed.FS

func New() (*sql.DB, error) {
	uri := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	// migrations := &migrate.EmbedFileSystemMigrationSource{
	// 	FileSystem: migrationsFS,
	// 	Root:       "migrations",
	// }

	// if _, err := migrate.Exec(db, "postgres", migrations, migrate.Up); err != nil {
	// 	return nil, err
	// }

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
