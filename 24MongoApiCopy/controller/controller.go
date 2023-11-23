package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohit/controller/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// initialising constants

const mongoString = "mongodb+srv://rohitvarshney87:Rohit1234@cluster0.crt6ax0.mongodb.net/?retryWrites=true&w=majority"
const dbName = "PrimeVideos"
const colName = "WatchList"

var collection *mongo.Collection

// init function run only ones without even calling by main function
//Creating the MongoDB Connection

func init() {
	clientOption := options.Client().ApplyURI(mongoString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection reference is ready...")

}

// Helper functions for controller

func updateOneMovie(movieId string) bool {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	updated, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	if updated.ModifiedCount == 0 {
		fmt.Println("Update failed id not found")
		return false
	}
	fmt.Println("Updated count ", updated.ModifiedCount)
	return true

}

func deleteOneMovie(movieId string) bool {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleted, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	if deleted.DeletedCount == 0 {
		fmt.Println("deleted Unsuccessfully id not found")
		return false
	}
	fmt.Println("deleted successfully with delete count : ", deleted.DeletedCount)
	return true

}

func deleteAllMovie() {
	filter := bson.M{}
	collection.DeleteMany(context.Background(), filter)

}

func insertOneMovie(movie model.PrimeVideos) bool {

	filter := bson.M{"movie": movie.Movie}
	cursor, err1 := collection.Find(context.Background(), filter)
	if err1 != nil {
		log.Fatal(err1)
	}

	if !cursor.Next(context.Background()) {
		inserted, err := collection.InsertOne(context.Background(), movie)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Movie inserted Successfully with id : ", inserted.InsertedID)
		return true

	}

	fmt.Println("Movie inserted UnSuccessfully with id : ")
	return false

}

func getAllMovies() []primitive.M {
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return movies
}

// Actual Controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func InsertOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.PrimeVideos
	_ = json.NewDecoder(r.Body).Decode(&movie)

	if insertOneMovie(movie) {
		json.NewEncoder(w).Encode(movie)
		return
	}
	json.NewEncoder(w).Encode("Duplicate Movie found")

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	if updateOneMovie(params["id"]) {
		json.NewEncoder(w).Encode(params["id"])
		return
	}
	json.NewEncoder(w).Encode("id not found")

}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "Delete")

	params := mux.Vars(r)
	if deleteOneMovie(params["id"]) {
		json.NewEncoder(w).Encode(params["id"])
		return
	}
	json.NewEncoder(w).Encode("movie with given id not found")

}
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "Delete")

	deleteAllMovie()
	json.NewEncoder(w).Encode("All movies deleted...")

}
