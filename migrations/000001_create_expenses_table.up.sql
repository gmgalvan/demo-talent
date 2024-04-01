CREATE TABLE expenses (
    id VARCHAR(255) PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    date_creation BIGINT NOT NULL
);