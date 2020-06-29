package router

import (
	"../middleware"
	"github.com/gorilla/mux"
)

//Router takes zero inputs and returns the router to main
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", middleware.IndexPage)
	router.HandleFunc("/add", middleware.AddTask)
	router.HandleFunc("/edit", middleware.EditTask)
	router.HandleFunc("/delete", middleware.DeleteTask)
	router.HandleFunc("/getAll", middleware.GetAllTasks)
	return router
}
