package http

import (
	"errors"
	"log"
	"net/http"
	"saga-golang/api-order/internal/controller/order"
	"saga-golang/api-order/pkg/models"
	"strconv"

	"github.com/google/uuid"
)

type Handler struct {
	ctrl *order.Controller
}

func New(ctrl *order.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, req *http.Request) {
	log.Println("API-Order - Creating order")
	id, err := uuid.Parse((req.FormValue("id")))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	item_id, err := uuid.Parse((req.FormValue("item_id")))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	quantity, err := strconv.Atoi(req.FormValue("quantity"))
	// Accepts quantity < 0 for compensating
	if quantity < 1 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err, quantity)
		return
	}
	new_order := &models.Order{
		ID:       id,
		ItemID:   item_id,
		Quantity: uint(quantity),
		Status:   models.Processing,
	}
	ctx := req.Context()
	err = h.ctrl.Create(ctx, new_order)
	if err != nil && errors.Is(err, order.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateOrderStatus(w http.ResponseWriter, req *http.Request) {
	log.Println("API-Order - Updating order")
	id, err := uuid.Parse((req.FormValue("id")))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	status := models.OrderStatus(req.FormValue("status"))
	log.Println("API-Order - Updating order2", status)
	if status != models.Processing && status != models.Completed && status != models.Failed {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := req.Context()
	log.Println("API-Order - Updating order3")
	err = h.ctrl.UpdateStatus(ctx, id, status)
	if err != nil && errors.Is(err, order.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
