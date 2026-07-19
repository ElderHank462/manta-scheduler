package models

type TaskGroup struct {
	Name	string	`json:"name"`
	Color	string	`json:"color"`
	Tasks	[]Task	`json:"tasks"`
}