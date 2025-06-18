package domain

import "time"

type Customer struct {
	ID        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}
type Product struct {
	ID            int     `db:"id" json:"id"`
	Name          string  `db:"name" json:"name"`
	Description   string  `db:"description" json:"description"`
	Price         float64 `db:"price" json:"price"` //TODO: При работе с денежными суммами лучше - shopspring/decimal
	StockQuantity int     `db:"stock_quantity" json:"stock_quantity"`
}

type Invoice struct {
	ID         int       `db:"id" json:"id"`
	Number     string    `db:"number" json:"number"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	CustomerID int       `db:"customer_id" json:"customer_id"`
}
type InvoiceItem struct {
	ID           int     `db:"id" json:"id"`
	InvoiceID    int     `db:"invoice_id" json:"invoice_id"`
	ProductID    int     `db:"product_id" json:"product_id"`
	ProductName  string  `db:"product_name" json:"product_name"`
	ProductPrice float64 `db:"product_price" json:"product_price"` //TODO: При работе с денежными суммами лучше - shopspring/decimal
	Quantity     int     `db:"quantity" json:"quantity"`
	Total        float64 `db:"total" json:"total"` //TODO: При работе с денежными суммами лучше - shopspring/decimal
}

type AddItemInput struct {
	InvoiceID int `json:"invoice_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
