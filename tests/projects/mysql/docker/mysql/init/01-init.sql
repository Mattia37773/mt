USE appDb;

-- Kunden
CREATE TABLE customers (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Kategorien
CREATE TABLE categories (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

-- Produkte
CREATE TABLE products (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    category_id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,

    CONSTRAINT fk_product_category
        FOREIGN KEY (category_id)
        REFERENCES categories(id)
);

-- Bestellungen
CREATE TABLE orders (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    order_date DATETIME DEFAULT NOW(),

    CONSTRAINT fk_order_customer
        FOREIGN KEY (customer_id)
        REFERENCES customers(id)
);

-- Bestellpositionen
CREATE TABLE order_items (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    order_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL,

    CONSTRAINT fk_item_order
        FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_item_product
        FOREIGN KEY (product_id)
        REFERENCES products(id)
);

-- Rollen
CREATE TABLE roles (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- Benutzer
CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    role_id BIGINT NOT NULL,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,

    CONSTRAINT fk_user_role
        FOREIGN KEY (role_id)
        REFERENCES roles(id)
);

--------------------------------------------------
-- Seed Daten
--------------------------------------------------

INSERT INTO categories(name) VALUES
('Books'),
('Electronics'),
('Food');

INSERT INTO products(category_id, name, price) VALUES
(1, 'Docker Handbook', 29.99),
(1, 'Spring Boot Guide', 34.90),
(2, 'Mechanical Keyboard', 89.99),
(2, '27 Inch Monitor', 249.99),
(3, 'Chocolate', 2.49),
(3, 'Coffee', 9.99);

INSERT INTO customers(firstname, lastname, email) VALUES
('Max', 'Mustermann', 'max@test.de'),
('Erika', 'Musterfrau', 'erika@test.de'),
('John', 'Doe', 'john@test.de');

INSERT INTO orders(customer_id) VALUES
(1),
(1),
(2),
(3);

INSERT INTO order_items(order_id, product_id, quantity) VALUES
(1, 1, 1),
(1, 3, 2),
(2, 6, 3),
(3, 2, 1),
(3, 5, 5),
(4, 4, 1);

INSERT INTO roles(name) VALUES
('ADMIN'),
('USER');

INSERT INTO users(role_id, username, password) VALUES
(1, 'admin', 'admin123'),
(2, 'tester', 'tester123'),
(2, 'demo', 'demo123');