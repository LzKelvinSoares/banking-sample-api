package domain

import (
	"banking/dto"
	"banking/errs"
)

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	UpdateBalance(id string, amount float64, transactionOperator string) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}
