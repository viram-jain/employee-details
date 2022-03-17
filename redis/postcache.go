package redis

import "github.com/viram-jain/employee-details/model"

type PostCache interface {
	Set(key string, value model.Employee)
	Get(key string) *model.Employee
}
