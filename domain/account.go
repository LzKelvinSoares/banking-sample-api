package domain

import (
	"banking/dto"
	"time"

	"github.com/LzKelvinSoares/lzk-go-lib/errs"
)

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain banking/domain AccountRepository
type AccountRepository interface {
	GetById(id string) (*Account, *errs.AppError)
	Save(Account) (*Account, *errs.AppError)
	UpdateBalance(id string, amount float64, transactionOperator string) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDTO() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

func NewAccount(customerId, accountType string, amount float64) Account {
	return Account{
		CustomerId:  customerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}
