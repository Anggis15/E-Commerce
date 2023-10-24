CREATE TABLE products
(
    id   VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    product_gender VARCHAR(255) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    desc_product VARCHAR(255) NOT NULL,
    price VARCHAR(255) NOT NULL,
    stock INT NOT NULL,
    created_at VARCHAR(255) NOT NULL,
    updated_at VARCHAR(255) NOT NULL,
    is_active INT NOT NULL
    PRIMARY KEY (id)
) ENGINE = InnoDB;