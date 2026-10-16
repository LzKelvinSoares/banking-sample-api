package domain

import (
	"database/sql"

	"github.com/LzKelvinSoares/lzk-go-lib/logger"

	"github.com/LzKelvinSoares/lzk-go-lib/errs"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"
)

type CustomerRepositorySQLServerDB struct {
	client *sqlx.DB
}

func (d CustomerRepositorySQLServerDB) FindAll(status *Status) ([]Customer, *errs.AppError) {
	findAllSql := `SELECT 
					customer_id, 
					customername as name, 
					city, 
					zipcode, 
					date_of_birth, 
					status 
				FROM 
					customers`
	var err error
	customers := make([]Customer, 0)
	if status != nil {
		findAllSql += " where status = @p1"
		err = d.client.Select(&customers, findAllSql, status)
	} else {
		err = d.client.Select(&customers, findAllSql)
	}

	if err != nil {
		logger.Error("Error while scanning customer: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	}

	return customers, nil
}

func (d CustomerRepositorySQLServerDB) ById(id string) (*Customer, *errs.AppError) {
	findByIdSql := `SELECT 
						customer_id,
						customername as name, 
						city, 
						zipcode, 
						date_of_birth, 
						status 
					FROM 
						customers 
					WHERE 
						customer_id = @p1`

	// row := d.client.QueryRow(findByIdSql, id)
	var c Customer
	err := d.client.Get(&c, findByIdSql, id)
	// err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not Found")
		} else {
			logger.Error("Error while scanning customer: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected DB Error")
		}
	}

	return &c, nil
}

func NewCustomerRepositorySQLServerDB(dbClient *sqlx.DB) CustomerRepositorySQLServerDB {
	return CustomerRepositorySQLServerDB{
		client: dbClient,
	}
}
