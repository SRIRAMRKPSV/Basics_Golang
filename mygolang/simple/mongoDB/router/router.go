package router

import (
	"simple/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movie", controller.GetAll).Methods("GET")
	router.HandleFunc("/api/movie", controller.Create).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAll).Methods("PUT")
	router.HandleFunc("/api/movie", controller.Delete).Methods("DELETE")

	return router
}
