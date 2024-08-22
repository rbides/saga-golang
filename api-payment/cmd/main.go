package main

import (
	"fmt"
	"net/http"
	"saga-golang/api-payment/internal/controller/payment"
	httphandler "saga-golang/api-payment/internal/handler/http"
)

func main() {
	fmt.Println("Starting api-payment service")
	ctrl := payment.New()
	handler := httphandler.New(ctrl)
	http.Handle("/process-payment", http.HandlerFunc(handler.ProcessPayment))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
