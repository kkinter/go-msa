package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      //type
		true,         // durable?
		false,        //autuo-deleted?
		false,        //internal ?
		false,        //no-wait?
		nil,          //args ?

	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name?
		false, // durable?
		false, // delete whne unused?
		true,  // exclusive?
		false, // no-wait ?
		nil,   // args?
	)
}
