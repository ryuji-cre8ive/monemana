package database

import (
	"embed"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// migrate "github.com/rubenv/sql-migrate"
	"os"
)

var migrationsFS embed.FS

func New() (*gorm.DB, error) {
	uri := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  uri,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// migrations := &migrate.EmbedFileSystemMigrationSource{
	// 	FileSystem: migrationsFS,
	// 	Root:       "migrations",
	// }

	// sqlDB, err := db.DB()
	// if err != nil {
	// 	return nil, err
	// }

	// if _, err := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up); err != nil {
	// 	return nil, err
	// }

	return db, nil
}
