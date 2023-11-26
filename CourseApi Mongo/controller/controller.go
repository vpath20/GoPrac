package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohit/courseApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoString = "mongodb+srv://rohitvarshney87:Rohit1234@cluster0.crt6ax0.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Courses"
const colName = "Available Courses"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(mongoString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection Reference is ready...")
}

// Controller Helper Functions
func getOneCourse(courseId string) []primitive.M {
	id, _ := primitive.ObjectIDFromHex(courseId)
	filter := bson.M{"_id": id}
	var courses []primitive.M

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	if cursor.Next(context.Background()) {
		var course bson.M
		err := cursor.Decode(&course)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Successfully fetched course")
		courses = append(courses, course)
		return courses
	}
	fmt.Println("Course not found")
	return courses

}
func updateOneCourse(courseId string, course model.Course) bool {
	id, _ := primitive.ObjectIDFromHex(courseId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": course}
	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal("this", err)
	}
	if updated.ModifiedCount > 0 {
		fmt.Println("updated successful", updated.ModifiedCount)
		return true
	}
	fmt.Println("updated Unsuccessful", updated.ModifiedCount)
	return false

}

func insertOneCourse(course model.Course) bool {
	filter := bson.M{"author.fullname": course.Author.FullName, "coursename": course.CourseName}
	cursor, err := collection.Find(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	if !cursor.Next(context.Background()) {
		inserted, err := collection.InsertOne(context.Background(), course)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Course inserted successfully with course id: ", inserted.InsertedID)
		return true
	}
	fmt.Println("Can't insert course duplicate found")
	return false

}

func getAllCourse() []primitive.M {
	filter := bson.M{}
	var courses []primitive.M

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(cursor)
	}

	for cursor.Next(context.Background()) {
		var course bson.M
		err := cursor.Decode(&course)
		if err != nil {
			log.Fatal(err)
		}
		courses = append(courses, course)
	}

	return courses
}

func deleteOneCourse(courseId string) bool {
	id, _ := primitive.ObjectIDFromHex(courseId)
	filter := bson.M{"_id": id}

	deleted, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	if deleted.DeletedCount == 0 {
		fmt.Println("delete unsuccessful")
		return false
	}
	fmt.Println("Deleted successfully ", deleted.DeletedCount)
	return true

}

func deleteAllCourse() {
	filter := bson.M{}
	deleted, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All courses deleted", deleted.DeletedCount)
}

// Actual Controllers
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	courses := getAllCourse()
	json.NewEncoder(w).Encode(courses)

}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	course := getOneCourse(params["id"])
	json.NewEncoder(w).Encode(course)

}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var course model.Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if insertOneCourse(course) {
		json.NewEncoder(w).Encode(course)
		return
	}
	json.NewEncoder(w).Encode("duplicate course found")

}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	var course model.Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if updateOneCourse(params["id"], course) {
		json.NewEncoder(w).Encode(course)
		return
	}
	json.NewEncoder(w).Encode("Course Update Unsuccessful")

}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	if deleteOneCourse(params["id"]) {
		json.NewEncoder(w).Encode("course deleted")
		return
	}
	json.NewEncoder(w).Encode("course not found")

}

func DeleteAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAllCourse()
	json.NewEncoder(w).Encode("course deleted")

}
