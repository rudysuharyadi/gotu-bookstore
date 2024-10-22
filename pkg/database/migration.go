package database

import (
	"embed"
	"github.com/pressly/goose/v3"
	"gorm.io/gorm"
)

func (n Database) EmbedMigrations(databaseDialector gorm.Dialector, embedMigrations embed.FS, migrationsFolder string) error {
	db, err := n.DB.DB()
	if err != nil {
		return err
	}

	goose.SetBaseFS(embedMigrations)

	if err = goose.SetDialect(databaseDialector.Name()); err != nil {
		return err
	}

	if err = goose.Up(db, migrationsFolder); err != nil {
		return err
	}

	return nil
}
