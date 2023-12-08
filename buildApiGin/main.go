package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

var courses []Course

// Middleware / helper - file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	r := gin.Default()

	courses = append(courses, Course{CourseId: "2", CourseName: "Java", CoursePrice: 254, Author: &Author{FullName: "Rohit varshney", Website: "github.com/rohit"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "JavaScript", CoursePrice: 199, Author: &Author{FullName: "Rajat sharma", Website: "github.com/rajat"}})
	r.GET("/ping", servHome)
	r.GET("/courses", getAllCourses)
	r.GET("/courses/:id", getOneCourse)
	r.POST("/courses", createOneCourse)
	r.PUT("/courses/:id", updateOneCourse)
	r.DELETE("/courses/:id", deleteOneCourse)
	r.DELETE("/courses", deleteAllCourse)
	r.Run()
}

func servHome(c *gin.Context) {
	c.String(http.StatusOK, "Hello!")
}
func getAllCourses(c *gin.Context) {
	fmt.Println("Get all courses")
	c.JSON(http.StatusOK, courses)
}

func getOneCourse(c *gin.Context) {
	fmt.Println("Get One course Request")

	// grab id from request

	params := c.Param("id")
	// loop through course, find matching id
	for _, course := range courses {
		if course.CourseId == params {
			c.JSON(http.StatusOK, course)
			return
		}
	}
	// json.NewEncoder(w).Encode("No course found with given id ")
	c.JSON(http.StatusBadRequest, "course not found")

}

func createOneCourse(c *gin.Context) {
	fmt.Println("create one course")

	// what if : body is empty

	// data like -{}

	var course Course
	c.BindJSON(&course)
	if course.IsEmpty() {
		c.JSON(http.StatusOK, "no data found")
		return
	}

	// TODO : check only if title is duplicate
	// loop, title, JSON

	for _, co := range courses {
		if co.CourseName == course.CourseName {
			c.JSON(http.StatusOK, "duplicate data found")
			return

		}
	}

	//generate a unique id, string
	//append course into courses

	rand.Seed(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	c.JSON(http.StatusOK, course)

}

func updateOneCourse(c *gin.Context) {
	fmt.Println("Update course")

	// grab id from request
	params := c.Param("id")

	// loop, id, remove, add with my id

	for index, course := range courses {
		if course.CourseId == params {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			c.BindJSON(&course)
			courses = append(courses, course)
			c.JSON(http.StatusOK, course)
			return
		}
	}
	// json.NewEncoder(w).Encode("No course found with given id ")
	c.JSON(http.StatusOK, "No course found with given id "+params)

}

func deleteOneCourse(c *gin.Context) {
	fmt.Println("Deleting one course")

	params := c.Param("id")

	// loop, id, remove

	for index, course := range courses {
		if course.CourseId == params {
			courses = append(courses[:index], courses[index+1:]...)
			c.JSON(http.StatusOK, "course deleted successfully")
			return

		}
	}
	// json.NewEncoder(w).Encode("No course found with this id")
	c.JSON(http.StatusOK, "No course found with given id "+params)

}

func deleteAllCourse(c *gin.Context) {
	fmt.Println("Deleting all courses")

	courses = nil
	c.JSON(http.StatusOK, "all courses deleted")

}
