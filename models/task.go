/*
Defines the task type.

*/

package models

import "time"

type Task struct {
	Name		string	`json:"name"`
	Description	string	`json:"description"`
	Duration	int	`json:"duration"`
	DueDate		time.Time	`json:"due_date"`
}

// type TaskGroup struct {
// 	Name	string	`json:"name"`
// 	Color	string	`json:"color"`
// 	Tasks	[]Task	`json:"tasks"`
// }