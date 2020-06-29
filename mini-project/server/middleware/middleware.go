package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IndexPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("MyResponse"))
}
func GetAllTasks(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("All Tasks"))
}

func AddTask(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Addd"))
}

func DeleteTask(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Delete"))
}
func EditTask(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Edit"))
}

const connectionString = ""
const databaseName = ""

var collection *mongo.Collection

const colelctionName = "todolist"

func init() {
	//Seeting the client options
	clientOptions := options.Client().ApplyURI(connectionString)
	//conenct to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Mongobd")
	collection = client.Database(databaseName).Collection(colelctionName)
}
