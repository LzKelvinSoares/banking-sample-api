package app

import (
	"banking/builder"
	"banking/constants"
	"banking/domain"
	"banking/service"
	"log"
	"net/http"
	"time"

	"github.com/LzKelvinSoares/lzk-go-lib/environment"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {
	serverVariables := environment.GetServerVariables()
	router := mux.NewRouter()
	dbClient := getDBClient()

	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)
	transactionRepositoryDB := domain.NewTransactionRepositoryDB(dbClient)
	customerBuilder := builder.NewCustomerBuilder()

	ch := CustomerHandler{
		service:       service.NewCustomerService(customerRepositoryDB, customerBuilder),
		statusService: service.NewStatusService(),
	}

	ah := AccountHandler{
		service:            service.NewAccountService(accountRepositoryDB),
		transactionService: service.NewTransactionService(transactionRepositoryDB, accountRepositoryDB),
	}

	// define routes
	router.
		HandleFunc(constants.GET_CUSTOMERS, ch.GetAllCustomers).
		Methods(http.MethodGet).
		Name(constants.ROUTE_NAME_GET_CUSTOMERS)

	router.
		HandleFunc(constants.GET_CUSTOMER_BY_ID, ch.GetCustomer).
		Methods(http.MethodGet).
		Name(constants.ROUTE_NAME_GET_CUSTOMER)

	router.
		HandleFunc(constants.NEW_ACCOUNT, ah.NewAccount).
		Methods(http.MethodPost).
		Name(constants.ROUTE_NAME_NEW_ACCOUNT)

	router.
		HandleFunc(constants.MAKE_TRANSACTION, ah.MakeTransaction).
		Methods(http.MethodPost).
		Name(constants.ROUTE_NAME_MAKE_TRANSACTION)

	am := AuthMiddleware{
		repo: domain.NewAuthRepository(),
	}
	router.
		Use(am.authorizationHandler())

	// starting server
	log.Fatal(http.ListenAndServe(serverVariables.GetAddressToServe(), router))

}

func getDBClient() *sqlx.DB {
	dbVariables := environment.GetDBVariables()
	client, err := sqlx.Open("mysql", dbVariables.GetDataSource())
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
