package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"student_rest/db"
	"student_rest/routes"
	"student_rest/utils"
)

func main() {
	utils.LoadFixture("./repositories/testdata/register_course/register_course.sql")

	_db , _ := db.ConnectDB()
	r := routes.CreateRoutes(_db)
	log.Fatal(http.ListenAndServe(":8080", r))
}
