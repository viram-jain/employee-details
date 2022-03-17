package config

import (
	"os"

	"github.com/streadway/amqp"
)

func RabbitMQConnection() *amqp.Channel {
	conn, err := amqp.Dial("amqp://" + os.Getenv("RMQ_USERNAME") + ":" + os.Getenv("RMQ_PASSWORD") + "@" + os.Getenv("RMQ_HOST") + os.Getenv("RMQ_PORT"))
	if err != nil {
		sugarLogger.Errorf("Failed to dial to RMQ %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		sugarLogger.Errorf("Failed to create channel %s", err.Error())
	}
	return ch
}
