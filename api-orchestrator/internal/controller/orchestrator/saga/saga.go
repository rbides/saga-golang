package saga

import (
	"context"
	"log"
	"saga-golang/api-order/pkg/models"
)

type LocalTransaction func(ctx context.Context, order *models.Order) error
type CompensatingTransaction func(ctx context.Context, order *models.Order) error

type SagaStep struct {
	Transaction LocalTransaction
	Compensate  CompensatingTransaction
}

func SagaWorkflow(ctx context.Context, steps map[string]SagaStep) error {
	log.Println("Initializing Saga Workflow")

	return nil
}
