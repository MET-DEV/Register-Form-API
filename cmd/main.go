package main

import (
	"net/http"

	"../handlers"
	"../migrations"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	migrations.InitialMigration()
	r := mux.NewRouter()
	r.HandleFunc("/user", handlers.AddUser).Methods("POST")
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8080", handler)
}
