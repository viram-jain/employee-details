package main

import (
	"encoding/json"
	"os"
	"rmqandredis/config"
	"rmqandredis/model"
	"rmqandredis/redis"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var (
	sugarLogger *zap.SugaredLogger
	postCache   redis.PostCache = redis.NewRedisCache(os.Getenv("REDIS_URL"), 0, 10)
)

func init() {
	config.Loadenv()
	sugarLogger = config.InitLogger()
	config.ConnectToMongo()
	config.ConnectToMysql()
}

func main() {
	ch := config.RabbitMQConnection()
	msgs, err := ch.Consume(
		os.Getenv("RMQ_QUEUE"),
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		sugarLogger.Errorf("Failed to consume message %s", err.Error())
		return
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var employee model.Employee
			err = json.Unmarshal(d.Body, &employee)
			if err != nil {
				sugarLogger.Errorf("Failed to unmarshal %s", err)
			}
			var post *model.Employee = postCache.Get(employee.EmpID)
			var res = config.GetDetailsFromMysql(employee.EmpID)
			if post == nil {
				postCache.Set(employee.EmpID, res)
			}
			config.AddToMongo(employee)
			err = d.Ack(true)
			if err != nil {
				sugarLogger.Error(err)
			}
		}
	}()
	<-forever
}
