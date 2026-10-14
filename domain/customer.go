package domain

import (
	"banking/builder"
	"banking/dto"

	"github.com/LzKelvinSoares/lzk-go-lib/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) BuildCustomerDTO(b builder.CustomerBuilder) dto.CustomerResponse {
	customerResponse := b.WithId(c.Id).
		WithName(c.Name).
		WithCity(c.City).
		WithZipCode(c.Zipcode).
		WithDateOfBirth(c.DateofBirth).
		WithStatus(c.Status).
		CreateResponse()
	return customerResponse
}

type CustomerRepository interface {
	FindAll(status *Status) ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}
