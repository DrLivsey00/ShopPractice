CREATE TABLE  devices(
    device_id SERIAL PRIMARY KEY,
    device_name TEXT NOT NULL ,
    device_price FLOAT NOT NULL ,
    device_description TEXT default '' NOT NULL,
    number_in_stock INT default 0 NOT NULL
)