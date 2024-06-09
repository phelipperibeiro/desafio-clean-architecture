package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/rabbitmq/amqp091-go"

	"github.com/phelipperibeiro/desafio-clean-architecture/pkg/events"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp091.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp091.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (orderCreatedHandler *OrderCreatedHandler) Handle(event events.EventInterface, waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()

	fmt.Printf("Order created: %#v", event.GetPayload())

	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp091.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	err := orderCreatedHandler.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)

	if err != nil {
		return
	}
}
