package main

import (
	"encoding/json"	
	"net/http"

	"github.com/gorilla/mux"
)

func SetHeader(w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetTodos(w http.ResponseWriter, req *http.Request){
	SetHeader(w)
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, req *http.Request){
	SetHeader(w)		
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
	SetHeader(w)		
	params := mux.Vars(req)
	var todo Todo
	_ = json.NewDecoder(req.Body).Decode(&todo)
	todo.ID = params["id"]
	todos = append(todos,todo)
	json.NewEncoder(w).Encode(todos)
}

func DeleteTodo(w http.ResponseWriter, req *http.Request){
	SetHeader(w)		
	params := mux.Vars(req)
	for index, item := range todos {
		if item.ID == params["id"]{
			todos = append(todos[:index], todos[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}
