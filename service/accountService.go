package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	a := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)

	if newAccount, err := s.repo.Save(a); err != nil {
		return nil, err
	} else {
		return newAccount.ToNewAccountResponseDTO(), nil
	}
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{
		repo: repo,
	}
}
