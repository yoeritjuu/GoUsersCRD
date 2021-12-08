package main

import (
	"log"
	"net/http"

	"github.com/Yoeritjuu/GoUserCRD/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yoeritjuu/GoUsersCRD/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
