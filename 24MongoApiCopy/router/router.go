package router

import (
	"github.com/gorilla/mux"
	"github.com/rohit/controller/controller"
)

func Routers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.InsertOneMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	r.HandleFunc("/api/movies", controller.DeleteAllMovies).Methods("DELETE")

	return r

}
