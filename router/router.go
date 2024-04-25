package router

import (
	"github.com/gorilla/mux"

	"mongoapi/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/movie", controller.GetMyAllMovie).Methods("GET")
	router.HandleFunc("/api/movies/{id}", controller.GetMyMovie).Methods("GET")
	router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")
	return router
}
