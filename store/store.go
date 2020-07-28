package store

import (
	"github.com/iamwm/connector/util/rabbitmq/wrapper"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	Mongo    *mongo.Client
	Rabbitmq *wrapper.Connection
}

var STORE = &Store{}
