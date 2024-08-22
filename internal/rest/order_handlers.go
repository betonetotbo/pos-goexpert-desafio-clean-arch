package rest

import (
	"context"
	"encoding/json"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/usecase"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/utils"
	"net/http"
)

type (
	OrderHandlers interface {
		ListOrdersHandler(w http.ResponseWriter, r *http.Request)
		CreateOrderHandler(w http.ResponseWriter, r *http.Request)
	}

	orderHandlersImpl struct {
		create *usecase.CreateOrderUseCase
		list   *usecase.ListOrdersUseCase
	}
)

func NewHandler(ctx context.Context) OrderHandlers {
	return &orderHandlersImpl{
		create: usecase.NewCreateOrderUseCase(ctx),
		list:   usecase.NewListOrdersUseCase(ctx),
	}
}

func (h *orderHandlersImpl) ListOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := h.list.Execute(&usecase.ListOrdersInputDTO{
		Limit:  int32(utils.GetQueryParamInt(r, "limit", 10)),
		Offset: int32(utils.GetQueryParamInt(r, "offset", 0)),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&orders.Orders)
}

func (h *orderHandlersImpl) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var in usecase.CreateOrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := h.create.Execute(&in)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&order)
}
