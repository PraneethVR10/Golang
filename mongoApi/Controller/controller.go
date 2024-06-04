package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models "github.com/PraneethVR10/MongoApi/Model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://praneethvr:<praneeth2001>@cluster0.i6ve5io.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "hotstar"
const colName = "watchlist"

//Most important 

var collection *mongo.Collection

//connect with mongodb

func init() {
	//client option


	clientOption:= options.Client().ApplyURI(connectionString)

	//connect to Mongodb

	client, err := mongo.Connect(context.TODO(), clientOption)

	if client != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection is a success")

	collection = client.Database(dbName).Collection(colName)

	//collection reference

	fmt.Println("Collection reference is ready")
} 

//MONGODB helpers 

//insert 1 record 

func insertOneMovie(movie models.Hotstar){
	inserted , err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in DB with ID:", inserted.InsertedID)
}

func insertOneSeries(series models.Hotstar) {
	inserted1, err := collection.InsertOne(context.Background(), series)

	if err != nil {
	   log.Fatal(err)
	}

	fmt.Println("Inserted 1 series in DB with ID:", inserted1.InsertedID)
}

// update 1 record 

func updateOneMovie(movieId string) {
  id , _ := primitive.ObjectIDFromHex(movieId)
  filter := bson.M{"_id": id}
  update := bson.M{"$set":bson.M{"watched": true}}

  result , err:= collection.UpdateOne(context.Background(), filter, update)

  if err != nil {
	log.Fatal(err)
  }

  fmt.Println("modified count:" , result.ModifiedCount)
}

//delete 1 record

func deleteOneMovie(movieId string) {
	id,_ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}


  delete , err:= collection.DeleteOne(context.Background(), filter,)

  if err != nil {
	log.Fatal(err)
  }

  fmt.Println("movie got deleted with delete count:" , delete)
}

//delete all records from the database

func deleteAllMovie() int64 {
	deletResult, err  := collection.DeleteMany(context.Background(), bson.D{}, nil) 

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies delete:", deletResult.DeletedCount)
	return deletResult.DeletedCount

}

func getAllMovies() []primitive.M{
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie) 
		if err != nil {
			log.Fatal(err)

		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())

	return movies
}

//Actual Controller - file 

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) { 

	w.Header().Set("content-Type"," application/x-www-form-urlencode")

	allMovies := getAllMovies()
     json.NewEncoder(w).Encode(allMovies)


    
}





     

