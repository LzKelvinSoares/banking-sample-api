package domain

import (
	"database/sql"
	"strconv"

	"github.com/LzKelvinSoares/lzk-go-lib/logger"

	"github.com/LzKelvinSoares/lzk-go-lib/errs"

	"github.com/jmoiron/sqlx"
)

type AccountRepositorySQLServerDB struct {
	client *sqlx.DB
}

func (d AccountRepositorySQLServerDB) GetById(id string) (*Account, *errs.AppError) {
	findByIdSql := `SELECT 
						account_id, 
						customer_id, 
						opening_date, 
						account_type, 
						amount, 
						status 
					FROM 
						accounts 
					WHERE 
						account_id = @p1`
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

func (d AccountRepositorySQLServerDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := `INSERT INTO 
					accounts(
						customer_id, 
						opening_date, 
						account_type, 
						amount, 
						status
					) VALUES (
					 	@p1,
						@p2, 
						@p3, 
						@p4, 
						@p5
					);
					SELECT CONVERT(bigint, SCOPE_IDENTITY());`
	var id int64
	err := d.client.QueryRow(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status).Scan(&id)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositorySQLServerDB) UpdateBalance(id string, amount float64, transactionOperator string) (*Account, *errs.AppError) {
	sqlUpdate := "UPDATE accounts SET amount = (amount + @p1) WHERE account_id = @p2"
	if transactionOperator == "-" {
		sqlUpdate = "UPDATE accounts SET amount = (amount - @p1) WHERE account_id = @p2"
	}
	_, err := d.client.Exec(sqlUpdate, amount, id)
	if err != nil {
		logger.Error("Error while updating account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	account, appErr := d.GetById(id)
	if appErr != nil {
		logger.Error("Error while updating account")
		return nil, appErr
	}
	return account, nil
}

func NewAccountRepositorySQLServerDB(dbClient *sqlx.DB) AccountRepositorySQLServerDB {
	return AccountRepositorySQLServerDB{
		client: dbClient,
	}
}
