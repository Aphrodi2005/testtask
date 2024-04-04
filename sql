CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    order_number INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL
);
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    primary_shelf VARCHAR(1) NOT NULL
);
CREATE TABLE shelves (
    shelf_id SERIAL PRIMARY KEY,
    shelf_name VARCHAR(255) NOT NULL
);
CREATE TABLE product_shelf (
    product_id INT NOT NULL,
    shelf_id INT NOT NULL,
    PRIMARY KEY (product_id, shelf_id)
);

INSERT INTO products (name, primary_shelf) VALUES
('Ноутбук', 'А'),
('Телевизор', 'А'),
('Телефон', 'Б'),
('Системный блок', 'Ж'),
('Часы', 'Ж'),
('Микрофон', 'Ж');

INSERT INTO shelves (shelf_name) VALUES
('А'),
('Б'),
('Ж'),
('З'),
('В');

INSERT INTO product_shelf (product_id, shelf_id) VALUES
(1, 1),
(2, 1),
(3, 2),
(3, 4),
(3, 5),
(4, 3),
(5, 3),
(6, 3);

INSERT INTO orders (order_number, product_id, quantity) VALUES
(10, 1, 2),
(10, 3, 1),
(10, 6, 1),
(11, 2, 3),
(14, 1, 3),
(14, 4, 4),
(15, 5, 1);