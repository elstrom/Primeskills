package main

import (
	"comments-api/database"
	"comments-api/handlers"
	"comments-api/models"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// Hubungkan ke database
	database.ConnectDatabase()

	// Migrasikan model ke database
	database.MigrateDatabase(&models.Comment{})

	// Atur router
	r := mux.NewRouter()
	h := handlers.Handler{DB: database.DB}

	// Tambahkan endpoint API
	r.HandleFunc("/comments", h.CreateComment).Methods(http.MethodPost)
	r.HandleFunc("/comments", h.GetComments).Methods(http.MethodGet)
	r.HandleFunc("/comments/{id}", h.DeleteComment).Methods(http.MethodDelete)

	// Jalankan server
	http.ListenAndServe(":8080", r)
}
