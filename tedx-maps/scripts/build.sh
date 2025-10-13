#!/bin/bash
echo "🔧 Building tedx-maps..."

# Выходим при ошибках
set -e

# Очистка старых бинарников
rm -f ./bin/tedx-maps

# Сборка проекта
go build -o ./bin/tedx-maps ./cmd/server/main.go

echo "✅ Build complete! Binary saved to ./bin/tedx-maps"
