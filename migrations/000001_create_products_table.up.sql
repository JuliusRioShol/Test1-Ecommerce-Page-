CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY, 
    name TEXT  NOT NULL,
    description TEXT NOT NULL,
    brand TEXT NOT NULL,
    image_url TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL CHECK (price >= 0),
    rating NUMERIC(2,1) CHECK (rating >= 0 AND rating <= 5),
    created_at TIMESTAMPTZ DEFAULT NOW()
);