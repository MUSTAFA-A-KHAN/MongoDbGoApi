package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"mongoapi/model"
)

const connectionString = "mongodb+srv://Mkhan62608gmailcom:pass%40123@cluster0.zuzzadg.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "theatre"
const colName = "watchlist"

var collection *mongo.Collection

//Connect with mongo db

func init() {
	//Client options
	clientOption := options.Client().ApplyURI(connectionString)

	client, error := mongo.Connect(context.TODO(), clientOption)

	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("mongo connected successfully")

	collection = client.Database(dbName).Collection(colName)

	//collection is ready
	fmt.Println("Collection instance is ready")
}

//Mongo DB helpers -file

// insert one record
func insertOneMovie(movie model.Theatre) {
	inserted, error := collection.InsertOne(context.Background(), movie)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("inserted onw movie with id:", inserted.InsertedID)
}

// update 1 record
func updateonemovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified value:", result.ModifiedCount)
}

// delete one record
func deleteOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex((movieID))
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted:", deleteCount)
}

// delete all from movie db
func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("no. of movies deleted", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from DB
func getAllMovie() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for curr.Next(context.Background()) {
		var movie bson.M
		err := curr.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer curr.Close(context.Background())

	return movies
}

// getOne
func getOneMovie(movieID string) bson.Raw {
	id, err := primitive.ObjectIDFromHex((movieID))
	filter := bson.M{"_id": id}
	movie, err := collection.FindOne(context.Background(), filter).Raw()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie:", movie)
	return movie
}

// actual controlller

func GetMyAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovie()
	json.NewEncoder(w).Encode(allMovies)

}
func GetMyMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Allow-Control-Allow-Methods", "GET")
	fmt.Println("inside GETmovie")
	params := mux.Vars(r)
	fmt.Println(params["id"])
	movie := getOneMovie(params["id"])
	json.NewEncoder(w).Encode(movie.String())
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Theatre
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateonemovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	// params := mux.Vars(r)
	// deleteOneMovie(params["id"])
	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
