package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"student_rest/db"
	"student_rest/routes"
)

func main() {
	db.ConnectDB()
	r := routes.CreateRoutes()
	log.Fatal(http.ListenAndServe(":8080", r))
}
