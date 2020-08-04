package store

import (
	"github.com/go-redis/redis/v8"
	"github.com/iamwm/connector/util/conf"
	"github.com/iamwm/connector/util/rabbitmq/wrapper"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	Mongo        *mongo.Client
	MongoConf    *conf.MongoConfig
	Rabbitmq     *wrapper.Connection
	RabbitmqConf *conf.MqConfig
	RedisConf    *conf.RedisConfig
	Redis        *redis.Client
}

var STORE = &Store{}
