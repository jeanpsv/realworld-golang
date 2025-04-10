package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jeanpsv/realworld-golang/internal/repository/db"
	"github.com/jeanpsv/realworld-golang/internal/rest"
	"github.com/jeanpsv/realworld-golang/tag"
)

func main() {

	dbConn, err := openDB("realworld:realworld@/realworld_dev?parseTime=true")
	if err != nil {
		log.Fatal("failed to open connection to database", err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal("failed to ping database", err)
	}

	defer dbConn.Close()

	tagRepo := db.NewTagRepository(dbConn)

	tagService := tag.NewService(tagRepo)

	router := mux.NewRouter()
	rest.NewTagHandler(router, tagService)

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
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
