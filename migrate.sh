source .env


MIGRATION_DIR="internal/db/migrations"
echo "Migrating database..."
migrate -database "$DATABASE_URL" -path "$MIGRATION_DIR" force 1
migrate -database "$DATABASE_URL" -path "$MIGRATION_DIR" down
migrate -database "$DATABASE_URL" -path "$MIGRATION_DIR" up