package routes

import (
	"github.com/example/simple-REST/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterFilmStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/Film/", controllers.CreateFilm).Methods("POST")
	router.HandleFunc("/Film/", controllers.GetAllFilms).Methods("GET")
	router.HandleFunc("/Film/{FilmId}", controllers.GetFilmById).Methods("GET")
	router.HandleFunc("/Film/{FilmId}", controllers.UpdateFilm).Methods("PUT")
	router.HandleFunc("/Film/{FilmId}", controllers.DeleteFilm).Methods("DELETE")
}