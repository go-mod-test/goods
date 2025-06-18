package storage

import (
	"context"
	"fmt"

	"github.com/go-mod-test/goods/internal/domain"
)

func (s *Storage) GetAllProductFromInvoice(ctx context.Context, invoiceId int) ([]domain.InvoiceItem, error) {
	const op = "storage.GetAllProductFromInvoice"

	query := `SELECT id, invoice_id, poroduct_id, product_name, product_price, quantity FROM invoice_items WHERE invoice_id = $1`

	rows, err := s.Db.QueryContext(ctx, query, invoiceId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var invoiceItems []domain.InvoiceItem
	for rows.Next() {
		var invoiceItem domain.InvoiceItem
		if err = rows.Scan(&invoiceItem.ID, &invoiceItem.InvoiceID, &invoiceItem.ProductID, &invoiceItem.ProductName, &invoiceItem.ProductPrice, &invoiceItem.Quantity); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		invoiceItems = append(invoiceItems, invoiceItem)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return invoiceItems, nil
}

func (s *Storage) CreateProductInInvoice(ctx context.Context, input domain.AddItemInput) error {
	const op = "storage.CreateProductInInvoice"

	var name string
	var price float64

	queryProduct := `SELECT name, price FROM products WHERE id = $1`
	rows, err := s.Db.QueryContext(ctx, queryProduct, input.ProductID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&name, &price); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	query := `INSERT INTO invoice_items (invoice_id, product_id, product_name, product_price, quantity) VALUES ($1, $2, $3, $4, $5)`
	_, err = s.Db.ExecContext(ctx, query, input.InvoiceID, input.ProductID, name, price, input.Quantity)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) DelProductFromInvoice(ctx context.Context, invoiceId int) error {
	const op = "storage.DelProductFromInvoice"

	query := `DELETE FROM invoice_items WHERE invoice_id = $1`

	_, err := s.Db.ExecContext(ctx, query, invoiceId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
