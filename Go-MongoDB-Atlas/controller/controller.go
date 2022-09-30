package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/deepraj02/go_mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://maverick:Deepraj12.@cluster0.jlk5heg.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"
const colName = "watchlist"

// referencing to a mongodb Collection
var collection *mongo.Collection

// connect with mongodb
// init method only runs once
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Success")
	//creating Collections

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection Reference is Ready")

}

//mongo DB helpers -sep files
//insert 1 record

func insertOneMovie(movie model.Netfilx) { //-> as defined in the package name
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Movie with ID: ", inserted.InsertedID)
}

// UpdateOneRecord
func updateOneMethod(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count", result.ModifiedCount)
}

// delete 1 record
func deleteOneRecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie Deleted with ID :", result)
}

// delete all record
func deleteAllRecord() int64 {
	deletedCount, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of Movies deleted :", deletedCount.DeletedCount)
	return deletedCount.DeletedCount
}

func getAllMovies() {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cur.Next(context.Background()){
		var movie bson.M
		err:=cur.Decode(&movie)
		if err!=nil{
			log.Fatal(err)
		}
		movies=append(movies, movie)
	}
	defer cur.Close(context.Background())
}
