package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"simple/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://sriramrkpsv:sriramrkpsv@cluster0.qf9ypis.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	clientoption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientoption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB has created Sucessfully")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("collections are ready")
}

// to insert the value in the database
func adddata(movie model.Nexflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the data is inserted", inserted.InsertedID)
}

// to update the movies in the database

func updatedata(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the value is updated in the database", result.ModifiedCount)

}

// to delete one value in the database

func deleteOne(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteone, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the data has been deleted", deleteone)

}

// to get all the movies
func getall() []primitive.M {
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

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	allmovies := getall()
	json.NewEncoder(w).Encode(allmovies)

}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Controller-Allow-Method", "POST")

	var movie model.Nexflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	adddata(movie)

}

func MarkAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Controller-Allow-Method", "POST")

	params := mux.Vars(r)
	updatedata(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Controller-Allow-Method", "POST")

	params := mux.Vars(r)

	deleteOne(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
