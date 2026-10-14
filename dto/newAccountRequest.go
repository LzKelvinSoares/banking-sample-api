package dto

import "github.com/LzKelvinSoares/lzk-go-lib/errs"

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (n NewAccountRequest) Validate() *errs.AppError {
	if n.Amount < 5000 {
		return errs.NewValidationError("To open new account, you need to deposit at least 5000.00")
	}

	if n.AccountType != "saving" && n.AccountType != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
