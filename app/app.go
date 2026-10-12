package app

import (
	"banking/builder"
	"banking/domain"
	"banking/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variables not defined!")
	}
}

func Start() {
	sanityCheck()

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandler{
	// 	service: service.NewCustomerService(domain.NewCustomerRepositoryStub()),
	// }
	ch := CustomerHandler{
		service:       service.NewCustomerService(domain.NewCustomerRepositoryDB(), builder.NewCustomerBuilder()),
		statusService: service.NewStatusService(),
	}

	// define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}
