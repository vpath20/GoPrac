package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Database

var courses []Course

func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("=== Api === ")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "2", CourseName: "Java", CoursePrice: 254, Author: &Author{FullName: "Rohit varshney", Website: "github.com/rohit"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "JavaScript", CoursePrice: 199, Author: &Author{FullName: "Rajat sharma", Website: "github.com/rajat"}})
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Server Created Successfully</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all courses details")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting particular course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Can't find course with this id")

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please enter some data")
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	rand.Seed(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100))
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating the particular course")
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == param["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = param["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode("Course updated successfully")
			break
		}
	}
	json.NewEncoder(w).Encode("Course not fount with given id")

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting the course")
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == param["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted...")
			break
		}
	}
	json.NewEncoder(w).Encode("Course didn't found with the given id")

}
