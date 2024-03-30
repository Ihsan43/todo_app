package controller

import (
	"encoding/json"
	"net/http"
	"todo_app/model"
	"todo_app/usecase"

	"github.com/gorilla/mux"
)

type todoController struct {
	todoUseCase usecase.TodoUseCase
}

func NewTodoController(todoUseCase usecase.TodoUseCase) *todoController {
	return &todoController{
		todoUseCase: todoUseCase,
	}
}

func (u *todoController) AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo model.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if todo.Name == "" || todo.Status == "" {
		http.Error(w, "Name dan status tidak boleh kosong", http.StatusBadRequest)
		return
	}

	addedTodo, err := u.todoUseCase.CreateTodo(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(map[string]any{
		"Messege": http.StatusCreated,
		"Data":    addedTodo,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func (c *todoController) GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	getTodo, err := c.todoUseCase.GetTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if getTodo == nil {
		http.Error(w, "id is not found", http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(map[string]any{
		"Messege": http.StatusFound,
		"Data":    getTodo,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func (c *todoController) UpdatedTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	var todo model.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTodo, err := c.todoUseCase.UpdateTodo(id, todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if updatedTodo == nil {
		http.Error(w, "id is not found", http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(map[string]interface{}{
		"Message": "Todo updated successfully",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	todo, err := c.todoUseCase.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if todo == nil {
		http.Error(w, "id is not found", http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(map[string]interface{}{
		"Message": "Todo Delete successfully",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func (c *todoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	todos, err := c.todoUseCase.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseJSON, err := json.Marshal(map[string]any{
		"Messege": http.StatusFound,
		"Data":    todos,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
