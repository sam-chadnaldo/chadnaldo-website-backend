CREATE TABLE tokens (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE, -- Пользователь, владелец токенов
    amount INTEGER CHECK(amount > 0),                   -- Количество токенов
    created_at TIMESTAMP DEFAULT NOW(),                 -- Дата добавления токена
    updated_at TIMESTAMP DEFAULT NOW()                  -- Дата последнего обновления
);