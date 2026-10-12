package domain

import (
	"banking/errs"
	"banking/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDB struct {
	client *sqlx.DB
}

func (d TransactionRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	sqlInsert := "INSERT INTO transactions(account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		logger.Error("Error while creating new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted it: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	t.TransactionId = strconv.FormatInt(id, 10)
	return &t, nil
}

func NewTransactionRepositoryDB(dbClient *sqlx.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{
		client: dbClient,
	}
}
