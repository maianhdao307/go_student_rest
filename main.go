package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"student_rest/db"
	"student_rest/routes"
)

func main() {
	_db , _ := db.ConnectDB()
	r := routes.CreateRoutes(_db)
	log.Fatal(http.ListenAndServe(":8080", r))
}
