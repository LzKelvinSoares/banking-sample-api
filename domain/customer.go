package domain

import "banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll(status *Status) ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}
