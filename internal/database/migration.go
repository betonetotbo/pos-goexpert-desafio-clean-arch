package database

import (
	"context"
	"gorm.io/gorm"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(ctx context.Context) error {
	log.Println("Migrando o esquema da base de dados...")

	g := ctx.Value(DBContextKey).(*gorm.DB)

	db, err := g.DB()
	if err != nil {
		return err
	}

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)

	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}
	log.Println("Base de dados atualizada")
	return nil
}
