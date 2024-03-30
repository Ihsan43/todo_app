package repository

import (
	"database/sql"
	"time"
	"todo_app/model"
)

type TodoRepository interface {
	InsertTimeTable(todo model.Todo) (*model.Todo, error)
	GetTodoById(id string) (*model.Todo, error)
	UpdateTodo(id string, todo model.Todo) (*model.Todo, error)
	DeleteTodoById(id string) (model.Todo, error)
	GetAllTodos() ([]model.Todo, error)
}

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) InsertTimeTable(todo model.Todo) (*model.Todo, error) {
	createTimeTable := model.Todo{}

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	timeNow := time.Now()
	err = tx.QueryRow("INSERT INTO Todo (name, status, created_at, updated_at) VALUES ($1,$2,$3,$4) RETURNING id, name, status, created_at, updated_at",
		todo.Name, todo.Status, timeNow, todo.Updated_At).
		Scan(&createTimeTable.Id, &createTimeTable.Name, &createTimeTable.Status, &createTimeTable.Created_At, &createTimeTable.Updated_At)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &createTimeTable, nil
}

func (r *todoRepository) GetTodoById(id string) (*model.Todo, error) {

	var todo model.Todo

	err := r.db.QueryRow("Select id, name, status, created_at, updated_at FROM todo where id = $1", id).
		Scan(&todo.Id, &todo.Name, &todo.Status, &todo.Created_At, &todo.Updated_At)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *todoRepository) UpdateTodo(id string, todo model.Todo) (*model.Todo, error) {

	var updateTodo model.Todo

	// if todo.Id == "" {
	// 	return nil,errors.New("id cannot be empty")
	// }

	err := r.db.QueryRow("UPDATE todo SET name = $1, status = $2, updated_at = $3 WHERE id = $4 RETURNING id, name, status, updated_at",
		todo.Name, todo.Status, time.Now(), id).Scan(&todo.Id, &todo.Name, &todo.Status, &todo.Updated_At)

	if err != nil {
		return nil, err
	}

	return &updateTodo, nil
}

func (r *todoRepository) DeleteTodoById(id string) (model.Todo, error) {

	var todo model.Todo

	_, err := r.db.Exec("DELETE FROM todo WHERE id = $1", id)
	if err != nil {
		return model.Todo{}, err
	}

	return todo, nil
}

func (r *todoRepository) GetAllTodos() ([]model.Todo, error) {
	var todos []model.Todo

	rows, err := r.db.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.Id, &todo.Name, &todo.Status, &todo.Created_At, &todo.Updated_At)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
