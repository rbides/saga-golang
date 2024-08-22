package main

import (
	"fmt"
	"net/http"
	"saga-golang/api-order/internal/controller/order"
	httphandler "saga-golang/api-order/internal/handler/http"
	"saga-golang/api-order/internal/repository/postgresql"
)

func main() {
	fmt.Println("Starting api-order service")
	repo := postgresql.New()
	ctrl := order.New(repo)
	handler := httphandler.New(ctrl)
	http.Handle("/create-order", http.HandlerFunc(handler.CreateOrder))
	http.Handle("/update-status", http.HandlerFunc(handler.UpdateOrderStatus))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
