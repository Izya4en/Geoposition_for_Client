mkdir -p auth-service/{cmd,internal/{server,handler,service,repo,model,middleware,config},migrations}
cat > auth-service/.env.example <<EOF
# DB
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=authdb

# JWT
JWT_SECRET=supersecretkeyyoushouldchange
ACCESS_TTL=15m
REFRESH_TTL=168h

# NATS
NATS_URL=nats://nats:4222
EOF
