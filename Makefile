# --- Config ---
MIGRATE_BIN ?= migrate
MIGRATIONS_PATH ?= internal/infrastructure/persistence/postgres/migrations

# Default DB connection (override via `make ... DB_URL=...`)
DB_URL ?= postgresql://postgres:password@localhost:5432/edugov?sslmode=disable

# --- Helpers ---
.PHONY: help
help:
	@echo "Migration commands:"
	@echo "  make migrate-up               Run all up migrations"
	@echo "  make migrate-down             Run all down migrations (prompts in some versions)"
	@echo "  make migrate-drop             Drop everything in DB (danger!)"
	@echo "  make migrate-force VERSION=V  Force set migration version"
	@echo "  make migrate-version          Print current migration version"
	@echo "  make migrate-goto VERSION=V   Migrate to a specific version"
	@echo "  make migrate-steps N=1        Move N steps (use N=-1 to go down)"
	@echo "  make migrate-create NAME=x    Create new migration files"
	@echo ""
	@echo "Overrides:"
	@echo "  DB_URL=... MIGRATIONS_PATH=... MIGRATE_BIN=..."

# --- Migrations ---
.PHONY: migrate-up
migrate-up:
	$(MIGRATE_BIN) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	$(MIGRATE_BIN) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down

.PHONY: migrate-drop
migrate-drop:
	$(MIGRATE_BIN) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" drop

.PHONY: migrate-force
migrate-force:
	@if [ -z "$(VERSION)" ]; then echo "ERROR: provide VERSION. Example: make migrate-force VERSION=5"; exit 1; fi
	$(MIGRATE_BIN) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $(VERSION)

.PHONY: migrate-version
migrate-version:
	$(MIGRATE_BIN) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version

.PHONY: migrate-goto
migrate-goto:
	@if [ -z "$(VERSION)" ]; then echo "ERROR: provide VERSION. Example: make migrate-goto VERSION=5"; exit 1; fi
	$(MIGRATE_BIN) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" goto $(VERSION)

.PHONY: migrate-steps
migrate-steps:
	@if [ -z "$(N)" ]; then echo "ERROR: provide N. Example: make migrate-steps N=1 (or N=-1)"; exit 1; fi
	$(MIGRATE_BIN) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" steps $(N)

.PHONY: migrate-create
migrate-create:
	@if [ -z "$(NAME)" ]; then echo "ERROR: provide NAME. Example: make migrate-create NAME=add_employee_fields"; exit 1; fi
	$(MIGRATE_BIN) create -ext sql -dir $(MIGRATIONS_PATH) -seq $(NAME)

