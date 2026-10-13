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

func (n NewTransactionRequest) IsTransactionTypeWithdrawal() bool {
	return n.TransactionType == constants.WITHDRAWAL
}

func (n NewTransactionRequest) IsTransactionTypeDeposit() bool {
	return n.TransactionType == constants.DEPOSIT
}

func (n NewTransactionRequest) Validate() *errs.AppError {
	if !n.IsTransactionTypeWithdrawal() && !n.IsTransactionTypeDeposit() {
		return errs.NewValidationError("Transaction type should be Withdrawal or Deposit")
	}

	if n.Amount < 0 {
		return errs.NewValidationError("The new transaction cannot be under zero")
	}
	return nil
}
