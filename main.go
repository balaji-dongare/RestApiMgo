package main

import (
	"fmt"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2"

	"github.com/RestApiMgo/controller"
	"github.com/RestApiMgo/employee"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func getData(session *mgo.Session) {
	// Collection People
	c := session.DB("RestAssignment").C("Employee")
	// Query One
	result := employee.Employee{}
	err := c.Find(bson.M{"fname": "Bala-G"}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Phone", result)
}
func main() {
	// //Connect to DB
	// var session = mongocon.Connect("localhost")
	// getData(session)
	//listening port
	var port = ":" + strconv.Itoa(3000)
	//init router
	router := mux.NewRouter()
	//enpoints
	router.HandleFunc("/getAllEmployee", controller.GetAllEmployees).Methods("GET")
	router.HandleFunc("/getEmployee/{fname}", controller.GetEmployee).Methods("GET")
	router.HandleFunc("/addNewEmployee", controller.InsertEmployee).Methods("POST")
	router.HandleFunc("/updateEmployee/{fname}", controller.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/deleteEmployee/{fname}", controller.DeleteEmployee).Methods("DELETE")
	fmt.Printf("application listening port%s\n", port)
	http.ListenAndServe(port, router)

}
