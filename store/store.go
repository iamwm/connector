package store

import (
	"github.com/iamwm/connector/util/conf"
	"github.com/iamwm/connector/util/rabbitmq/wrapper"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	Mongo        *mongo.Client
	MongoConf    *conf.MongoConfig
	Rabbitmq     *wrapper.Connection
	RabbitmqConf *conf.MqConfig
}

var STORE = &Store{}
