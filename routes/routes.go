package routes

import (
	"go-rest/controllers"
	"go-rest/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlewares.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/persona", controllers.GetAllPersona).Methods("GET")
	r.HandleFunc("/api/persona/{id}", controllers.GetPersonaById).Methods("GET")
	r.HandleFunc("/api/persona", controllers.CreatePersona).Methods("POST")
	r.HandleFunc("/api/persona/{id}", controllers.DeletePersona).Methods("DELETE")
	r.HandleFunc("/api/persona/{id}", controllers.EditPersona).Methods("PUT")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", handlers.CORS((handlers.AllowedOrigins([]string{"*"})))(r)))
}
