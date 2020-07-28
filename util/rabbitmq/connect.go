package mq

import (
	"github.com/iamwm/connector/exception"
	"github.com/iamwm/connector/store"
	"github.com/iamwm/connector/util/conf"
	"github.com/iamwm/connector/util/rabbitmq/wrapper"
	"log"
)

func InitRabbitMqHandler() {
	log.Println("START INIT MQ HANDLER")
	var globalStore = store.STORE
	rabbitmqConfig, err := conf.InitRabbitmqConfig("")
	if err != nil {
		log.Println("Mq config is invalid")
		log.Println(exception.ErrRabbitmqConfig)
		return
	}
	globalStore.RabbitmqConf = rabbitmqConfig
	conn, err := wrapper.Dial("amqp://" + rabbitmqConfig.User + ":" + rabbitmqConfig.Password + "@" + rabbitmqConfig.Host + ":" + rabbitmqConfig.Port + "/")
	if err != nil {
		log.Println("Mq connect failed")
		return
	}

	globalStore.Rabbitmq = conn
	log.Println("MQ HANDLER INIT COMPLETE")
}
