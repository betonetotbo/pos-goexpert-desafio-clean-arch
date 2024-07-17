package rest

import (
	"encoding/json"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/pb"
	"log"
	"net/http"
)

type (
	OrderHandlers interface {
		ListOrdersHandler(w http.ResponseWriter, r *http.Request)
		CreateOrderHandler(w http.ResponseWriter, r *http.Request)
	}

	orderHandlersImpl struct {
		service pb.OrderServiceServer
	}
)

func NewHandler(service pb.OrderServiceServer) OrderHandlers {
	return &orderHandlersImpl{service: service}
}

func (h *orderHandlersImpl) ListOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := h.service.ListOrders(r.Context(), nil)
	if err != nil {
		log.Fatalf("Error while listing orders: %v", err)
	}
	json.NewEncoder(w).Encode(&orders.Orders)
}

func (h *orderHandlersImpl) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	order := &pb.Order{}
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Fatalf("Failed to decode request body: %v", err)
	}

	order, err = h.service.CreateOrder(r.Context(), order)
	if err != nil {
		log.Fatalf("Failed to create order: %v", err)
	}
	json.NewEncoder(w).Encode(&order)
}
