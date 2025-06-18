package storage

import (
	"context"
	"fmt"

	"github.com/go-mod-test/goods/internal/domain"
)

func (s *Storage) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	const op = "storage.GetAllProducts"

	query := `SELECT id, name, description, price, stock_quantity FROM products`

	rows, err := s.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.StockQuantity); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		products = append(products, product)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return products, nil
}

func (s *Storage) GetOneProduct(ctx context.Context, id int) (domain.Product, error) {
	const op = "storage.GetOneProduct"

	query := `SELECT id, name, description, price, stock_quantity FROM products WHERE id = $1`

	row := s.Db.QueryRowContext(ctx, query, id)

	var product domain.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.StockQuantity); err != nil {
		return domain.Product{}, fmt.Errorf("%s: %w", op, err)
	}

	err := row.Err()
	if err != nil {
		return domain.Product{}, fmt.Errorf("%s: %w", op, err)
	}

	return product, nil
}

func (s *Storage) CreateProduct(ctx context.Context, product domain.Product) error {
	const op = "storage.CreateProduct"

	query := `INSERT INTO products (name, description, price, stock_quantity) VALUES ($1, $2, $3, $4)`

	_, err := s.Db.ExecContext(ctx, query, product.Name, product.Description, product.Price, product.StockQuantity)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
func (s *Storage) UpdateProduct(ctx context.Context, product domain.Product) error {
	const op = "storage.UpdateProduct"

	query := `UPDATE products SET name = $1, description = $2, price = $3, stock_quantity = $4 WHERE id = $5`

	_, err := s.Db.ExecContext(ctx, query, product.Name, product.Description, product.Price, product.StockQuantity, product.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
func (s *Storage) DelOneProduct(ctx context.Context, id int) error {
	const op = "storage.DelOneProduct"

	query := `DELETE FROM products WHERE id = $1`

	_, err := s.Db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
