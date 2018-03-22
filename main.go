package main

import (
	"net/http"
	"log"
)

func example () {
	todos = append(todos, Todo{ID: "1", CreatedBy: "Me", Desc: "Just a test", Completed: false})
	todos = append(todos, Todo{ID: "2", CreatedBy: "You", Desc: "Just a second test", Completed: true})	
}

func main () {
	router := NewRouter()
	example()
	log.Fatal(http.ListenAndServe(":8080", router))
}
