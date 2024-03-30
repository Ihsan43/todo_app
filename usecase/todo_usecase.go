package usecase

import (
	"strconv"
	"todo_app/model"
	"todo_app/repository"
)

type TodoUseCase interface {
	CreateTodo(todo model.Todo) (*model.Todo, error)
	GetTodo(id string) (*model.Todo, error)
	UpdateTodo(id string, todo model.Todo) (*model.Todo, error)
	GetTodos() ([]model.Todo, error)
	Delete(id string) (*model.Todo, error)
}

type todoUseCase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUseCase(todoRepo repository.TodoRepository) TodoUseCase {
	return &todoUseCase{
		todoRepo: todoRepo,
	}
}

func (u *todoUseCase) CreateTodo(todo model.Todo) (*model.Todo, error) {
	return u.todoRepo.InsertTimeTable(todo)
}

func (u *todoUseCase) GetTodo(id string) (*model.Todo, error) {
	return u.todoRepo.GetTodoById(id)
}

func (u *todoUseCase) UpdateTodo(id string, payload model.Todo) (*model.Todo, error) {

	// var todo *model.Todo
	// var err error

	// todo, err = u.GetTodo(id)
	// if err != nil {
	// 	return nil, err
	// }

	// todo, err = u.UpdateTodo(id, payload)
	// if err != nil {
	// 	return nil, err
	// }

	// newTodo := &model.Todo{
	// 	Id:         todo.Id,
	// 	Name:       todo.Name,
	// 	Status:     todo.Status,
	// 	Created_At: todo.Created_At,
	// 	Updated_At: todo.Updated_At,
	// }

	return u.todoRepo.UpdateTodo(id, payload)
}

func (u *todoUseCase) Delete(id string) (*model.Todo, error) {
	// Hapus entitas dengan ID yang ditentukan
	todo, err := u.todoRepo.DeleteTodoById(id)
	if err != nil {
		return nil, err
	}

	// Perbarui urutan ID untuk entitas yang tersisa setelah penghapusan
	todos, err := u.todoRepo.GetAllTodos()
	if err != nil {
		return nil, err
	}
	for i := range todos {
		todos[i].Id = strconv.Itoa(i + 1)
		if _, err := u.todoRepo.UpdateTodo(todos[i].Id, todos[i]); err != nil {
			return nil, err
		}
	}

	return &todo, nil
}

func (u *todoUseCase) GetTodos() ([]model.Todo, error) {
	return u.todoRepo.GetAllTodos()
}
