package employee

import (
	"time"
)

//Employee Struct
type Employee struct {
	//ID          bson.ObjectId `bson:"_id,omitempty"`
	FName       string    `json:"fname"`
	LName       string    `json:"lname"`
	Designation string    `json:"designation"`
	Phone       int64     `json:"phone"`
	JoiningDate time.Time `json:"joiningdate"`
}
