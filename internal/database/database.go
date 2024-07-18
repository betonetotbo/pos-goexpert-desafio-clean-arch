package database

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type (
	DBContextKeyValue string

	Order struct {
		ID       string      `gorm:"size:191;index" json:"id"`
		Customer string      `json:"customer"`
		Date     time.Time   `json:"date"`
		Items    []OrderItem `json:"items"`
		Total    float64     `json:"total"`
	}

	OrderItem struct {
		ID       string  `gorm:"size:191;primaryKey" json:"id"`
		OrderID  string  `gorm:"size:191" json:"-"`
		Order    Order   `json:"-"`
		Product  string  `json:"product"`
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity"`
		Total    float64 `json:"total"`
	}
)

var DBContextKey DBContextKeyValue = "db-context-key"

func newDB(host string, port int, user, password, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true&charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	log.Printf("Conectando a base de dados %s:%d/%s (user %s)", host, port, database, user)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func PrepareDatabase(cmd *cobra.Command, args []string) {
	root := cmd.Parent()
	for root.Parent() != nil {
		root = root.Parent()
	}
	host, _ := root.PersistentFlags().GetString("db-host")
	port, _ := root.PersistentFlags().GetInt("db-port")
	user, _ := root.PersistentFlags().GetString("db-user")
	pass, _ := root.PersistentFlags().GetString("db-password")
	dbName, _ := root.PersistentFlags().GetString("db-name")

	db, err := newDB(host, port, user, pass, dbName)
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %+v", err)
	}
	ctx := context.WithValue(context.Background(), DBContextKey, db)

	cmd.SetContext(ctx)
}

// Hooks

func (o *Order) BeforeCreate(_ *gorm.DB) error {
	o.ID = uuid.New().String()
	return nil
}

func (i *OrderItem) BeforeCreate(_ *gorm.DB) error {
	i.ID = uuid.New().String()
	return nil
}
