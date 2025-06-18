package storage

import (
	"context"
	"fmt"

	"github.com/go-mod-test/goods/internal/domain"
)

func (s *Storage) GetAllInvoices(ctx context.Context) ([]domain.Invoice, error) {
	const op = "storage.GetAllInvoices"
	query := `SELECT id, number, created_at, customer_id FROM invoices`

	rows, err := s.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var invoices []domain.Invoice
	for rows.Next() {
		var invoice domain.Invoice
		if err = rows.Scan(&invoice.ID, &invoice.Number, &invoice.CreatedAt, &invoice.CustomerID); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		invoices = append(invoices, invoice)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return invoices, nil

}

func (s *Storage) GetOneInvoice(ctx context.Context, id int) (domain.Invoice, error) {
	const op = "storage.GetOneInvoice"

	query := `SELECT id, number, created_at, customer_id FROM invoices WHERE id = $1`

	row := s.Db.QueryRowContext(ctx, query, id)

	var invoice domain.Invoice
	if err := row.Scan(&invoice.ID, &invoice.Number, &invoice.CreatedAt, &invoice.CustomerID); err != nil {
		return domain.Invoice{}, fmt.Errorf("%s: %w", op, err)
	}

	err := row.Err()
	if err != nil {
		return domain.Invoice{}, fmt.Errorf("%s: %w", op, err)
	}

	return invoice, nil
}

func (s *Storage) CreateInvoice(ctx context.Context, invoice domain.Invoice) error {
	const op = "storage.CreateInvoice"

	query := `INSERT INTO invoices (number, created_at, customer_id) VALUES ($1, $2, $3)`

	_, err := s.Db.ExecContext(ctx, query, invoice.Number, invoice.CreatedAt, invoice.CustomerID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
func (s *Storage) UpdateInvoice(ctx context.Context, invoice domain.Invoice) error {
	const op = "storage.UpdateInvoice"

	query := `UPDATE invoices SET number = $1, created_at = $2, customer_id = $3 WHERE id = $4`

	_, err := s.Db.ExecContext(ctx, query, invoice.Number, invoice.CreatedAt, invoice.CustomerID, invoice.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
func (s *Storage) DelOneInvoice(ctx context.Context, id int) error {
	const op = "storage.DelOneInvoice"

	query := `DELETE FROM invoices WHERE id = $1`

	_, err := s.Db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
