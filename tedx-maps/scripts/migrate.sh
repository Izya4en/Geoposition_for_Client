#!/bin/bash
echo "🚀 Running database migrations..."

set -e

DB_NAME=${DB_NAME:-tedx_maps}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-postgres}
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}

echo "Applying schema.sql..."
psql "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}" -f ./internal/repository/migrations/schema.sql

echo "✅ Migration complete!"
