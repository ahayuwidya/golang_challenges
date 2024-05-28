DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS variants;

CREATE TABLE products (
  id                INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  product_name      VARCHAR(128) NOT NULL UNIQUE,
  created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO products
  (product_name)
VALUES
  ('Matcha'),
  ('Choco'),
  ('Berry'),
  ('Cheesecake'),
  ('Mango');

-- SELECT * FROM products;

CREATE TABLE variants (
  id                INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  variant_name      VARCHAR(128) NOT NULL UNIQUE,
  quantity          INT NOT NULL,
  product_id        INT NOT NULL,
  created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (product_id) REFERENCES products(id)
);

INSERT INTO variants
  (variant_name, quantity, product_id)
VALUES
  ('Matcha Almond', 100, 1),
  ('Matcha Cheesecake', 80, 1),
  ('Choco Oreo', 15, 2),
  ('Choco Hazelnut', 90, 2),
  ('Berry Lemon', 30, 3),
  ('Strawberry Cheesecake', 70, 4),
  ('Mango Banana', 30, 5);

-- SELECT * FROM variants;