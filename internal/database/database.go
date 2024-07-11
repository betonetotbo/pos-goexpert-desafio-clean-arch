package database

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type (
	DBContextKeyValue string

	Order struct {
		ID       string `gorm:"primaryKey"`
		Customer string
		Date     time.Time
		Products []Product `gorm:"many2many:order_products;"`
		Total    float64
	}
	Product struct {
		ID          string `gorm:"primaryKey"`
		Description string
		Price       float64
	}
)

var DBContextKey DBContextKeyValue = "db-context-key"

func NewDB(host string, port int, user, password, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func GetDB(ctx context.Context) *gorm.DB {
	return ctx.Value(DBContextKey).(*gorm.DB)
}

func RunMigrations(ctx context.Context) error {
	log.Println("Migrando o esquema da base de dados...")
	db := GetDB(ctx)
	return db.AutoMigrate(&Order{}, &Product{})
}
