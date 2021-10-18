package main

import (
	"encoding/json"

	"github.com/viramjainkaleyra/employee-details/config"
	"github.com/viramjainkaleyra/employee-details/model"
	"github.com/viramjainkaleyra/employee-details/publisher"
)

//main function
func main() {
	sugarLogger := config.InitLogger()
	employees := model.EmployeeJSON
	for _, employee := range employees {
		employeeJson, err := json.Marshal(employee)
		if err != nil {
			sugarLogger.Errorf("Cannot encode to JSON %s", err.Error())
		}
		publisher.Publisher([]byte(employeeJson))
	}
}
