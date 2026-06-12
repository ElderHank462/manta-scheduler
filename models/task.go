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
	Group		int	`json:"group"`
}
