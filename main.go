package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

type Todo struct {
	ID string `json:"id,omitempty"`
	CreatedBy string `json:"author,omitempty"`
	Desc string `json:"desc,omitempty"`
}

var todos []Todo

func GetTodos(w http.ResponseWriter, req *http.Request){
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for _, item := range todos {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
}

func CreateTodo(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	var todo Todo
	_ = json.NewDecoder(req.Body).Decode(&todo)
	todo.ID = params["id"]
	todos = append(todos,todo)
	json.NewEncoder(w).Encode(todos)
}

func DeleteTodo(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for index, item := range todos {
		if item.ID == params["id"]{
			todos = append(todos[:index], todos[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

func main () {
	router := mux.NewRouter()
	todos = append(todos, Todo{ID: "1", CreatedBy: "Me", Desc: "Just a test"})
	todos = append(todos, Todo{ID: "2", CreatedBy: "You", Desc: "Just a second test"})	
	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", CreateTodo).Methods("POST")			
	router.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
