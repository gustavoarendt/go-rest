package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest/database"
	"go-rest/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func GetAllPersona(w http.ResponseWriter, r *http.Request) {
	var persona []models.Persona
	database.DB.Find(&persona)
	json.NewEncoder(w).Encode(persona)
}

func GetPersonaById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var persona models.Persona
	database.DB.First(&persona, id)
	json.NewEncoder(w).Encode(persona)
}

func CreatePersona(w http.ResponseWriter, r *http.Request) {
	var persona models.Persona
	json.NewDecoder(r.Body).Decode(&persona)
	database.DB.Create(&persona)
	json.NewEncoder(w).Encode(persona)
}

func DeletePersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var persona models.Persona
	database.DB.First(&persona, id)
	database.DB.Delete(&persona)
	json.NewEncoder(w).Encode(persona)
}

func EditPersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var persona models.Persona
	database.DB.First(&persona, id)
	json.NewDecoder(r.Body).Decode(&persona)
	database.DB.Save(&persona)
	json.NewEncoder(w).Encode(persona)
}
