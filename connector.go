package connector

import (
	"github.com/iamwm/connector/util/mongo"
	mq "github.com/iamwm/connector/util/rabbitmq"
)

func MongoConnect(confPath string) bool {
	return mongo.InitMongoHandler(confPath)
}

func RabbitmqConnect(confPath string) bool {
	return mq.InitRabbitMqHandler(confPath)
}
