CREATE TABLE transactions
(
    id   VARCHAR(255) NOT NULL,
    product_id VARCHAR(255) NOT NULL,
    promo_name VARCHAR(255) NOT NULL,
    promo_desc VARCHAR(255) NOT NULL,
    promo INT NOT NULL,
    created_at VARCHAR(255) NOT NULL,
    updated_at VARCHAR(255) NOT NULL,
    is_active int NOT NULL
    PRIMARY KEY (id)
) ENGINE = InnoDB;