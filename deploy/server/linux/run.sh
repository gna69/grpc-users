export PORT=8080

export KAFKA_HOST=localhost
export KAFKA_PORT=29092

export CLICKHOUSE_HOST=localhost
export CLICKHOUSE_PORT=9000

export REDIS_HOST=localhost
export REDIS_PORT=6379

export PG_HOST=localhost
export PG_PORT=6001
export PG_DB=users_db
export PG_USER=pguser
export PG_PASS=pg_password

chmod +x main

./main