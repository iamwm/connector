package exception

import "errors"

var ErrMongoConfig = errors.New("mongo config error")
var ErrRabbitmqConfig = errors.New("rabbitmq config error")
var ErrRedisConfig = errors.New("redis config error")
