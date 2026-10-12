package dto

import (
	"banking/constants"
	"banking/errs"
)

type NewTransactionRequest struct {
	AccountId       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

func (n NewTransactionRequest) Validate() *errs.AppError {
	if n.Amount < 0 {
		return errs.NewValidationError("The new transaction cannot be under zero")
	}

	if n.TransactionType != constants.WITHDRAWAL && n.TransactionType != constants.DEPOSIT {
		return errs.NewValidationError("Transaction type should be Withdrawal or Deposit")
	}
	return nil
}
