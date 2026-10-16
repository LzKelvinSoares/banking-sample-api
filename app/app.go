package app

import (
	"banking/builder"
	"banking/constants"
	"banking/domain"
	"banking/service"
	"log"
	"net/http"

	libConstants "github.com/LzKelvinSoares/lzk-go-lib/constants"
	"github.com/LzKelvinSoares/lzk-go-lib/environment"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"
)

func Start() {
	serverVariables := environment.GetServerVariables()
	router := mux.NewRouter()
	dbType := libConstants.SQLServer
	dbClient := environment.GetDBClient(dbType)

	accountRepositoryDB, customerRepositoryDB, transactionRepositoryDB := GetDBRepositories(dbType, dbClient)

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

func GetDBRepositories(dbType libConstants.DBType, dbClient *sqlx.DB) (
	domain.AccountRepository,
	domain.CustomerRepository,
	domain.TransactionRepository,
) {
	var customerRepositoryDB domain.CustomerRepository
	var accountRepositoryDB domain.AccountRepository
	var transactionRepositoryDB domain.TransactionRepository

	switch dbType {
	case libConstants.MySQL:
		customerRepositoryDB = domain.NewCustomerRepositoryDB(dbClient)
		accountRepositoryDB = domain.NewAccountRepositoryDB(dbClient)
		transactionRepositoryDB = domain.NewTransactionRepositoryDB(dbClient)
	case libConstants.SQLServer:
		customerRepositoryDB = domain.NewCustomerRepositorySQLServerDB(dbClient)
		accountRepositoryDB = domain.NewAccountRepositorySQLServerDB(dbClient)
		transactionRepositoryDB = domain.NewTransactionRepositorySQLServerDB(dbClient)
	}

	return accountRepositoryDB, customerRepositoryDB, transactionRepositoryDB
}
