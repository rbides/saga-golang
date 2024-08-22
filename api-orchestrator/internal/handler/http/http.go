package http

import (
	"net/http"
	"saga-golang/api-orchestrator/internal/controller/orchestrator"
	"strconv"

	"saga-golang/api-order/pkg/models"

	"github.com/google/uuid"
)

type Handler struct {
	ctrl *orchestrator.Controller
}

func New(ctrl *orchestrator.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) ProcessOrder(w http.ResponseWriter, req *http.Request) {
	item_id, err := uuid.Parse((req.FormValue("item_id")))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	quantity, err := strconv.Atoi(req.FormValue("quantity"))
	if quantity < 1 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	order := &models.Order{
		ID:       uuid.New(),
		ItemID:   item_id,
		Quantity: uint(quantity),
		Status:   models.Processing,
	}
	ctx := req.Context()
	err = h.ctrl.ProcessOrder(ctx, order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
