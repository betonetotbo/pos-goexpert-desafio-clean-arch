package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/pb"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func newDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)

	rows := sqlmock.NewRows([]string{"version"}).AddRow("5.7")
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(rows)

	driver := mysql.New(mysql.Config{Conn: db})
	gdb, err := gorm.Open(driver, &gorm.Config{})
	assert.NoError(t, err)
	return gdb, mock
}

func TestListOrders(t *testing.T) {
	// assemble
	db, mock := newDB(t)
	svc := NewOrderService(db)

	rows := sqlmock.NewRows([]string{"id", "customer", "date", "total"}).AddRow("1", "Tester", time.Now(), 1.0)
	mock.ExpectQuery("SELECT * FROM `orders`").WillReturnRows(rows)

	rows = sqlmock.NewRows([]string{"id", "order_id", "product", "price", "quantity", "total"}).AddRow("1.1", "1", "Product", 1.0, 1, 1.0)
	mock.ExpectQuery("SELECT * FROM `order_items` WHERE `order_items`.`order_id` = ?").WithArgs("1").WillReturnRows(rows)

	// act
	orders, err := svc.ListOrders(nil, nil)

	// verify
	if assert.NoError(t, err) {
		assert.Len(t, orders.Orders, 1)
	}
}

func TestCreateOrder(t *testing.T) {
	// assemble
	db, mock := newDB(t)
	svc := NewOrderService(db)
	order := &pb.Order{
		Customer: "Guy",
		Items: []*pb.OrderItem{
			{
				Product:  "Product",
				Price:    10.0,
				Quantity: 2,
			},
		},
	}

	beforeInsert := time.Now().UTC()
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `orders` (`id`,`customer`,`date`,`total`) VALUES (?,?,?,?)").WithArgs(sqlmock.AnyArg(), "Guy", sqlmock.AnyArg(), 20.0).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("INSERT INTO `order_items` (`id`,`order_id`,`product`,`price`,`quantity`,`total`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `order_id`=VALUES(`order_id`)").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "Product", 10.0, 2, 20.0).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// act
	order, err := svc.CreateOrder(nil, order)

	// verify
	if assert.NoError(t, err) {
		assert.Equal(t, "Guy", order.Customer)
		date, err := time.Parse(time.RFC3339, order.Date)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, beforeInsert, date)
		assert.Equal(t, 20.0, order.Total)
		assert.Len(t, order.Items, 1)

		assert.Equal(t, "Product", order.Items[0].Product)
		assert.Equal(t, 10.0, order.Items[0].Price)
		assert.Equal(t, int32(2), order.Items[0].Quantity)
		assert.Equal(t, 20.0, order.Items[0].Total)
	}
}
