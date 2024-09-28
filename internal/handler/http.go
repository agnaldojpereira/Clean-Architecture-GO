package handler

import (
	"encoding/json"
	"net/http"

	"github.com/seu-usuario/projeto-pedidos/internal/model"
	"github.com/seu-usuario/projeto-pedidos/internal/service"
)

// HTTPHandler lida com as requisições HTTP
type HTTPHandler struct {
	orderService *service.OrderService
}

// NewHTTPHandler cria uma nova instância de HTTPHandler
func NewHTTPHandler(orderService *service.OrderService) *HTTPHandler {
	return &HTTPHandler{orderService: orderService}
}

// ServeHTTP implementa a interface http.Handler
func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listOrders(w, r)
	case http.MethodPost:
		h.createOrder(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// listOrders retorna a lista de todos os pedidos
func (h *HTTPHandler) listOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderService.ListOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// createOrder cria um novo pedido
func (h *HTTPHandler) createOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.orderService.CreateOrder(&order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
