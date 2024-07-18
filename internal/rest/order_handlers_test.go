package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ServiceMock struct {
	mock.Mock
	pb.UnimplementedOrderServiceServer
}

func (m *ServiceMock) ListOrders(ctx context.Context, e *emptypb.Empty) (*pb.OrderList, error) {
	args := m.Called()
	return args.Get(0).(*pb.OrderList), args.Error(1)
}

func (m *ServiceMock) CreateOrder(ctx context.Context, order *pb.Order) (*pb.Order, error) {
	args := m.Called()
	return args.Get(0).(*pb.Order), args.Error(1)
}

func TestListOrders(t *testing.T) {
	// Assemble
	m := &ServiceMock{}
	handler := NewHandler(m)
	m.On("ListOrders").Return(&pb.OrderList{Orders: []*pb.Order{{
		Id:       "id",
		Customer: "customer",
		Date:     "2024-07-18T01:14:37Z",
		Total:    1.0,
		Items: []*pb.OrderItem{{
			Id:       "id",
			Product:  "product",
			Price:    1.0,
			Quantity: 1,
			Total:    1.0,
		}},
	}}}, nil)
	w := httptest.NewRecorder()

	// Act
	handler.ListOrdersHandler(w, httptest.NewRequest(http.MethodGet, "/orders", nil))

	resp := w.Result()
	data, _ := io.ReadAll(resp.Body)
	assert.Equal(t, `[{"id":"id","customer":"customer","date":"2024-07-18T01:14:37Z","total":1,"items":[{"id":"id","product":"product","price":1,"quantity":1,"total":1}]}]
`, string(data))
}

func TestCreateOrder(t *testing.T) {
	// Assemble
	m := &ServiceMock{}
	handler := NewHandler(m)
	order := &pb.Order{
		Id:       "id",
		Customer: "customer",
		Date:     "2024-07-18T01:14:37Z",
		Total:    1.0,
		Items: []*pb.OrderItem{{
			Id:       "id",
			Product:  "product",
			Price:    1.0,
			Quantity: 1,
			Total:    1.0,
		}},
	}
	m.On("CreateOrder").Return(order, nil)
	w := httptest.NewRecorder()
	data, _ := json.Marshal(order)
	body := bytes.NewReader(data)

	// Act
	handler.CreateOrderHandler(w, httptest.NewRequest(http.MethodPost, "/orders", body))

	resp := w.Result()
	data, _ = io.ReadAll(resp.Body)
	assert.Equal(t, `{"id":"id","customer":"customer","date":"2024-07-18T01:14:37Z","total":1,"items":[{"id":"id","product":"product","price":1,"quantity":1,"total":1}]}
`, string(data))
}
