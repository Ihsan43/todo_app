package routes

import (
	"fmt"
	"net/http"
	"todo_app/controller"
	"todo_app/db"
	"todo_app/repository"
	"todo_app/usecase"

	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {

	db := db.InitDB()

	todoRep := repository.NewTodoRepository(db)
	todoUseCase := usecase.NewTodoUseCase(todoRep)
	todoController := controller.NewTodoController(todoUseCase)

	router := mux.NewRouter()

	router.HandleFunc("/todo", todoController.AddTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", todoController.GetTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", todoController.UpdatedTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", todoController.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todos", todoController.GetAllTodos).Methods("GET")

	fmt.Println("Running in server:8080")
	http.ListenAndServe(":8080", router)
}
