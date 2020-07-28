package mq

import (
	"github.com/iamwm/connector/exception"
	"github.com/iamwm/connector/store"
	"github.com/iamwm/connector/util/conf"
	"github.com/iamwm/connector/util/rabbitmq/wrapper"
	"log"
)

func InitRabbitMqHandler(confPath string) bool {
	log.Println("START INIT MQ HANDLER")
	var globalStore = store.STORE
	rabbitmqConfig, err := conf.InitRabbitmqConfig(confPath)
	if err != nil {
		log.Println("Mq config is invalid")
		log.Println(exception.ErrRabbitmqConfig)
		return false
	}
	globalStore.RabbitmqConf = rabbitmqConfig
	conn, err := wrapper.Dial("amqp://" + rabbitmqConfig.User + ":" + rabbitmqConfig.Password + "@" + rabbitmqConfig.Host + ":" + rabbitmqConfig.Port + "/")
	if err != nil {
		log.Println("Mq connect failed")
		return false
	}

	globalStore.Rabbitmq = conn
	log.Println("MQ HANDLER INIT COMPLETE")
	return true
}
