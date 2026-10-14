package service

import (
	"banking/builder"
	"banking/domain"
	"banking/dto"

	"github.com/LzKelvinSoares/lzk-go-lib/errs"
)

//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package=service banking/service CustomerService
type CustomerService interface {
	GetAllCustomers(status *domain.Status) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo  domain.CustomerRepository
	build builder.CustomerBuilder
}

func (s DefaultCustomerService) GetAllCustomers(status *domain.Status) ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	var response []dto.CustomerResponse

	for _, c := range customers {
		response = append(response, c.BuildCustomerDTO(s.build))
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.BuildCustomerDTO(s.build)
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository, build builder.CustomerBuilder) DefaultCustomerService {
	return DefaultCustomerService{
		repo:  repository,
		build: build,
	}
}
