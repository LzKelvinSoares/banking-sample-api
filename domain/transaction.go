package domain

import (
	"banking/dto"
	"banking/errs"
)

type Transaction struct {
	TransactionId   string
	AccountId       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

type TransactionRepository interface {
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}

func (t Transaction) ToNewTransactionResponseDTO() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId: t.TransactionId,
		Balance:       0,
	}
}
