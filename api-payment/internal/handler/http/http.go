package http

import (
	"log"
	"net/http"
	"saga-golang/api-payment/internal/controller/payment"
)

type Handler struct {
	ctrl *payment.Controller
}

func New(ctrl *payment.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) ProcessPayment(w http.ResponseWriter, req *http.Request) {
	log.Println("API-Payment - ProcessPayment")
	ctx := req.Context()
	err := h.ctrl.ProcessPayment(ctx)
	if err != nil {
		log.Println("API-Payment - Err2", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
