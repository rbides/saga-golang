package order

import (
	"context"
	"errors"
	"log"
	"saga-golang/api-order/internal/repository"
	"saga-golang/api-order/pkg/models"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("not found")

type orderRepository interface {
	Create(ctx context.Context, order *models.Order) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status models.OrderStatus) error
}

type Controller struct {
	repo orderRepository
}

func New(repo orderRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) Create(ctx context.Context, order *models.Order) error {
	log.Println("API-Order - Controller")
	err := c.repo.Create(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) UpdateStatus(ctx context.Context, id uuid.UUID, status models.OrderStatus) error {
	log.Println("API-Order - Updating order controller")
	err := c.repo.UpdateStatus(ctx, id, status)
	log.Println(err)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return ErrNotFound
	}

	return nil
}
