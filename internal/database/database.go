package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/config"
	"github.com/spf13/cobra"
	"log"
)

type (
	dbContextKeyValue string
)

var dbContextKey dbContextKeyValue = "db-context-key"

func newDB(host string, port int, user, password, database string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true&charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	log.Printf("Conectando a base de dados %s:%d/%s (user %s)", host, port, database, user)
	return sql.Open("mysql", dsn)
}

func PrepareDatabase(cmd *cobra.Command, _ []string) error {
	cfg := config.GetConfig(cmd.Context())

	db, err := newDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		return fmt.Errorf("falha ao conectar ao banco de dados: %+v", err)
	}
	ctx := context.WithValue(cmd.Context(), dbContextKey, db)

	cmd.SetContext(ctx)
	return nil
}

func GetDB(ctx context.Context) *sql.DB {
	return ctx.Value(dbContextKey).(*sql.DB)
}
