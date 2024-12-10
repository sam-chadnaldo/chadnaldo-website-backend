CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,               
    user_id INT REFERENCES users(id) ON DELETE CASCADE,  
    payment_provider_id INT REFERENCES payment_providers(id), 
    tokens_amount DECIMAL(18, 8) NOT NULL,                -- Количество токенов
    payment_currency VARCHAR(10) NOT NULL,               -- Валюта платежа
    total_price DECIMAL(18, 8) NOT NULL,                 -- Общая сумма
    status VARCHAR(50) NOT NULL,                         -- Статус транзакции
    created_at TIMESTAMP DEFAULT NOW(),  
    updated_at TIMESTAMP DEFAULT NOW()   
);