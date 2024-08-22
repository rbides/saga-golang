package postgresql

import (
	"context"
	"database/sql"
	"log"
	"os"
	"saga-golang/api-inventory/internal/repository"
	"saga-golang/api-inventory/pkg/models"
	"time"

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

// Get inventory item by ID
func (r *Repository) Get(ctx context.Context, item_id uuid.UUID) ([]models.Item, error) {
	log.Println("Repo - Getting item ", item_id)
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM inventory WHERE id::text = $1", item_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var res []models.Item
	for rows.Next() {
		var id uuid.UUID
		var name string
		var stock uint
		var price float64
		var created_at time.Time
		var updated_at time.Time
		if err := rows.Scan(
			&id,
			&name,
			&stock,
			&price,
			&created_at,
			&updated_at,
		); err != nil {
			if err == sql.ErrNoRows {
				log.Println("Err Not found")
				return nil, repository.ErrNotFound
			}
			return nil, err
		}

		res = append(res, models.Item{
			ID:        id,
			Name:      name,
			Stock:     stock,
			Price:     price,
			CreatedAt: created_at,
			UpdatedAt: updated_at,
		})
	}
	return res, nil
}

// Update inventory stock
func (r *Repository) Update(ctx context.Context, item_id uuid.UUID, quantity int) error {
	_, err := r.db.ExecContext(
		ctx,
		"UPDATE inventory SET stock = stock - $1 WHERE id = $2",
		quantity,
		item_id,
	)

	return err
}
