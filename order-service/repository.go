package main

import "database/sql"

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(userID int, product string) (*Order, error) {
	result, err := r.db.Exec(
		"INSERT INTO orders(user_id, product) VALUES (?, ?)",
		userID, product,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(
		"SELECT id, user_id, product, created_at FROM orders WHERE id = ?",
		id,
	)

	var o Order
	if err := row.Scan(&o.ID, &o.UserID, &o.Product, &o.CreatedAt); err != nil {
		return nil, err
	}

	return &o, nil
}

func (r *OrderRepository) List() ([]Order, error) {
	rows, err := r.db.Query(
		"SELECT id, user_id, product, created_at FROM orders ORDER BY id",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.Product, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}
