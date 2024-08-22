package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Gateway struct {
	// TODO registry
}

func New() *Gateway {
	return &Gateway{}
}

func (g *Gateway) ProcessPayment(ctx context.Context) error {
	url := "http://api-payment:8083/process-payment"
	log.Println("Calling ", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("Err ", err)
		return err
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("something went wrong")
	}
	return nil
}
