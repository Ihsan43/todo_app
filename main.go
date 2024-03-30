package main

import (
	"todo_app/routes"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	routes.Setup(r)
}
