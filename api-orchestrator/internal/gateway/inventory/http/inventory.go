package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type Gateway struct {
	// TODO registry
}

func New() *Gateway {
	return &Gateway{}
}

func (g *Gateway) UpdateStock(ctx context.Context, item_id uuid.UUID, quantity int) error {
	url := "http://api-inventory:8082/stock-update"
	log.Println("Calling ", url)
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", item_id.String())
	values.Add("quantity", strconv.Itoa(int(quantity)))
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
