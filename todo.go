package main

type Todo struct {
	ID string `json:"id,omitempty"`
	CreatedBy string `json:"author,omitempty"`
	Desc string `json:"desc,omitempty"`
	Completed bool `json:"completed,omitempty"`
}

var todos []Todo
