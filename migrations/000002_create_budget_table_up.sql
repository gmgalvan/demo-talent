CREATE TABLE budgets (
    id VARCHAR(255) PRIMARY KEY,
    description TEXT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    start_date VARCHAR(255) NOT NULL,
    end_date DATE NOT NULL
);
