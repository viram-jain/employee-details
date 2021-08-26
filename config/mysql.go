package config

import (
	"database/sql"
	"os"
	"rmqandredis/model"
)

var (
	err error
	db  *sql.DB
)

//ConnectToMysql creates a mysql connection
func ConnectToMysql() {
	db, err = sql.Open("mysql", os.Getenv("MYSQL_STRING"))
	if err != nil {
		sugarLogger.Errorf("Failed to connect to db %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		sugarLogger.Errorf("Failed to ping to db %s", err.Error())
	}
}

//GetDetailsFromMysql retrieves details from mysql database
func GetDetailsFromMysql(id string) model.Employee {
	var employee model.Employee
	err = db.QueryRow("select name, age, manager,position from employee where id = ?", id).Scan(&employee.Name, &employee.Age, &employee.Manager, &employee.Position)
	if err != nil {
		sugarLogger.Errorf("Failed to fetch details %s", err.Error())
	}
	return employee
}
