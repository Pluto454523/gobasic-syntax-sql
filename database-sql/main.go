package main

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Customer struct {
	CustomerID int
	Fname      string
	Lname      string
}

var db *sqlx.DB

func main() {

	var err error
	// db, err = sql.Open("mysql", "root:170745*@tcp(127.0.0.1:3306)/gobaisc_db")
	db, err = sqlx.Open("mysql", "root:170745*@tcp(localhost:3306)/gobaisc_db")
	if err != nil {
		panic(err)
	}

	// ** INSERT
	// infoCustomers := []Customer{
	// 	{
	// 		CustomerID: 1,
	// 		Fname:      "Nawin",
	// 		Lname:      "Khamchun",
	// 	},
	// 	{
	// 		CustomerID: 2,
	// 		Fname:      "Pannawat",
	// 		Lname:      "Imsin",
	// 	},
	// 	{
	// 		CustomerID: 3,
	// 		Fname:      "Thanapat",
	// 		Lname:      "Rattanasang",
	// 	},
	// }
	// for _, v := range infoCustomers {
	// 	err = AddCustomer(v)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("Insert %v %v success.\n", v.Fname, v.Lname)
	// }

	// ** UPDATE
	// infoCustomer := Customer{
	// 	CustomerID: 1,
	// 	Fname:      "Nawin Update",
	// }
	// err = UpdateCustomer(infoCustomer)
	// if err != nil {
	// 	panic(err)
	// }

	// ** DELETE
	// err = DeleteCustomer(3)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// ** SELECt ALL
	// customers, err := GetCustomerX() // TODO: customers is slice(pass by reference)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, v := range customers {
	// 	fmt.Printf("%#v\n", v)
	// }
	// newCustomers := customers // TODO: newCustomers recive address from customers
	// newCustomers[0] = Customer{
	// 	Fname: "Nobi",
	// 	Lname: "Nobita",
	// }
	// fmt.Println(customers[0])

	// ** SELECT BY ID
	// dataByID, err := GetCustomerXById(1) // TODO: dataByID is pointer(pass by reference)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%#v\n", dataByID)
	// newData := dataByID // TODO: newData recive address from dataByID
	// newData.Fname = "newData Pointer"
	// fmt.Printf("%v\n", dataByID)
}

func GetCustomerX() ([]Customer, error) {

	query := "SELECT fname,lname FROM customer"
	customers := []Customer{} //** slice สำหรับเก็บข้อมูล Customer ทั้งหมดที่ได้มาจาก Database
	err := db.Select(&customers, query)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func GetCustomerXById(id int) (*Customer, error) {

	query := "SELECT fname, lname FROM customer where CustomerID=?"
	customers := Customer{} //** slice สำหรับเก็บข้อมูล Customer ทั้งหมดที่ได้มาจาก Database
	err := db.Get(&customers, query, id)
	if err != nil {
		return nil, err
	}

	return &customers, nil
}

func GetCustomer() ([]Customer, error) {

	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "SELECT fname , lname FROM customer"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	customersData := []Customer{} //** slice สำหรับเก็บข้อมูล Customer ทั้งหมดที่ได้มาจาก Database
	for rows.Next() {
		customer := Customer{} //** ไว้สำหรับรับค่าจาก Database ที่ Scan มา
		err = rows.Scan(&customer.CustomerID, &customer.Fname, &customer.Lname)
		if err != nil {
			return nil, err
		}
		customersData = append(customersData, customer) //**
	}

	return customersData, nil
}

func GetCustomerById(id int) (*Customer, error) {

	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "SELECT CustomerID, fname, lname FROM customer where CustomerID=?"
	row := db.QueryRow(query, id)

	// query := "SELECT CustomerID, fname FROM customer where CustomerID=? and fname=?"
	// row := db.QueryRow(query, id, name)

	customer := Customer{}
	err = row.Scan(&customer.CustomerID, &customer.Fname, &customer.Lname)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func AddCustomer(customer Customer) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO customer(CustomerID, fname,lname) VALUES(?,?,?)"
	result, err := tx.Exec(query, customer.CustomerID, customer.Fname, customer.Lname)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affected <= 0 {
		return errors.New("Can not insert")
	}

	tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func UpdateCustomer(customer Customer) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "UPDATE customer set fname=? where CustomerID=?"
	result, err := tx.Exec(query, customer.Fname, customer.CustomerID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affected <= 0 {
		return errors.New("Can not update")
	}

	tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func DeleteCustomer(id int) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "delete from customer where CustomerID=?"
	result, err := tx.Exec(query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affected <= 0 {
		return errors.New("Can not delete")
	}

	tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
