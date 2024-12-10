CREATE TABLE payment_providers (
    id SERIAL PRIMARY KEY,            
    name VARCHAR(255) NOT NULL,       
    created_at TIMESTAMP DEFAULT NOW(),  
    updated_at TIMESTAMP DEFAULT NOW()   
);