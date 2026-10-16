package domain

import (
	"strconv"

	"github.com/LzKelvinSoares/lzk-go-lib/logger"

	"github.com/LzKelvinSoares/lzk-go-lib/errs"

	"github.com/jmoiron/sqlx"
)

type TransactionRepositorySQLServerDB struct {
	client *sqlx.DB
}

func (d TransactionRepositorySQLServerDB) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	sqlInsert := `INSERT INTO 
					transactions(
						account_id, 
						amount, 
						transaction_type, 
						transaction_date
					) VALUES (
					 	@p1, 
						@p2, 
						@p3, 
						@p4
					);
					SELECT CONVERT(bigint, SCOPE_IDENTITY());`
	var id int64
	err := d.client.QueryRow(sqlInsert, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate).Scan(&id)
	if err != nil {
		logger.Error("Error while creating new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	t.TransactionId = strconv.FormatInt(id, 10)
	return &t, nil
}

func NewTransactionRepositorySQLServerDB(dbClient *sqlx.DB) TransactionRepositorySQLServerDB {
	return TransactionRepositorySQLServerDB{
		client: dbClient,
	}
}
