package service

import (
	"banking/constants"
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"banking/logger"
	"time"
)

type TransactionService interface {
	Withdrawal(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
	Deposit(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repo        domain.TransactionRepository
	accountRepo domain.AccountRepository
}

func buildTransaction(req dto.NewTransactionRequest, transactionType string) domain.Transaction {
	return domain.Transaction{
		TransactionId:   "",
		AccountId:       req.AccountId,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
		TransactionType: transactionType,
		Amount:          req.Amount,
	}
}

func getOperator(transactionType string) string {
	operator := "+"
	if transactionType == constants.WITHDRAWAL {
		operator = "-"
	}
	return operator
}

func (s DefaultTransactionService) updateAccountBalance(newTransaction domain.Transaction) (*dto.NewTransactionResponse, *errs.AppError) {
	updatedAccount, err := s.accountRepo.UpdateBalance(newTransaction.AccountId, newTransaction.Amount, getOperator(newTransaction.TransactionType))
	if err != nil {
		return nil, err
	}
	response := newTransaction.ToNewTransactionResponseDTO()
	response.Balance = updatedAccount.Amount
	return &response, nil
}

func (s DefaultTransactionService) Withdrawal(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	account, appErr := s.accountRepo.GetById(req.AccountId)
	if appErr != nil {
		logger.Error("Error while retrieving account")
		return nil, appErr
	}

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if req.Amount > account.Amount {
		return nil, errs.NewValidationError("Insuficient funds")
	}

	t := buildTransaction(req, constants.WITHDRAWAL)
	newTransaction, err := s.repo.SaveTransaction(t)
	if err != nil {
		return nil, err
	}
	return s.updateAccountBalance(*newTransaction)
}

func (s DefaultTransactionService) Deposit(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	t := buildTransaction(req, constants.DEPOSIT)
	_, appErr := s.repo.SaveTransaction(t)
	if appErr != nil {
		return nil, appErr
	}
	return s.updateAccountBalance(t)
}

func NewTransactionService(repo domain.TransactionRepository, accountRepo domain.AccountRepository) DefaultTransactionService {
	return DefaultTransactionService{
		repo:        repo,
		accountRepo: accountRepo,
	}
}
