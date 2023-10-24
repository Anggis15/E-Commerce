CREATE TABLE transactions
(
    id   VARCHAR(255) NOT NULL,
    user_id_comment VARCHAR(255) NOT NULL,
    product_id VARCHAR(255) NOT NULL,
    comment VARCHAR(255) NOT NULL,
    rating INT NOT NULL,
    created_at VARCHAR(255) NOT NULL,
    updated_at VARCHAR(255) NOT NULL
    PRIMARY KEY (id)
) ENGINE = InnoDB;