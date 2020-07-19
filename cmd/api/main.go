package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/bgildson/enext-challenge/api/database"
	"github.com/bgildson/enext-challenge/api/handler"
	"github.com/bgildson/enext-challenge/api/repository"
	"github.com/bgildson/enext-challenge/api/service"
)

func main() {
	// parse params to obtain api configurations
	gamesJSONPath := flag.String("games-json-path", "./games.json", "should inform the path for the games json file")
	port := flag.Int64("port", 8080, "indicates which port the api should listen")
	flag.Parse()

	// bind app layers
	db, err := database.NewJSONDatabase(*gamesJSONPath)
	if err != nil {
		log.Fatalf("could not create the database: %v", err)
	}
	r := repository.NewJSONGamesRepository(db)
	s := service.NewGamesService(r)
	h := handler.NewGamesHandler(s)

	// generate http router
	router := chi.NewRouter()

	// add minimum middlewares to log requests and recover from panics
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// bind handlers
	router.Get("/games", h.GetAll)
	router.Get("/games/{id}", h.GetOne)

	// serve api
	addr := fmt.Sprintf(":%d", *port)
	fmt.Println("Running api on", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
