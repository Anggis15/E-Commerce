CREATE TABLE user_profile
(
    id   VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    gender VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    created_at VARCHAR(255) NOT NULL,
    updated_at VARCHAR(255) NOT NULL,
    is_active INT NOT NULL
    PRIMARY KEY (id)
) ENGINE = InnoDB;