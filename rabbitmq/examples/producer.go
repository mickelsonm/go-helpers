package main

import (
	"encoding/json"
	"log"

	"github.com/mickelsonm/go-helpers/rabbitmq"
)

func main() {
	exchange := rabbitmq.Exchange{
		Name:       "exchange",
		RoutingKey: "hacker",
	}

	prod, err := rabbitmq.NewProducer(exchange, nil)
	if err != nil {
		log.Println(err)
		return
	}

	mess := "this is a really good test"
	js, err := json.Marshal(mess)
	if err != nil {
		return
	}

	//message is just []byte...doesn't have to be json
	if err = prod.SendMessage(js); err != nil {
		log.Println(err)
		return
	}
}
