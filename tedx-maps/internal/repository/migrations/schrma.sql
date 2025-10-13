CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role TEXT DEFAULT 'viewer',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS points (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    type TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS routes (
    id SERIAL PRIMARY KEY,
    from_point INT REFERENCES points(id) ON DELETE CASCADE,
    to_point INT REFERENCES points(id) ON DELETE CASCADE,
    distance_km DOUBLE PRECISION,
    duration TEXT,
    path JSONB,
    created_at TIMESTAMP DEFAULT NOW()
);
