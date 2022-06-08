package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/trivektor/gorm-resumes-api/pkg/database"
	"github.com/trivektor/gorm-resumes-api/pkg/models"
)

func CreateResume(w http.ResponseWriter, r *http.Request) {
	var input CreateResumeInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(422)
		return
	}

	resume := models.Resume{Title: input.Title, Description: input.Description}

	result := database.DB.Create(&resume)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	response := make(map[string]int)
	response["resume_id"] = int(resume.Id)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err);
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetResumes(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal([]int{})
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func GetResume(w http.ResponseWriter, r *http.Request) {

}

func UpdateResume(w http.ResponseWriter, r *http.Request) {

}

func DeleteResume(w http.ResponseWriter, r *http.Request) {

}