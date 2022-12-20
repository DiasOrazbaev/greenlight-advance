package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"os"
	"time"
)

// Declare a string containing the application version number
const version = "1.0.0"

// Define a config struct.
type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  int
	}
}

// Define an application struct to hold dependencides for our HTTP handlers, helpers, and
// middleware.
type application struct {
	config config
	logger *log.Logger
	pool   *pgxpool.Pool
}

func main() {
	// Declare an instance of the config struct.
	var cfg config

	// Read the value of the port and env command-line flags into the config struct.
	// We default to using the port number 4000 and the environment "development" if no
	// corresponding flags are provided.
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment (development|staging|production")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "postgres dsn for connection")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conn", 100, "max open connection of pgx connection pool")
	flag.IntVar(&cfg.db.maxIdleTime, "db-max-idle-time", 30, "max idle time")

	flag.Parse()

	// Initialize a new logger which writes messages to the STDOUT stream, prefixed
	// with the current date and time.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	pool, err := pgxpool.New(context.Background(), cfg.db.dsn)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Declare an instance of the application struct, containing the config struct and the logger.
	app := &application{
		config: cfg,
		logger: logger,
		pool:   pool,
	}

	// Use the httprouter instance returned by app.routes as the server handler.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server.
	logger.Printf("starting the %s server on %s", cfg.env, srv.Addr)
	if err = srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
