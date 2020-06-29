package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"../models"
)

func IndexPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("MyResponse"))
}
func GetAllTasks(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("All Tasks"))
}

func AddTask(res http.ResponseWriter, req *http.Request) {
	var task models.ToDoList
	_ = json.NewDecoder(req.Body).Decode(&task)
	fmt.Println(task)
	_, err := insertTask(task)
	if err != nil {
		res.Write([]byte("Error"))

	} else {
		json.NewEncoder(res).Encode(task)
	}

	//res.Write([]byte("Addd"))
}

func DeleteTask(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Delete"))
}
func EditTask(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Edit"))
}

const connectionString = "mongodb://go_user:gouser2020@ds159185.mlab.com"
const databaseName = "go_db"

var collection *mongo.Collection

const colelctionName = "todolist"

func init() {
	//Seeting the client options
	//clientOptions := options.Client().ApplyURI(connectionString)
	//conenct to mongodb
	//client, err := mongo.Connect(context.TODO(), clientOptions)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Mongobd")
	collection = client.Database(databaseName).Collection(colelctionName)
}

func insertTask(task models.ToDoList) (insertedId string, errorResult error) {

	insertResult, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
		errorResult = err
	}

	fmt.Println("Inserted a record", insertResult.InsertedID)
	insertedId = "inserted successfully"
	return
}
