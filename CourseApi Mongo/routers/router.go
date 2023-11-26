package routers

import (
	"github.com/gorilla/mux"
	"github.com/rohit/courseApi/controller"
)

func Routers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/courses", controller.GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", controller.GetOneCourse).Methods("GET")
	r.HandleFunc("/course", controller.CreateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", controller.UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", controller.DeleteOneCourse).Methods("DELETE")
	r.HandleFunc("/courses", controller.DeleteAllCourses).Methods("DELETE")

	return r

}
