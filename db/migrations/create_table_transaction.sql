CREATE TABLE transactions
(
    id   VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    product_id VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    total_price VARCHAR(255) NOT NULL,
    created_at VARCHAR(255) NOT NULL,
    payment VARCHAR(255) NOT NULL
    PRIMARY KEY (id)
) ENGINE = InnoDB;