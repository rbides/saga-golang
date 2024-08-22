package orchestrator

import (
	"context"
	"errors"
	"log"

	"saga-golang/api-orchestrator/internal/controller/orchestrator/saga"
	"saga-golang/api-order/pkg/models"

	"github.com/google/uuid"
)

type Controller struct {
	orderGateway     orderGateway
	inventoryGateway inventoryGateway
	paymentGateway   paymentGateway
}

func New(
	orderGateway orderGateway,
	inventoryGateway inventoryGateway,
	paymentGateway paymentGateway,
) *Controller {
	return &Controller{orderGateway, inventoryGateway, paymentGateway}
}

type orderGateway interface {
	CreateOrder(ctx context.Context, order *models.Order) error
	UpdateOrderStatus(ctx context.Context, id uuid.UUID, status models.OrderStatus) error
}

type inventoryGateway interface {
	UpdateStock(ctx context.Context, item_id uuid.UUID, quantity int) error
}

type paymentGateway interface {
	ProcessPayment(ctx context.Context) error
}

func (c *Controller) ProcessOrder(ctx context.Context, order *models.Order) error {
	log.Println("Initializing Order Workflow", order.ID)
	wf := c.buildSaga()
	var j int
	for i, step := range wf {
		if err := step.Transaction(ctx, order); err != nil {
			log.Println("Transaction failed.")
			for j = i; j >= 0; j-- {
				if i == j {
					continue
				}
				if err := wf[j].Compensate(ctx, order); err != nil {
					log.Printf("Compensating failed, you should have a retry maybe.")
				}
			}
			log.Println("Ended Workflow with Error")
			return errors.New("WF Error")
		}
	}
	log.Println("Ended Workflow with Success", order.ID)
	return nil
}

func (c *Controller) buildSaga() [3]saga.SagaStep {
	return [3]saga.SagaStep{
		saga.SagaStep{Transaction: c.createOrder, Compensate: c.compensateOrder},
		saga.SagaStep{Transaction: c.updateStock, Compensate: c.compensateStock},
		saga.SagaStep{Transaction: c.processPayment, Compensate: c.compensatePayment},
	}
}

func (c *Controller) createOrder(ctx context.Context, order *models.Order) error {
	log.Println("Creating Order")
	if err := c.orderGateway.CreateOrder(ctx, order); err != nil {
		log.Println("Failed creating order")
		return err
	}
	log.Println("Created")

	return nil
}

func (c *Controller) compensateOrder(ctx context.Context, order *models.Order) error {
	log.Println("Compensating order")
	if err := c.orderGateway.UpdateOrderStatus(ctx, order.ID, models.Failed); err != nil {
		return err
	}
	log.Println("Compensated")
	return nil
}

func (c *Controller) updateStock(ctx context.Context, order *models.Order) error {
	log.Println("Updating Stock")
	if err := c.inventoryGateway.UpdateStock(ctx, order.ItemID, int(order.Quantity)); err != nil {
		log.Println("Failed updating stock")
		return err
	}
	log.Println("Updated")

	return nil
}

func (c *Controller) compensateStock(ctx context.Context, order *models.Order) error {
	log.Println("Compensating Stock")
	if err := c.inventoryGateway.UpdateStock(ctx, order.ItemID, int(order.Quantity)*(-1)); err != nil {
		return err
	}
	log.Println("Compensated")
	return nil
}

func (c *Controller) processPayment(ctx context.Context, order *models.Order) error {
	log.Println("Processing payment")
	if err := c.paymentGateway.ProcessPayment(ctx); err != nil {
		log.Println("Failed processing payment")
		return err
	}
	log.Println("Processed")
	return nil
}

func (c *Controller) compensatePayment(ctx context.Context, order *models.Order) error {
	// Does nothing, will never be called anyway
	log.Println("Compensating payment")
	log.Println("Compensated")
	return nil
}
