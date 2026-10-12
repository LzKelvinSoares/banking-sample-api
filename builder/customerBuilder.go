package builder

import (
	"banking/dto"
)

type CustomerBuilder struct {
	customerResponse dto.CustomerResponse
}

func (b *CustomerBuilder) WithName(name string) *CustomerBuilder {
	b.customerResponse.Name = name
	return b
}

func (b *CustomerBuilder) WithId(id string) *CustomerBuilder {
	b.customerResponse.Id = id
	return b
}

func (b *CustomerBuilder) WithCity(city string) *CustomerBuilder {
	b.customerResponse.City = city
	return b
}

func (b *CustomerBuilder) WithZipCode(zipcode string) *CustomerBuilder {
	b.customerResponse.Zipcode = zipcode
	return b
}

func (b *CustomerBuilder) WithDateOfBirth(dateOfBirth string) *CustomerBuilder {
	b.customerResponse.DateofBirth = dateOfBirth
	return b
}

func (b *CustomerBuilder) WithStatus(status string) *CustomerBuilder {
	statusToUse := "active"
	if status == "0" {
		statusToUse = "inactive"
	}
	b.customerResponse.Status = statusToUse
	return b
}

func (b *CustomerBuilder) CreateResponse() dto.CustomerResponse {
	return b.customerResponse
}

func NewCustomerBuilder() CustomerBuilder {
	return CustomerBuilder{
		customerResponse: dto.CustomerResponse{},
	}
}
