#!/bin/sh

set -e

# load .env
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
else
  echo "❌ .env file not found!"
  exit 1
fi

echo "🚀 start the app"

# migrate sql schema
/app/migrate \
  -path db/migration \
  -database "postgresql://$DATABASE_USERNAME:$DATABASE_PASSWORD@$DATABASE_HOST:5432/$DATABASE_NAME?sslmode=disable" \
  --verbose up

echo "✅ finish migrate sql schema"
exec "$@"
