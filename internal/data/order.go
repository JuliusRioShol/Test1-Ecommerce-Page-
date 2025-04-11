package data

import (
	"database/sql"
	"fmt"
	"time"
)

// validator needs to be implemented
/*
	Here is our order Model
*/
// We create our "product" struct
type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userID"` // get this from session/auth
	ProductID int       `json:"productID"`
	Quantity  int       `json:"quantity"`
	Size      string    `json:"size"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type OrderModel struct {
	DB *sql.DB
}

func (m *OrderModel) Insert(order Order) error {
	// Debug print the order values
	fmt.Printf("Inserting order: %+v\n", order)
	query := `
			INSERT INTO orders (user_id, product_id, quantity, size, price, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.Exec(query,
		order.UserID,
		order.ProductID,
		order.Quantity,
		order.Size,
		order.Price,
		order.CreatedAt,
		order.UpdatedAt,
	)
	fmt.Printf("Inserted order: %+v\n", order)
	return err
}
