// Package controller Employee API.
//
// the purpose of this application is to provide an application
// that is using go code to define an  Rest API
//
//     Schemes: http, https
//     Host: localhost:3000
//     BasePath: /v1
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
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

// swagger:operation GET /Employees Employee getAllEmployee
//
// Get Employee
//
// Returns existing Employees
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: employee data
//     schema:
//      $ref: "#/definitions/Employee"
//   '405':
//     description: Method Not Allowed, likely url is not correct
//   '403':
//     description: Forbidden, you are not allowed to undertake this operation

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

// swagger:operation GET /Employee/{fname} Employee getEmployee
//
// Get Employee
//
// Returns existing Employee filtered by fname
//
// ---
// produces:
// - application/json
// parameters:
//  - name: fname
//    type: string
//    in: path
//    required: true
// responses:
//   '200':
//     description: employee data
//     schema:
//      "$ref": "#/definitions/Employee"
//   '405':
//     description: Method Not Allowed, likely url is not correct
//   '403':
//     description: Forbidden, you are not allowed to undertake this operation

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

// swagger:operation POST /Employee Employee addNewEmployee
//
// Add new Employee
//
// Returns new Employee
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: employee
//   in: body
//   description: add employee data
//   required: true
//   schema:
//     "$ref": "#/definitions/Employee"
// responses:
//   '200':
//     description: Employee response
//     schema:
//       "$ref": "#/definitions/Employee"
//   '409':
//     description: Conflict
//   '405':
//     description: Method Not Allowed, likely url is not correct
//   '403':
//     description: Forbidden, you are not allowed to undertake this operation

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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(emp)
}

// swagger:operation PUT /Employee/{fname} Employee updateEmployee
//
// Update Employee
//
// Update existing Employee filtered by fname
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: fname
//   type: string
//   in: path
//   required: true
// - name: employee
//   in: body
//   description: add employee data
//   required: true
//   schema:
//     "$ref": "#/definitions/Employee"
// responses:
//   '200':
//     description: Employee response
//     schema:
//       "$ref": "#/definitions/Employee"
//   '409':
//     description: Conflict
//   '405':
//     description: Method Not Allowed, likely url is not correct
//   '403':
//     description: Forbidden, you are not allowed to undertake this operation

//UpdateEmployee update employee
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

// swagger:operation DELETE /deleteEmployee/{fname} Employee deleteEmployee
//
// Delete Employee
//
// Delete existing Employee filtered by fname
//
// ---
// produces:
// - application/json
// parameters:
//  - name: fname
//    type: string
//    in: path
//    required: true
// responses:
//   '200':
//     description: delete employee sucessfully
//     schema:
//       "$ref": "#/definitions/Employee"
//   '405':
//     description: Method Not Allowed, likely url is not correct
//   '403':
//     description: Forbidden, you are not allowed to undertake this operation

//DeleteEmployee delete employee
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
	json.NewEncoder(w).Encode("")
}
