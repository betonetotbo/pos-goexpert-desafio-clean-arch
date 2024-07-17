package database

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func NewDB(host string, port int, user, password, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true&charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
