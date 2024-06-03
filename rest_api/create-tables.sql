-- DROP TABLE orders;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS items;

CREATE TABLE orders;
CREATE TABLE items;

-- CREATE TABLE orders (
--     id INTEGER PRIMARY KEY,              -- Auto-incrementing primary key
--     customer_name VARCHAR(255) NOT NULL, -- Name of the customer
--     ordered_at TIMESTAMP NOT NULL,      -- Time when the order was placed
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Record creation timestamp
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Record update timestamp
-- );

-- -- Create the items table (child table)
-- CREATE TABLE items (
--     id SERIAL PRIMARY KEY,              -- Auto-incrementing primary key
--     item_name VARCHAR(255) NOT NULL,         -- Name of the item
--     item_desc TEXT,                   -- Description of the item
--     quantity INTEGER NOT NULL,          -- Quantity of the item
--     order_id INTEGER NOT NULL,          -- Foreign key referencing the orders table
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Record creation timestamp
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Record update timestamp
--     FOREIGN KEY (order_id) REFERENCES orders(id) -- Foreign key constraint
-- );