-- Покупатели
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL
);

-- Товары
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price NUMERIC(10,2) NOT NULL,
    stock_quantity INT NOT NULL DEFAULT 0 --Количество на складе
);

-- Накладные
CREATE TABLE IF NOT EXISTS invoices (
    id SERIAL PRIMARY KEY,
    number TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    customer_id INT NOT NULL REFERENCES customers(id)
);

-- Позиции в накладной
CREATE TABLE IF NOT EXISTS invoice_items (
    id SERIAL PRIMARY KEY,
    invoice_id INT NOT NULL REFERENCES invoices(id) ON DELETE CASCADE,
    product_id INT NOT NULL REFERENCES products(id),
    product_name TEXT NOT NULL,   --намеренно дублирую, как назывался товар на момент продажи
    product_price NUMERIC(10,2) NOT NULL, -- намеренно дублирую, какая цена была на момент продажи
    quantity INT NOT NULL, --количесто товаров
    total NUMERIC(10,2) GENERATED ALWAYS AS (product_price * quantity) STORED
);