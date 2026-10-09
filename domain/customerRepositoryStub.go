package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Kelvin Luz", City: "João Pessoa", Zipcode: "58075000", DateofBirth: "30/07/1991", Status: "1"},
		{Id: "1002", Name: "Teste Luz", City: "João Pessoa", Zipcode: "58075000", DateofBirth: "08/03/1991", Status: "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
