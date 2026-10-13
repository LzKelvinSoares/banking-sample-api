package constants

const (
	GET_CUSTOMERS      = "/customers"
	GET_CUSTOMER_BY_ID = "/customers/{id:[0-9]+}"
	NEW_ACCOUNT        = "/customers/{id:[0-9]+}/account"
	MAKE_TRANSACTION   = "/customers/{id:[0-9]+}/account/{account_id:[0-9]+}/{transaction_type}"
)

const (
	ROUTE_NAME_GET_CUSTOMERS    = "GetAllCustomers"
	ROUTE_NAME_GET_CUSTOMER     = "GetCustomer"
	ROUTE_NAME_NEW_ACCOUNT      = "NewAccount"
	ROUTE_NAME_MAKE_TRANSACTION = "MakeTransaction"
)
