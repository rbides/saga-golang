package payment

import (
	"context"
	"errors"
	"log"
	"time"
)

var ErrTimeout = errors.New("timeout")

const TIMEOUT = 4 // Use < 5 to force a timeout

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) ProcessPayment(ctx context.Context) error {
	log.Println("Processing payment")
	mychan := make(chan string, 2)

	go func() {
		time.Sleep(time.Second * 5)
		mychan <- "Success"
	}()

	select {
	case out := <-mychan:
		log.Println(out)
	case <-time.After(time.Second * TIMEOUT):
		log.Println("Payment timeout")
		return ErrTimeout
	}
	return nil
}
