package employee

import (
	"time"
)

// Employee employeestruct
// swagger:model
type Employee struct {
	FName       string    `json:"fname"`
	LName       string    `json:"lname"`
	Designation string    `json:"designation"`
	Phone       int64     `json:"phone"`
	JoiningDate time.Time `json:"joiningdate"`
}
