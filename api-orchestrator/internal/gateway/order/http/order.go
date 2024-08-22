package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"saga-golang/api-order/pkg/models"
	"strconv"

	"github.com/google/uuid"
)

type Gateway struct {
	// TODO registry
}

func New() *Gateway {
	return &Gateway{}
}

func (g *Gateway) CreateOrder(ctx context.Context, order *models.Order) error {
	url := "http://api-order:8081/create-order"
	log.Println("Calling ", url)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", order.ID.String())
	values.Add("item_id", order.ItemID.String())
	values.Add("quantity", strconv.Itoa(int(order.Quantity)))
	req.URL.RawQuery = values.Encode()
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

func (g *Gateway) UpdateOrderStatus(ctx context.Context, id uuid.UUID, status models.OrderStatus) error {
	url := "http://api-order:8081/update-status"
	log.Println("Calling ", url)
	req, err := http.NewRequest(http.MethodPost, url, nil) // Should be PUT here
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", id.String())
	values.Add("status", string(status))
	log.Println("DEBUG ", id, status)
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	log.Println("Err", err)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("something went wrong")
	}
	return nil
}
