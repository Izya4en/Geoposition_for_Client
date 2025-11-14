CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS terminals (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  latitude DOUBLE PRECISION,
  longitude DOUBLE PRECISION,
  cash_available BIGINT DEFAULT 0,
  status TEXT DEFAULT 'online',
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS reservations (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  terminal_id UUID REFERENCES terminals(id) ON DELETE CASCADE,
  user_id TEXT,
  amount BIGINT NOT NULL,
  status TEXT DEFAULT 'reserved',
  reserved_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
  completed_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS analytics (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  terminal_id UUID REFERENCES terminals(id) ON DELETE CASCADE,
  event_type TEXT,
  payload JSONB,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
