package connector

import (
	"github.com/iamwm/connector/util/mongo"
	mq "github.com/iamwm/connector/util/rabbitmq"
)

func MongoConnect() {
	mongo.InitMongoHandler()
}

func RabbitmqConnect() {
	mq.InitRabbitMqHandler()
}
