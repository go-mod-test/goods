package storage

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-mod-test/goods/internal/config"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db *sql.DB
}

// GetAllCustomers implements customer.GetAllCustom.

func NewStorage(dbPath config.Db) (*Storage, error) {
	const op = "storage.NewStorage"

	dsn := createDSN(dbPath)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		price NUMERIC(10,2) NOT NULL,
		stock_quantity INT NOT NULL DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS invoices (
		id SERIAL PRIMARY KEY,
		number TEXT NOT NULL UNIQUE,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		customer_id INT NOT NULL REFERENCES customers(id)
	);

	CREATE TABLE IF NOT EXISTS invoice_items (
		id SERIAL PRIMARY KEY,
		invoice_id INT NOT NULL REFERENCES invoices(id) ON DELETE CASCADE,
		product_id INT NOT NULL REFERENCES products(id),
		product_name TEXT NOT NULL,
		product_price NUMERIC(10,2) NOT NULL,
		quantity INT NOT NULL,
		total NUMERIC(10,2) GENERATED ALWAYS AS (product_price * quantity) STORED
	);
	`

	if _, err := db.Exec(query); err != nil {
		return nil, fmt.Errorf("%s: failed to create schema: %w", op, err)
	}

	return &Storage{Db: db}, nil
}

func createDSN(dbPath config.Db) string {

	var dsn strings.Builder //:= "host=localhost port=5432 user=postgres password=secret dbname=mydb sslmode=disable"

	dsn.WriteString("host=")
	dsn.WriteString(dbPath.DBHost)
	dsn.WriteString(" ")

	dsn.WriteString("port=")
	dsn.WriteString(dbPath.DBPort)
	dsn.WriteString(" ")

	dsn.WriteString("user=")
	dsn.WriteString(dbPath.DBUser)
	dsn.WriteString(" ")

	dsn.WriteString("password=")
	dsn.WriteString(dbPath.DBPass)
	dsn.WriteString(" ")

	dsn.WriteString("dbname=")
	dsn.WriteString(dbPath.DBName)
	dsn.WriteString(" ")

	dsn.WriteString("sslmode=")
	dsn.WriteString(dbPath.Sslmode)

	return dsn.String()
}
