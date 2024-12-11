# Variables
DB_URL = postgres://myUser:mysecret_password@localhost:5432/myDatabase
MIGRATION_DIR = ./internal/database/migrations
GOOSE_CMD = goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)"

# Commands
create_table:
	@echo "Creating table..."
	# Replace with the SQL commands to create a table
	psql $(DB_URL) -c "CREATE TABLE example_table (id SERIAL PRIMARY KEY, name VARCHAR(100));"

migrate_up:
	@echo "Running migrations up..."
	$(GOOSE_CMD) up

migrate_down:
	@echo "Rolling back last migration..."
	$(GOOSE_CMD) down

migrate_status:
	@echo "Checking migration status..."
	$(GOOSE_CMD) status

create_migration:
	@echo "Creating new migration..."
	$(GOOSE_CMD) create $(name) sql

reset_migrations:
	@echo "Resetting migrations..."
	$(GOOSE_CMD) reset

# Help command to list available commands
help:
	@echo "Makefile commands:"
	@echo "  make create_table       - Create a table in the database"
	@echo "  make migrate_up         - Apply all up migrations"
	@echo "  make migrate_down       - Rollback last migration"
	@echo "  make migrate_status     - Show migration status"
	@echo "  make create_migration   - Create a new migration file (provide name with 'name' variable)"
	@echo "  make reset_migrations   - Reset all migrations"

# goose -dir ./internal/database/migrations postgres "postgres://myUser:mysecret_password@localhost:5432/myDatabase" up
# goose -dir ./internal/database/migrations create add-some-table sq

