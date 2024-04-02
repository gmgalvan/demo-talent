CREATE TABLE budgets (
    id VARCHAR(255) PRIMARY KEY,         
    amount NUMERIC NOT NULL,          
    start_date TIMESTAMP NOT NULL,    
    end_date TIMESTAMP NOT NULL       
);