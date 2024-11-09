-- Active: 1730083286169@@127.0.0.1@5432@inventarikantor@public


CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);




CREATE TABLE inventory_items (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    name VARCHAR(200) NOT NULL,
    photo_url VARCHAR(255),
    price DECIMAL(10, 2) NOT NULL,
    purchase_date DATE DEFAULT CURRENT_DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tokens (
    token_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

DROP TABLE tokens;


-- Kategori Barang

SELECT * FROM categories;

INSERT INTO categories (name, description) VALUES ('Elektronik', 'teknologi pembangkit listrik');

SELECT * FROM categories WHERE id = 1;

UPDATE categories SET name = 'Peralatan Elektronik', updated_at = CURRENT_TIMESTAMP WHERE id = 1;

DELETE FROM categories WHERE id = 1;

-- Barang Inventaris

SELECT * FROM inventory_items;

INSERT INTO inventory_items (category_id, name, photo_url, price, purchase_date) 
VALUES (1, 'Laptop Advan', 'https://example.com/laptop.jpg', 15000000, '2024-03-10');

SELECT * FROM inventory_items WHERE id = 1;

UPDATE inventory_items 
SET name = 'Laptop Dell Inspiron', price = 14000000 , updated_at = CURRENT_TIMESTAMP
WHERE id = 1;

DELETE FROM inventory_items WHERE id = 1;

SELECT * FROM inventory_items WHERE total_usage_days > 100;

--  Laporan Investasi dan Depresiasi

SELECT 
    SUM(price * POWER(0.9, DATE_PART('month', AGE(CURRENT_DATE, purchase_date)))) AS total_investment
FROM 
    inventory_items;


SELECT id, name, price * POWER(0.9, DATE_PART('month', AGE(CURRENT_DATE, purchase_date))) AS current_value, POWER(0.9, DATE_PART('month', AGE(CURRENT_DATE, purchase_date))) AS mounth
FROM inventory_items 
WHERE id = 1;


SELECT id, name, 
0.9 * 10000000 * DATE_PART('month', AGE(CURRENT_DATE, purchase_date::DATE)) 
    + (DATE_PART('year', AGE(CURRENT_DATE, purchase_date::DATE)) * 12) AS total_months
FROM inventory_items;


INSERT INTO users (username, password, email)
VALUES
    ('admin1', 'pass123', 'admin1@example.com');

SELECT * FROM users;

INSERT INTO tokens (user_id, token, expires_at)
VALUES
    (1, 'token_admin_1', NOW() + INTERVAL '30 day');


SELECT id, name, description FROM categories ORDER BY id;