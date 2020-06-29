package middleware

import (
	"net/http"
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


