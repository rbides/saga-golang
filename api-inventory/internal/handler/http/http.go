package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"saga-golang/api-inventory/internal/controller/inventory"
	"strconv"

	"github.com/google/uuid"
)

type Handler struct {
	ctrl *inventory.Controller
}

func New(ctrl *inventory.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) GetItem(w http.ResponseWriter, req *http.Request) {
	item_id, err := uuid.Parse((req.FormValue("id")))
	log.Println("Handler - Getting item", item_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := req.Context()
	history, err := h.ctrl.Get(ctx, item_id)
	if err != nil && errors.Is(err, inventory.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(history); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}

func (h *Handler) UpdateItemStock(w http.ResponseWriter, req *http.Request) {
	item_id, err := uuid.Parse((req.FormValue("id")))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	quantity, err := strconv.Atoi(req.FormValue("quantity"))
	// Accepts quantity < 0 for compensating
	if quantity == 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := req.Context()
	err = h.ctrl.Update(ctx, item_id, int(quantity))
	if err != nil && errors.Is(err, inventory.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil && errors.Is(err, inventory.ErrOutOfStock) {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
