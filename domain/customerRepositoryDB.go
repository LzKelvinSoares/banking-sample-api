package domain

import (
	"banking/errs"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll(status *Status) ([]Customer, *errs.AppError) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	var rows *sql.Rows
	var err error
	if status != nil {
		findAllSql += " where status = ?"
		rows, err = d.client.Query(findAllSql, status)
	} else {
		rows, err = d.client.Query(findAllSql)
	}

	if err != nil {
		log.Println("Error while scanning customer: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	}
	defer rows.Close()

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customer: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected DB Error")
		}
		customers = append(customers, c)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error while scanning customer: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	findByIdSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(findByIdSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not Found")
		} else {
			log.Println("Error while scanning customer: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected DB Error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{
		client: client,
	}
}
