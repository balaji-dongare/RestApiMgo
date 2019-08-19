package main

import (
	"fmt"
	"net/http"

	"github.com/TestGit/RestApiMgo/propparser"

	"github.com/RestApiMgo/controller"
	"github.com/gorilla/mux"
)

func main() {
	config := propparser.ReadConfig()
	var port = ":" + config.Port
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
