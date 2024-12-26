package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/imhasandl/grpc-go-project/services/common/genproto/orders"
	"github.com/imhasandl/grpc-go-project/services/orders/types"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		ordersService: orderService,
	}

	return handler
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req orders.CreateOrderRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		return
	}

	order := &orders.Order{
		OrderID: 42,
		CustomerID: req.GetCustomerID(),
		ProductID: req.GetProductID(),
		Quantity: req.GetQuantity(),
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	json.NewEncoder(w).Encode(res)
}

