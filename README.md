## **Go Backend Project Structure**

```
├── cmd
│   └── app
│       └── main.go                  # Application entry point: loads config, initializes DB, runs migrations, wires dependencies (repositories, use cases, handlers), starts server.
├── internal
│   ├── domain                       # Core business rules and entities (Innermost Layer)
│   │   ├── entities
│   │   │   ├── user.go              # Pure Go structs representing business entities (e.g., User, Product, Order).
│   │   │   ├── product.go           # Example Product entity.
│   │   │   └── order.go             # Example Order entity.
│   │   └── repositories             # Interfaces defining contracts for data persistence operations (abstractions).
│   │       ├── user_repository.go   # Interface for User data operations (e.g., Create, GetByID).
│   │       ├── product_repository.go# Interface for Product data operations (e.g., GetByID, UpdateStock).
│   │       └── order_repository.go  # Interface for Order data operations (e.g., Create, GetByID).
│   ├── application                  # Application-specific business logic (Use Cases / Interactors)
│   │   ├── usecases
│   │   │   ├── user_usecase.go      # User-related business logic, depends on `repositories.UserRepository` and `postgres.Store`.
│   │   │   └── order_usecase.go     # Example: Order-related business logic (e.g., PlaceOrder), depends on multiple repositories (User, Product, Order) and `postgres.Store` for transactions.
│   │   └── services                 # (Optional) For complex domain logic that doesn't fit within a single entity or use case directly.
│   │       └── some_domain_service.go
│   ├── infrastructure               # External concerns (Interface Adapters, Frameworks/Drivers)
│   │   ├── persistence              # Database implementations and schema management
│   │   │   ├── postgres             # PostgreSQL-specific implementations
│   │   │   │   ├── sqlc             # Generated Go code by SQLC (DO NOT EDIT MANUALLY)
│   │   │   │   │   ├── models.go    # Go structs representing database tables (e.g., User, Product, Order).
│   │   │   │   │   ├── queries.go   # Go functions for executing SQL queries (e.g., `GetUserByID`, `CreateUser`).
│   │   │   │   │   └── db.go        # SQLC database connection helper.
│   │   │   │   ├── query            # SQL files for SQLC to generate code from, organized by entity.
│   │   │   │   │   ├── user.sql     # SQL queries for the User entity.
│   │   │   │   │   ├── product.sql  # SQL queries for the Product entity.
│   │   │   │   │   └── order.sql    # SQL queries for the Order entity.
│   │   │   │   ├── store.go         # Implementation of the "Store" pattern; holds `sqlc.Queries` and provides `ExecTx` for transaction management.
│   │   │   │   ├── user_pg_repository.go   # Concrete implementation of `domain.UserRepository` using SQLC.
│   │   │   │   ├── product_pg_repository.go# Concrete implementation of `domain.ProductRepository` using SQLC.
│   │   │   │   └── order_pg_repository.go  # Concrete implementation of `domain.OrderRepository` using SQLC.
│   │   │   ├── migrations           # Database migration files (managed by `golang-migrate/migrate`).
│   │   │   │   ├── 000001_create_users_table.up.sql   # SQL to apply schema changes for users.
│   │   │   │   ├── 000001_create_users_table.down.sql # SQL to revert user table creation.
│   │   │   │   ├── 000002_create_products_table.up.sql# SQL to apply schema changes for products.
│   │   │   │   └── 000002_create_products_table.down.sql# SQL to revert product table creation.
│   │   │   └── inmemory             # (Optional) In-memory repository implementations, typically for faster testing or development.
│   │   │       └── user_inmemory_repository.go
│   │   ├── http                     # HTTP server, routing, and request/response handling (Delivery Mechanism)
│   │   │   ├── handlers
│   │   │   │   ├── user_handler.go  # HTTP request handlers for user-related endpoints; call `application.usecases.UserUsecase`.
│   │   │   │   └── order_handler.go # HTTP request handlers for order-related endpoints; call `application.usecases.OrderUsecase`.
│   │   │   └── middleware           # HTTP middleware functions, organized into separate files.
│   │   │       ├── logging.go       # Middleware for request logging.
│   │   │       ├── auth.go          # Middleware for authentication (e.g., JWT validation).
│   │   │       ├── recovery.go      # Middleware for panic recovery.
│   │   │       └── request_id.go    # Middleware for injecting a unique request ID.
│   │   ├── auth                     # Authentication/Authorization logic (e.g., JWT token parsing, user session management).
│   │   │   └── jwt.go
│   │   └── config                   # Application configuration loading (e.g., from .env files, YAML, environment variables).
│   │       └── config.go            # Structs and functions for loading configuration.
│   └── shared                       # Reusable components shared across layers (common utilities, custom errors).
│       ├── errors
│       │   └── errors.go            # Custom application-level error types (e.g., NotFoundError, ConflictError, BadRequestError) and helper functions.
│       └── utils
│           └── validator.go         # Utility functions for input validation.
├── pkg                              # Publicly consumable libraries (less common for typical microservices).
│   └── utils
│       └── logger.go                # A reusable logger utility (if you need it outside `internal`).
├── sqlc.yaml                        # SQLC configuration file (at project root) - defines schema and query paths, output settings.
├── go.mod                           # Go modules file - manages project dependencies.
├── go.sum                           # Go modules checksums - ensures integrity of dependencies.
└── Dockerfile                       # Dockerfile for containerizing the application.
```

---

### **Key Configuration Files & Directories Explained:**

* **`cmd/app/main.go`**: This is where your application starts. It's responsible for:
    * Loading configuration.
    * Initializing the database connection.
    * **Running database migrations** using `golang-migrate/migrate`.
    * Wiring up all dependencies (repositories, use cases, handlers).
    * Starting the HTTP server.
* **`internal/infrastructure/persistence/migrations/*.up.sql`**: These files define your database schema changes over time. They are the source of truth for SQLC to understand your database schema when generating code.
* **`internal/infrastructure/persistence/postgres/query/*.sql`**: This is where you'll write your actual SQL queries (e.g., `GetUserByID`, `CreateUser`). Each SQL file can correspond to a specific entity or logical group of operations. SQLC reads these files and generates type-safe Go code from them.
* **`internal/infrastructure/persistence/postgres/sqlc/`**: This directory is *generated* by SQLC. You generally don't edit files here directly. It contains:
    * **`models.go`**: Go structs representing your database tables (e.g., `User`).
    * **`queries.go`**: Type-safe Go functions to execute the SQL queries defined in your `*.sql` files.
* **`internal/infrastructure/persistence/postgres/user_pg_repository.go`**: This file implements the `domain/repositories/UserRepository` interface. It uses the functions generated by SQLC to interact with the database, bridging the gap between your application's use cases and the actual database.
* **`sqlc.yaml`**: Located at the project root, this configuration file tells SQLC where to find your schema (from migrations) and your queries, and where to output the generated Go code.

This structure provides a robust and maintainable foundation for your Go backend, adhering to Clean Architecture principles while leveraging powerful tools like SQLC and `golang-migrate` for efficient database interaction and schema management.
