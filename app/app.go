package app

import (
	"banking/domain"
	"banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandler{
	// 	service: service.NewCustomerService(domain.NewCustomerRepositoryStub()),
	// }
	ch := CustomerHandler{
		service:       service.NewCustomerService(domain.NewCustomerRepositoryDB()),
		statusService: service.NewStatusService(),
	}

	// define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
