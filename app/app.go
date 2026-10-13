package app

import (
	"banking/builder"
	"banking/constants"
	"banking/domain"
	"banking/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
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
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

func getDBClient() *sqlx.DB {
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbAddress := os.Getenv("DB_ADDRESS")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")

	// dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)
	// client, err := sqlx.Open("mysql", dataSource)
	client, err := sqlx.Open("mysql", "root:password@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
