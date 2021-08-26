package publisher

import (
	"fmt"
	"os"

	"rmqandredis/config"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var sugarLogger *zap.SugaredLogger

func init() {
	config.Loadenv()
	sugarLogger = config.InitLogger()
}

//Publisher method to publish message to queue
func Publisher(data []byte) {
	ch := config.RabbitMQConnection()
	_, err := ch.QueueDeclare(
		os.Getenv("RMQ_QUEUE"),
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		sugarLogger.Errorf("Failed to declare queue %s", err.Error())
		return
	}

	err = ch.Publish(
		"",
		os.Getenv("RMQ_QUEUE"),
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		},
	)

	if err != nil {
		sugarLogger.Errorf("Failed to publish message %s", err.Error())
		return
	}

	fmt.Println("Successfully Published message to Queue")
}
