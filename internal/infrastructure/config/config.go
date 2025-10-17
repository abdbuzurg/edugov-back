package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	//--- APPLICATION SETTINGS ---
	// Sets application mode
	// 2 modes: production and development(default)
	AppEnv string `env:"APP_ENV" env-default:"development"`

	// Default timeout for operations within use cases/business logic (in seconds)
	ContextTimeout int `env:"CONTEXT_TIMEOUT" env-default:"5"`

	// --- SERVER SETTINGS ---
	// Sets port for the server, by default 8080
	Port string `env:"PORT" env-default:"8080"`
	// Max duration for reading the entire request (header + body)
	ReadTimeout int `env:"READ_TIMEOUT" env-default:"15"`
	// Max duration before the server times out writing the response
	WriteTimeout int `env:"WRITE_TIMEOUT" env-default:"15"`
	//Max time to wait for the next request on a keep-alive connection
	IdleTimeout int `env:"IDLE_TIMEOUT" env-default:"60"`
	//Max time to wait for graceful server shutdown
	ShutdownTimeout int `env:"SHUTDOWN_TIMEOUT" env-default:"15"`

	// --- DATABASE SETTINGS ----
	// Full database connection string
	DatabaseURL string `env:"DATABASE_URL" env-required:"true"`
	// Max number of open database connections
	DBMaxOpenConns int `env:"DB_MAX_OPEN_CONSS" env-default:"25"`
	// Max number of idle connections in the pool
	DBMaxIdleConns int `env:"DB_MAX_IDLE_CONNS" env-default:"25"`
	// Max time a connection can be reused (in seconds)
	DBConnMaxLifetime int `env:"DB_CONN_MAX_LIFETIME" env-default:"300"`

	// --- LOGGING SETTINGS ---
	// Logging Level
	LogLevel string `env:"LOG_LEVEL" env-default:"info"`
	//Logging Format
	LogFormat string `env:"LOG_FORMAT" env-default:"text"`

	// ---AUTHENTICATION (JWT) SETTINGS ---
	// Secret key used for signing and verifying JWTs
	JWTAccessSecret  string `env:"JWT_ACCESS_SECRET" env-required:"true"`
	JWTRefreshSecret string `env:"JWT_REFRESH_SECRET" env-required:"true"`
	// JWTs valid state duration (in hours)
	JWTAccessExpiryHours  int `env:"JWT_ACCESS_EXPIRY_HOURS" env-default:"2"`
	JWTRefreshExpiryHours int `env:"JWT_REFRESH_EXPIRY_DAYS" env-default:"7"`
	//Cookies
	CookieSecure bool   `env:"COOKIE_SECURE" env-required:"true"`
	CookieDomain string `env:"COOKIE_DOMAIN" env-required:"true"`

	// --- Authorization (Casbin) Settings
	// Path to the Casbin model definition file
	CasbinModelPath string `env:"CASBIN_MODEL_PATH" env-default:"config/casbin/rbac_model.conf"`
}

func LoadConfig(path string) (*Config, error) {
	cfg := &Config{}

	// Attempt at reading from .env file
	errLocalEnv := cleanenv.ReadConfig(path+"/.env", cfg)
	if errLocalEnv != nil {
		log.Printf("error reading .env config: %v", errLocalEnv)
	}

	// Attempt at reading environmental variables provided
	errProvidedEnv := cleanenv.ReadEnv(cfg)
	if errProvidedEnv != nil {
		log.Printf("error reading environmental variables: %v", errProvidedEnv)
	}

	if errLocalEnv != nil && errProvidedEnv != nil {
		return nil, fmt.Errorf("error no config was loaded")
	}

	// Validation for config
	if cfg.Port == "" {
		return nil, fmt.Errorf("PORT environment variable is required")
	}
	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is required")
	}
	if cfg.JWTAccessSecret == "" {
		return nil, fmt.Errorf("JWT_ACCESS_SECRET environment variable is required")
	}
	if cfg.JWTRefreshSecret == "" {
		return nil, fmt.Errorf("JWT_REFRESH_SECRET environment variable is required")
	}
	if cfg.CookieDomain == "" {
		return nil, fmt.Errorf("COOKIE_DOMAIN environment variable is required")
	}
	if cfg.ContextTimeout <= 0 {
		return nil, fmt.Errorf("CONTEXT_TIMEOUT must be a positive integer")
	}
	if cfg.ReadTimeout <= 0 || cfg.WriteTimeout <= 0 || cfg.IdleTimeout <= 0 || cfg.ShutdownTimeout <= 0 {
		return nil, fmt.Errorf("all server timeouts should be positive integer")
	}
	if cfg.DBMaxOpenConns <= 0 || cfg.DBMaxIdleConns <= 0 || cfg.DBConnMaxLifetime <= 0 {
		return nil, fmt.Errorf("all database connection pool settings should be positive integers")
	}

	return cfg, nil
}
