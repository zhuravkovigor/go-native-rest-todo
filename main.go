package main

import (
	"fmt"
	_ "go-native-rest-todo/docs"
	"log"
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Todo REST API
// @version 1.0
// @description This is a simple todo API built with Go library net/http and PostgreSQL
// @host localhost:8080
// @BasePath /
func main() {
	InitDB()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	http.HandleFunc("GET /todos", GetTodos)
	http.HandleFunc("POST /todos", CreateTodo)
	http.HandleFunc("PUT /todos/{id}/rename", UpdateTodoName)
	http.HandleFunc("PUT /todos/{id}/done", UpdateTodoDone)
	http.HandleFunc("DELETE /todos/{id}", DeleteTodo)

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Printf("Server started on port http://localhost:%s", PORT)
	log.Printf("Swagger started on port http://localhost:%s/swagger/index.html", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}
