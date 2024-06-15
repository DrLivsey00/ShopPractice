CREATE TYPE status AS ENUM (
    'active',
    'shipment',
    'closed'
    );

CREATE TABLE orders(
    order_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL ,
    order_price FLOAT NOT NULL,
    order_status status DEFAULT 'active' NOT NULL,
    created_at TIMESTAMP NOT NULL ,
    closed_at TIMESTAMP
);