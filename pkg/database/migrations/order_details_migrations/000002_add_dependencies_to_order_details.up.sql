-- Первый внешний ключ: order_id
ALTER TABLE order_details
    ADD CONSTRAINT fk_order_id
        FOREIGN KEY (order_id) REFERENCES orders(order_id)
            ON DELETE CASCADE
            ON UPDATE CASCADE;


-- Второй внешний ключ: device_id
ALTER TABLE order_details
    ADD CONSTRAINT fk_device_id
        FOREIGN KEY (device_id) REFERENCES devices(device_id)
            ON DELETE CASCADE
            ON UPDATE CASCADE;
