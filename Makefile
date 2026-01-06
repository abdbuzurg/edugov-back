# Makefile for golang-migrate
# Usage examples:
#   make migrate-up
#   make migrate-up N=1
#   make migrate-down
#   make migrate-down N=1
#   make migrate-force V=1
#   make migrate-version
#
# You can override DB_URL and MIGRATIONS_PATH:
#   make migrate-up DB_URL='postgres://user:pass@localhost:5432/db?sslmode=disable'

MIGRATE          ?= migrate
MIGRATIONS_PATH  ?= internal/infrastructure/persistence/postgres/migrations
DB_URL           ?= postgres://postgres:password@localhost:5432/edugov?sslmode=disable

.PHONY: migrate-up migrate-down migrate-force migrate-version migrate-create

# Apply all up migrations, or N steps if N is provided (e.g., make migrate-up N=1)
migrate-up:
	@if [ -n "$(N)" ]; then \
		$(MIGRATE) -path "$(MIGRATIONS_PATH)" -database "$(DB_URL)" up $(N); \
	else \
		$(MIGRATE) -path "$(MIGRATIONS_PATH)" -database "$(DB_URL)" up; \
	fi

# Apply all down migrations, or N steps if N is provided (e.g., make migrate-down N=1)
migrate-down:
	@if [ -n "$(N)" ]; then \
		$(MIGRATE) -path "$(MIGRATIONS_PATH)" -database "$(DB_URL)" down $(N); \
	else \
		$(MIGRATE) -path "$(MIGRATIONS_PATH)" -database "$(DB_URL)" down; \
	fi

# Force-set the migration version (useful after a failed migration leaves DB "dirty")
# Example: make migrate-force V=3
migrate-force:
	@if [ -z "$(V)" ]; then \
		echo "Usage: make migrate-force V=<version>"; \
		exit 1; \
	fi
	$(MIGRATE) -path "$(MIGRATIONS_PATH)" -database "$(DB_URL)" force $(V)

# Show current migration version (and dirty state)
migrate-version:
	$(MIGRATE) -path "$(MIGRATIONS_PATH)" -database "$(DB_URL)" version

# Create new migration files (requires golang-migrate create support)
# Example: make migrate-create NAME=add_materials_constraints
migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make migrate-create NAME=<migration_name>"; \
		exit 1; \
	fi
	$(MIGRATE) create -ext sql -dir "$(MIGRATIONS_PATH)" -seq "$(NAME)"
