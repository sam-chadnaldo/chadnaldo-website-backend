CREATE TABLE users (
    id SERIAL PRIMARY KEY,              
    email VARCHAR(255) UNIQUE NOT NULL,  
    ton_wallet_address VARCHAR(255) UNIQUE NOT NULL,  
    created_at TIMESTAMP DEFAULT NOW(),  
    updated_at TIMESTAMP DEFAULT NOW()  
);