package storage

import (
	"context"
	"fmt"

	"github.com/go-mod-test/goods/internal/domain"
)

func (s *Storage) GetAllCustomers(ctx context.Context) ([]domain.Customer, error) {
	const op = "storage.GetAllCastomers"

	query := `SELECT id, first_name, last_name FROM customers`

	rows, err := s.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var customers []domain.Customer
	for rows.Next() {
		var customer domain.Customer
		if err = rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		customers = append(customers, customer)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return customers, nil
}

func (s *Storage) GetOneCustomer(ctx context.Context, id int) (domain.Customer, error) {
	const op = "storage.GetOneCastomer"

	qiery := `SELECT id, first_name, last_name FROM customers WHERE id = $1`

	row := s.Db.QueryRowContext(ctx, qiery, id)

	var customer domain.Customer
	if err := row.Scan(&customer.ID, &customer.FirstName, &customer.LastName); err != nil {
		return domain.Customer{}, fmt.Errorf("%s: %w", op, err)
	}

	err := row.Err()
	if err != nil {
		return domain.Customer{}, fmt.Errorf("%s: %w", op, err)
	}

	return customer, nil
}
func (s *Storage) CreateCustomer(ctx context.Context, customer domain.Customer) error {
	const op = "storage.CreateCastomer"

	query := `INSERT INTO customers (first_name, last_name) VALUES ($1, $2)`

	_, err := s.Db.ExecContext(ctx, query, customer.FirstName, customer.LastName)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
func (s *Storage) UpdateCustomer(ctx context.Context, id int, customer domain.Customer) error {
	const op = "storage.UpdateCastomer"

	query := `UPDATE customers SET first_name = $1, last_name = $2 WHERE id = $3`

	_, err := s.Db.ExecContext(ctx, query, customer.FirstName, customer.LastName, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) DelOneCustomer(ctx context.Context, id int) error {
	const op = "storage.DelOneCastomer"

	query := `DELETE FROM customers WHERE id = $1`

	_, err := s.Db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
