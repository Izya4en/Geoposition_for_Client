#!/bin/bash
echo "🌍 Starting TEDx Maps server..."

set -e

# Проверяем, что .env существует
if [ ! -f .env ]; then
  echo "⚠️  No .env file found. Please create it before running the server."
  exit 1
fi

# Экспортируем переменные окружения
export $(grep -v '^#' .env | xargs)

# Запуск приложения
go run ./cmd/server/main.go
