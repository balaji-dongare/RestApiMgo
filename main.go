package main

import (
	"fmt"
	"net/http"

	"github.com/TestGit/RestApiMgo/propparser"

	"github.com/RestApiMgo/controller"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//go:generate swagger generate spec -m -o ./swagger.json
func main() {
	config := propparser.ReadConfig()
	var port = ":" + config.Port
	//init router
	router := mux.NewRouter()
	corsObj := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/v1/Employees", controller.GetAllEmployees).Methods("GET", "OPTION")
	router.HandleFunc("/v1/Employee/{fname}", controller.GetEmployee).Methods("GET", "OPTION")
	router.HandleFunc("/v1/Employee", controller.InsertEmployee).Methods("POST", "OPTION")
	router.HandleFunc("/v1//Employee/{fname}", controller.UpdateEmployee).Methods("PUT", "OPTION")
	router.HandleFunc("/v1/Employee/{fname}", controller.DeleteEmployee).Methods("DELETE", "OPTION")

	fmt.Printf("application listening port%s\n", port)
	http.ListenAndServe(port, handlers.CORS(corsObj)(router))

}
