package controller

import (
	"encoding/json"
	"net/http"

	"github.com/RestApiMgo/employee"
	"github.com/RestApiMgo/mongocon"
	"github.com/TestGit/RestApiMgo/propparser"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/mgo.v2"
)

//Connect to DB
var config = propparser.ReadConfig()
var session = mongocon.Connect(config.ConnectionString)

// Index
var index = mgo.Index{
	Key:        []string{"fname", "lname"},
	Unique:     true,
	DropDups:   true,
	Background: true,
	Sparse:     true,
}

// Collection Employee
var c = session.DB(config.Database).C(config.Collection)

var err = c.EnsureIndex(index)

//GetAllEmployees fetch all emp fronm db
func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	var results []employee.Employee
	w.Header().Set("Content-Type", "application/json")
	//Query all
	err := c.Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

//GetEmployee fetch emp from db
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Query One on the basis of fname
	params := mux.Vars(r)
	fname := params["fname"]
	result := employee.Employee{}
	err := c.Find(bson.M{"fname": fname}).One(&result)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

//InsertEmployee in DB
func InsertEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp employee.Employee

	//Encode request body in Emp Struct
	_ = json.NewDecoder(r.Body).Decode(&emp)
	// Insert Query
	err := c.Insert(emp)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}

//UpdateEmployee fetch all emp fronm db
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Query One on the basis of fname
	params := mux.Vars(r)
	//path param
	fname := params["fname"]
	var emp employee.Employee
	//Encode request body in Emp Struct //updateded  Values
	_ = json.NewDecoder(r.Body).Decode(&emp)
	change := bson.M{"$set": emp}
	err := c.Update(bson.M{"fname": fname}, change)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(emp)
}

//DeleteEmployee fetch all emp fronm db
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Query One on the basis of fname
	params := mux.Vars(r)
	fname := params["fname"]
	err := c.Remove(bson.M{"fname": fname})
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Delete Successfully!")
}
