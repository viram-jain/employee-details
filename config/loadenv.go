package config

import (
	"github.com/viram-jain/employee-details/constant"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var sugarLogger *zap.SugaredLogger

//Loadenv loads the environment variables
func Loadenv() {
	err := godotenv.Load(constant.Envfile)
	if err != nil {
		sugarLogger.Errorf("Failed to load the env file : Error = %s", err.Error())
		return
	}
}
