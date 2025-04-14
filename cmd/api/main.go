package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jeanpsv/realworld-golang/internal/repository"
	mysqlRepository "github.com/jeanpsv/realworld-golang/internal/repository/mysql"
	"github.com/jeanpsv/realworld-golang/internal/rest"
	"github.com/jeanpsv/realworld-golang/services"
)

func main() {

	dbConn, err := repository.OpenDB("mysql", "realworld:realworld@/realworld_dev?parseTime=true")
	if err != nil {
		log.Fatal("failed to open connection to database", err)
	}
	defer dbConn.Close()

	router := mux.NewRouter()

	tagRepo := mysqlRepository.NewTagRepository(dbConn)
	tagService := services.NewTagService(tagRepo)
	rest.NewTagHandler(router, tagService)

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
