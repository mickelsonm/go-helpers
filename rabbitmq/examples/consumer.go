package main

import (
	"log"

	"github.com/mickelsonm/go-helpers/rabbitmq"
)

func main() {
	exchange := rabbitmq.Exchange{
		Name:       "exchange",
		RoutingKey: "hacker",
	}
	consumer, err := rabbitmq.NewConsumer("simple-consumer", "test-queue", exchange, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		select {
		case <-consumer.DoneChan:

		}
	}

	consumer.Close()
}
