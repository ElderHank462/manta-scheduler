/*
Defines the task type.

*/

package models

type Task struct {
	Name		string	`json:"name"`
	Description	string	`json:"description"`
	Duration	int	`json:"duration"`
	DueDate		string	`json:"due_date"`
	Group		string	`json:"group"`
}
