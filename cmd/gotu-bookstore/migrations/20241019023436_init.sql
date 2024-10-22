-- +goose Up
-- +goose StatementBegin
-- Create users table
CREATE TABLE users (
    id UUID NOT NULL CONSTRAINT users_pk PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    password TEXT,
    status VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE UNIQUE INDEX users_email_unique_index ON users(email)
WHERE deleted_at IS NULL;
-- Create books table
CREATE TABLE books (
    id UUID NOT NULL CONSTRAINT books_pk PRIMARY KEY,
    author VARCHAR(255),
    title VARCHAR(255),
    description TEXT,
    category VARCHAR(255),
    publisher VARCHAR(255),
    price NUMERIC(12, 2),
    isbn VARCHAR(13),
    language VARCHAR(255),
    publish_date TIMESTAMP,
    image_url VARCHAR,
    page INTEGER,
    rating NUMERIC(12, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE UNIQUE INDEX books_isbn_unique_index ON books(isbn)
WHERE deleted_at IS NULL;
-- Create shopping cart table
CREATE TABLE shopping_carts (
    id UUID NOT NULL CONSTRAINT shopping_carts_pk PRIMARY KEY,
    user_id UUID,
    CONSTRAINT shopping_carts_users_id_fk FOREIGN KEY (user_id) REFERENCES users(id),
    book_id UUID,
    CONSTRAINT shopping_carts_books_id_fk FOREIGN KEY (book_id) REFERENCES books(id),
    quantity INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
-- Create transactions table
CREATE TABLE transactions (
    id UUID NOT NULL CONSTRAINT transactions_pk PRIMARY KEY,
    user_id UUID,
    CONSTRAINT transactions_users_id_fk FOREIGN KEY (user_id) REFERENCES users(id),
    grand_total NUMERIC(12, 2),
    status VARCHAR(255),
    invoice_number VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
-- Create transaction_items table
CREATE TABLE transaction_items (
    id UUID NOT NULL CONSTRAINT transaction_items_pk PRIMARY KEY,
    transaction_id UUID,
    CONSTRAINT transactions_transaction_items_id_fk FOREIGN KEY (transaction_id) REFERENCES transactions(id),
    book_id UUID,
    CONSTRAINT transactions_books_id_fk FOREIGN KEY (book_id) REFERENCES books(id),
    quantity INTEGER,
    price NUMERIC(12, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE UNIQUE INDEX transactions_invoice_number_unique_index ON transactions(invoice_number)
WHERE deleted_at IS NULL;
CREATE SEQUENCE invoice_counter_seq START 1000000;
-- Add indexes for better query performance
CREATE INDEX idx_transactions_users ON transactions(user_id);
CREATE INDEX idx_transacition_details_transactions ON transaction_items(transaction_id);
CREATE INDEX idx_transaction_items_books ON transaction_items(book_id);
CREATE INDEX idx_books_title ON books USING gin (to_tsvector('english', title));
CREATE INDEX idx_books_author ON books USING gin (to_tsvector('english', author));
CREATE INDEX idx_books_description ON books USING gin (to_tsvector('english', description));
CREATE INDEX idx_books_isbn ON books USING btree (isbn);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE transaction_items,
transactions,
shopping_carts,
books,
users;
DROP SEQUENCE invoice_counter_seq;
-- +goose StatementEnd