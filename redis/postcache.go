package redis

import "github.com/viramjainkaleyra/employee-details/model"

type PostCache interface {
	Set(key string, value model.Employee)
	Get(key string) *model.Employee
}
