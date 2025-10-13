#!/bin/bash

# Создаём корневую папку проекта
mkdir -p tedx-maps
cd tedx-maps || exit

echo "📁 Создаю структуру проекта tedx-maps..."

# Основные директории
mkdir -p cmd/server
mkdir -p config
mkdir -p internal/{entity,dto,repository/{migrations},service,delivery/http,server,utils}
mkdir -p pkg/{mapsapi,auth}
mkdir -p scripts
mkdir -p test
mkdir -p web/{js,css}

# Основные файлы
touch cmd/server/main.go
touch config/{config.yaml,config.go}
touch internal/entity/{point.go,route.go,user.go,mapdata.go}
touch internal/dto/{point_dto.go,route_dto.go,map_dto.go,user_dto.go}
touch internal/repository/{point_repo.go,route_repo.go,user_repo.go,interfaces.go}
touch internal/repository/migrations/schema.sql
touch internal/service/{point_service.go,route_service.go,map_service.go,user_service.go}
touch internal/delivery/http/{handler.go,point_handler.go,route_handler.go,user_handler.go,map_handler.go}
touch internal/server/server.go
touch internal/utils/{logger.go,geo.go,response.go,errors.go,validation.go}
touch pkg/mapsapi/{client.go,serpapi_client.go,responses.go}
touch pkg/auth/{jwt.go,middleware.go}
touch scripts/{build.sh,run.sh,migrate.sh}
touch test/{point_service_test.go,route_service_test.go,api_integration_test.go}
touch web/index.html
touch web/js/map.js
touch web/css/style.css
touch .env go.mod go.sum README.md

echo "✅ Структура проекта tedx-maps успешно создана!"
