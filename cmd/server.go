package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/trivektor/gorm-resumes-api/pkg/database"
	"github.com/trivektor/gorm-resumes-api/pkg/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=root password=root dbname=postgres port=5432"
	var err error

	database.DB, err = gorm.Open(postgres.Open(dsn))

	if (err != nil) {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/resumes", handlers.GetResumes).Methods("GET")
	r.HandleFunc("/resumes", handlers.CreateResume).Methods("POST")
	r.HandleFunc("/resumes/{id}", handlers.GetResume).Methods("GET")
	r.HandleFunc("/resumes/{id}", handlers.UpdateResume).Methods("PATCH")
	r.HandleFunc("/resumes/{id}", handlers.DeleteResume).Methods("DELETE")

	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}