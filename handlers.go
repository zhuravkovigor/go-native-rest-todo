package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// GetTodos godoc
// @Summary Get all tasks
// @Tags todos
// @Produce json
// @Success 200 {array} Todo
// @Router /todos [get]
func GetTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, done FROM todos")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var t Todo

		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		todos = append(todos, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// CreateTodo godoc
// @Summary Create a new task
// @Tags todos
// @Produce json
// @Param todo body TodoUpdateRequest true "New task"
// @Success 200 {object} Todo
// @Router /todos [post]
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var t TodoUpdateRequest

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var created Todo

	err := db.QueryRow("INSERT INTO todos(id, title, done) VALUES(DEFAULT, $1, $2) RETURNING id, title, done", t.Title, t.Done).Scan(&created.ID, &created.Title, &created.Done)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(created)
}

// UpdateTodoName godoc
// @Summary Update a task name by id
// @Tags todos
// @Produce json
// @Param id path string true "Task ID"
// @Param todo body TodoUpdateNameRequest true "Updated task"
// @Success 200 {object} Todo
// @Router /todos/{id}/rename [put]
func UpdateTodoName(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(strings.TrimPrefix(r.URL.Path, "/todos/"), "/")[0]
	var t TodoUpdateNameRequest

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var updated Todo

	err := db.QueryRow("UPDATE todos SET title = $1 WHERE id = $2 RETURNING id, title, done", t.Title, id).Scan(&updated.ID, &updated.Title, &updated.Done)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(updated)
}

// UpdateTodoDone godoc
// @Summary Update a task done status by id
// @Tags todos
// @Produce json
// @Param id path string true "Task ID"
// @Param todo body TodoUpdateDoneRequest true "Updated task"
// @Success 200 {object} Todo
// @Router /todos/{id}/done [put]
func UpdateTodoDone(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(strings.TrimPrefix(r.URL.Path, "/todos/"), "/")[0]
	var t TodoUpdateDoneRequest

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var updated Todo

	err := db.QueryRow("UPDATE todos SET done = $1 WHERE id = $2 RETURNING id, title, done", t.Done, id).Scan(&updated.ID, &updated.Title, &updated.Done)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(updated)
}

// DeleteTodo godoc
// @Summary Delete a task by id
// @Tags todos
// @Produce json
// @Param id path string true "Task ID"
// @Success 204
// @Router /todos/{id} [delete]
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")

	_, err := db.Exec("DELETE FROM todos WHERE id=$1", id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
