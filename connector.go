package connector

import (
	"github.com/iamwm/connector/util/mongo"
	mq "github.com/iamwm/connector/util/rabbitmq"
	"github.com/iamwm/connector/util/redis"
)

func MongoConnect(confPath string) bool {
	return mongo.InitMongoHandler(confPath)
}

func RabbitmqConnect(confPath string) bool {
	return mq.InitRabbitMqHandler(confPath)
}

func RedisConnect(configPath string) bool {
	return redis.InitRedisHandler(configPath)
}
