package main

// @Success 200 {array} Todo
type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TodoUpdateRequest struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TodoUpdateDoneRequest struct {
	Done bool `json:"done"`
}

type TodoUpdateNameRequest struct {
	Title string `json:"title"`
}
