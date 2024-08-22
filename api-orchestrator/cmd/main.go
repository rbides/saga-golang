package main

import (
	"fmt"
	"net/http"
	"saga-golang/api-orchestrator/internal/controller/orchestrator"
	inventorygateway "saga-golang/api-orchestrator/internal/gateway/inventory/http"
	ordergateway "saga-golang/api-orchestrator/internal/gateway/order/http"
	paymentgateway "saga-golang/api-orchestrator/internal/gateway/payment/http"
	httphandler "saga-golang/api-orchestrator/internal/handler/http"
)

func main() {
	fmt.Println("Starting api-orchestrator service")
	ordergt := ordergateway.New()
	inventorygt := inventorygateway.New()
	paymentgt := paymentgateway.New()
	ctrl := orchestrator.New(ordergt, inventorygt, paymentgt)
	handler := httphandler.New(ctrl)
	http.Handle("/make-order", http.HandlerFunc(handler.ProcessOrder))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
