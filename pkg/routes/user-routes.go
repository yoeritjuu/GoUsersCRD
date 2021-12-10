package routes

import (
	"github.com/gorilla/mux"
	"github.com/yoeritjuu/GoUsersCRD/pkg/controllers"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{bookId}", controllers.DeleteUser).Methods("DELETE")
}
