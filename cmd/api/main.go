package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jeanpsv/realworld-golang/internal/data"
	"github.com/jeanpsv/realworld-golang/internal/router"
)

const version = "1.0.0"

func main() {
	var config router.HttpConfig

	flag.IntVar(&config.Port, "port", 4000, "API Server Port")
	flag.StringVar(&config.Env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&config.Db.Dsn, "db-dsn", "realworld:realworld@/realworld_dev?parseTime=true", "MySQL DSN")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(config)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()
	logger.Info("database connection pool established")

	app := &router.HttpApplication{
		Config: config,
		Logger: logger,
		Models: data.NewModels(db),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 10 * time.Minute,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}
	logger.Info("starting server", "addr", server.Addr, "env", config.Env)

	err = server.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(config router.HttpConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.Db.Dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
