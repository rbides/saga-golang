package main

import (
	"fmt"
	"net/http"
	"saga-golang/api-inventory/internal/controller/inventory"
	httphandler "saga-golang/api-inventory/internal/handler/http"
	"saga-golang/api-inventory/internal/repository/postgresql"
)

func main() {
	fmt.Println("Starting api-inventory service")
	repo := postgresql.New()
	ctrl := inventory.New(repo)
	handler := httphandler.New(ctrl)
	http.Handle("/item", http.HandlerFunc(handler.GetItem))
	http.Handle("/stock-update", http.HandlerFunc(handler.UpdateItemStock))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}

// compensate stock
// http methods
