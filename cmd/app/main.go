package main

import (
	"backend/internal/application/usecases"
	"backend/internal/infrastructure/config"
	"backend/internal/infrastructure/http/handlers"
	"backend/internal/infrastructure/http/middleware"
	"backend/internal/infrastructure/persistence/postgres"
	"backend/internal/infrastructure/security"
	"backend/internal/shared/utils"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// ---- Initialize Config ----
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}
	log.Printf("Configuration loaded Successfully. AppEnv: %s, Port: %s", cfg.AppEnv, cfg.Port)

	// ---- Initialize Database Connection ----
	//Background context for initial database connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Postgres instance
	pool, err := postgres.NewPostgresDB(ctx, cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer pool.Close()
	log.Println("Database connection established")

	// --- Database Migrations
	log.Println("Applying database migrations...")
	runMigrations(cfg.DatabaseURL)
	log.Println("Database migration applying successfully")

	// ---- Initilization of Store ----
	store := postgres.NewStore(pool)

	// ---- Initialization of Repositories ----
	userRepo := postgres.NewPGUserRepository(store)
	userSessionRepo := postgres.NewPGUserSessionRepository(store)
	employeeRepo := postgres.NewPgEmployeeRepository(store)
	employeeDetailsRepo := postgres.NewPGEmployeeDetailsRepository(store)
	employeeDegreeRepo := postgres.NewPgEmployeeDegreeRepository(store)
	employeeWorkExperienceRepo := postgres.NewPgEmployeeWorkExperienceRepository(store)
	employeePublicationRepo := postgres.NewPgEmployeePublicationRepository(store)
	employeeScientificAwardRepo := postgres.NewPgEmployeeScientificAwardRepository(store)

	// ---- Initilization of Security Components
	tokenManager := security.NewTokenManager(
		[]byte(cfg.JWTAccessSecret),
		[]byte(cfg.JWTRefreshSecret),
		time.Duration(cfg.JWTAccessExpiryHours)*60*60*time.Second,
		time.Duration(cfg.JWTRefreshExpiryHours)*24*60*60*time.Second,
	)
	validator := validator.New()

	// ---- Initialization of Use Cases ----
	authUC := usecases.NewAuthUsecase(userRepo, userSessionRepo, employeeRepo, store, tokenManager, validator)
	employeeUC := usecases.NewEmployeeUsecase(employeeRepo, store, validator)
	employeeDetailsUC := usecases.NewEmployeeDetailsUsecase(employeeDetailsRepo, store, validator)
	employeeDegreeUC := usecases.NewEmployeeDegreeUsecase(employeeDegreeRepo, validator)
	employeeWorkExperienceUC := usecases.NewEmployeeWorkExperienceUsecase(employeeWorkExperienceRepo, validator)
	employeePublicationUC := usecases.NewEmployeePublicationUsecase(employeePublicationRepo, validator)
	employeeScientificAwardUC := usecases.NewEmployeeScientificAwardUsecase(employeeScientificAwardRepo, validator)

	// ---- Initialization of HTTP Handlers ----
	authHandlers := handlers.NewAuthHandler(authUC, cfg.CookieDomain, cfg.CookieSecure)
	employeeHandlers := handlers.NewEmployeeHandler(employeeUC)
	employeeDetailsHandler := handlers.NewEmployeeDetailsHandler(employeeDetailsUC)
	employeeDegreeHandler := handlers.NewEmployeeDegreeHandler(employeeDegreeUC)
	employeeWorkExperienceHandler := handlers.NewEmployeeWorkExperienceHandler(employeeWorkExperienceUC)
	employeePublicationHandler := handlers.NewEmployeePublicationHandler(employeePublicationUC)
	employeeScientificAwardHandler := handlers.NewEmployeeScientificAwardHandler(employeeScientificAwardUC)

	// --- Initilization of Routes
	authMiddleware := middleware.CreateAuthMiddleware(tokenManager, utils.RespondWithError)
	mainMux := http.NewServeMux()
	mainMux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, r, http.StatusOK, map[string]string{"ping": "pong"})
	})

	// Auth Routes
	authMux := http.NewServeMux()
	authMux.HandleFunc("POST /register", authHandlers.Register)
	authMux.HandleFunc("POST /login", authHandlers.Login)
	authMux.HandleFunc("POST /refresh-token", authHandlers.RefreshToken)
	authMux.HandleFunc("POST /logout", authHandlers.Logout)
	mainMux.Handle("/auth/", http.StripPrefix("/auth", authMux))

	//Employee Handlers
	employeeMux := http.NewServeMux()

	employeeMux.HandleFunc("GET /{uid}", employeeHandlers.GetByUID)
	// ---- employee/detials
	employeeMux.HandleFunc("PUT /details", authMiddleware(employeeDetailsHandler.Update))
	employeeMux.HandleFunc("GET /details/{employeeID}", employeeDetailsHandler.GetByEmployeeID)
	// ---- employee/degree
	employeeMux.HandleFunc("GET /degree/{employeeID}", employeeDegreeHandler.GetByEmployeeIDAndLanguageCode)
	employeeMux.HandleFunc("PUT /degree", authMiddleware(employeeDegreeHandler.Update))
	employeeMux.HandleFunc("POST /degree", authMiddleware(employeeDegreeHandler.Create))
	employeeMux.HandleFunc("DELETE /degree/{id}", authMiddleware(employeeDegreeHandler.Delete))
	// ---- employee/work-experience
	employeeMux.HandleFunc("GET /work-experience/{employeeID}", employeeWorkExperienceHandler.GetByEmployeeIDAndLanguageCode)
	employeeMux.HandleFunc("POST /work-experience", employeeWorkExperienceHandler.Create)
	employeeMux.HandleFunc("PUT /work-experience", employeeWorkExperienceHandler.Update)
	employeeMux.HandleFunc("DELETE /work-experience/{id}", employeeWorkExperienceHandler.Delete)
	// ---- employee/publication
	employeeMux.HandleFunc("GET /publication/{employeeID}", employeePublicationHandler.GetByEmployeeIDAndLanguageCode)
	employeeMux.HandleFunc("POST /publication", employeePublicationHandler.Create)
	employeeMux.HandleFunc("PUT /publication", employeePublicationHandler.Update)
	employeeMux.HandleFunc("DELETE /publication/{id}", employeePublicationHandler.Delete)
	// ---- employee/scientific-award
	employeeMux.HandleFunc("GET /scientific-award/{employeeID}", employeeScientificAwardHandler.GetByEmployeeIDAndLanguageCode)
	employeeMux.HandleFunc("POST /scientific-award", employeeScientificAwardHandler.Create)
	employeeMux.HandleFunc("PUT /scientific-award", employeeScientificAwardHandler.Update)
	employeeMux.HandleFunc("DELETE /scientific-award/{id}", employeeScientificAwardHandler.Delete)

	mainMux.Handle("/employee/", http.StripPrefix("/employee", employeeMux))

	// ---- Server initialization ----
	mainMiddlewareStack := middleware.CreateMiddlewareStack(
		middleware.PanicRecoveryMiddleware(utils.RespondWithError),
		middleware.CORSMiddleware(middleware.DefaultCORSConfig),
		middleware.RequestIDMiddleware,
		middleware.LanguageMiddleware,
	)
	srv := &http.Server{
		Addr:         "localhost:" + cfg.Port,
		Handler:      mainMiddlewareStack(mainMux),
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
	}

	go func() {
		log.Printf("Server starting on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
		log.Println("Server Stopped gracefully")
	}()

	// --- Graceful Shutdown ---
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(cfg.ShutdownTimeout)*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}

func runMigrations(dbURL string) {
	m, err := migrate.New(
		"file://internal/infrastructure/persistence/postgres/migrations",
		dbURL,
	)

	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run database migrations: %v", err)
	}
	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Printf("Could not get database version: %v", err)
	} else if err == migrate.ErrNilVersion {
		log.Fatalf("Database is at version 0 (no migration applied yet)")
	} else {
		log.Printf("Database version: %d (dirty: %t)", version, dirty)
	}
}
