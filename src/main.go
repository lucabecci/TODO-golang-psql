package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lucabecci/TODO-golang-psql/src/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	var port string = "8080"

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/").Subrouter()

	apiRouter.HandleFunc("/todos", api.CreateTodo).Methods("POST")
	apiRouter.HandleFunc("/todos/{id}", api.GetTodo).Methods("GET")
	apiRouter.HandleFunc("/todos", api.GetTodos).Methods("GET")
	apiRouter.HandleFunc("/todos/{id}", api.DeleteTodo).Methods("DELETE")
	apiRouter.HandleFunc("/todos/{id}", api.UpdateTodo).Methods("PUT")

	fmt.Printf("Server running at port %s", port)

	http.ListenAndServe(":"+port, router)

}
