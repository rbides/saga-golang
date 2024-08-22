package postgresql

import (
	"context"
	"database/sql"
	"log"
	"os"
	"saga-golang/api-order/internal/repository"
	"saga-golang/api-order/pkg/models"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New() *Repository {
	godotenv.Load()
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Panic(err)
	}

	return &Repository{db}
}

// Create order
func (r *Repository) Create(ctx context.Context, order *models.Order) error {
	log.Println("Entering repo")
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO orders (id, item_id, quantity, status) VALUES ($1, $2, $3, $4)",
		order.ID,
		order.ItemID,
		order.Quantity,
		order.Status,
	)
	log.Println("DB err:", err)
	return err
}

// Update order status
func (r *Repository) UpdateStatus(ctx context.Context, id uuid.UUID, status models.OrderStatus) error {

	_, err := r.db.ExecContext(
		ctx,
		"UPDATE orders SET status = $1 WHERE id = $2",
		status,
		id,
		// UpdatedAt?
	)
	log.Println("Repo: ", err, id, status)
	if err != nil {
		return repository.ErrNotFound
	}

	return err
}
