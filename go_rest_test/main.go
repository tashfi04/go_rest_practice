package main

import (
	"database/sql"
	"go_rest_test/config"
	"go_rest_test/routes"
	"net/http"
)

var Db *sql.DB

func main() {
	

	config.Init()

	r := routes.Routers()

	http.ListenAndServe(":8000", r)

}
