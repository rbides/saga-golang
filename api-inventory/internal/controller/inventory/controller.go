package inventory

import (
	"context"
	"errors"
	"log"
	"saga-golang/api-inventory/internal/repository"
	"saga-golang/api-inventory/pkg/models"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("not found")
var ErrOutOfStock = errors.New("out of stock")

type inventoryRepository interface {
	Get(ctx context.Context, item_id uuid.UUID) ([]models.Item, error)
	Update(ctx context.Context, item_id uuid.UUID, quantity int) error
}

type Controller struct {
	repo inventoryRepository
}

func New(repo inventoryRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) Get(ctx context.Context, item_id uuid.UUID) ([]models.Item, error) {
	log.Println("Controller - Getting item ", item_id)
	res, err := c.repo.Get(ctx, item_id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}

	return res, err
}

func (c *Controller) Update(ctx context.Context, item_id uuid.UUID, quantity int) error {
	items, err := c.repo.Get(ctx, item_id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return ErrNotFound
	}
	if int(items[0].Stock)-quantity < 0 {
		log.Println(ErrOutOfStock)
		return ErrOutOfStock
	}
	err = c.repo.Update(ctx, item_id, quantity)
	if err != nil {
		return err
	}
	return nil
}
