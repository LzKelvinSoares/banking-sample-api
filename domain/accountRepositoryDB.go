package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) getById(id string) (*Account, *errs.AppError) {
	findByIdSql := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where account_id = ?"
	var a Account
	err := d.client.Get(&a, findByIdSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Account Not Found")
		} else {
			logger.Error("Error while scanning account: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected DB Error")
		}
	}

	return &a, nil
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts(customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted it: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDB) UpdateBalance(id string, amount float64, transactionOperator string) (*Account, *errs.AppError) {
	sqlUpdate := "UPDATE accounts SET amount = (amount + ?) WHERE account_id = ?"
	if transactionOperator == "-" {
		sqlUpdate = "UPDATE accounts SET amount = (amount - ?) WHERE account_id = ?"
	}
	_, err := d.client.Exec(sqlUpdate, amount, id)
	if err != nil {
		logger.Error("Error while updating account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	account, appErr := d.getById(id)
	if appErr != nil {
		logger.Error("Error while updating account")
		return nil, appErr
	}
	return account, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{
		client: dbClient,
	}
}
