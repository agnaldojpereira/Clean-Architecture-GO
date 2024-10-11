// pkg/rest/handler.go
package rest

import (
	"encoding/json"
	"net/http"

	"github.com/agnaldojpereira/Clean-Architecture-GO/internal/service"
)

type Handler struct {
	orderService *service.OrderService
}

func NewHandler(orderService *service.OrderService) *Handler {
	return &Handler{orderService: orderService}
}

func (h *Handler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderService.ListOrders(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
