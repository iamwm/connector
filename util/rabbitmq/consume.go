package mq

import (
	"github.com/iamwm/connector/store"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Println(err, msg)
	}
}

func StartConsume(mqProcess func(messages <-chan amqp.Delivery)) {
	mqConnection := store.STORE.Rabbitmq
	config := store.STORE.RabbitmqConf
	ch, err := mqConnection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.Qos(
		500,
		0,
		false,
	)
	failOnError(err, "Failed to set Qos")

	err = ch.ExchangeDeclare(
		config.Exchange, // name
		"topic",         // type
		false,           // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		config.Queue, // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for _, routingKey := range config.RoutingKeys {
		err = ch.QueueBind(
			q.Name,          // queue name
			routingKey,      // routing key
			config.Exchange, // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}

	msgs, err := ch.Consume(
		q.Name,      // queue
		"connector", // consumer
		false,       // auto ack
		false,       // exclusive
		false,       // no local
		false,       // no wait
		nil,         // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		mqProcess(msgs)
	}()

	log.Println("Waiting for logs. To exit press CTRL+C")
	<-forever
}
