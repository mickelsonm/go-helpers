package rabbitmq

import (
	"errors"
	"os"

	"github.com/streadway/amqp"
)

type Producer struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	Exchange Exchange
}

func NewProducer(exchange Exchange) (producer *Producer, err error) {
	if exchange.Type == "" {
		exchange.Type = "direct"
	}
	if err = exchange.Validate(); err != nil {
		return
	}

	var conn *amqp.Connection
	if os.Getenv("AMQP_HOST") != "" {
		conn, err = amqp.Dial(os.Getenv("AMQP_HOST"))
	} else {
		conn, err = amqp.Dial("amqp://localhost:5672")
	}
	if err != nil {
		return
	}

	producer = new(Producer)
	producer.conn = conn
	producer.Exchange = exchange
	producer.channel, err = conn.Channel()

	return
}

func (p *Producer) SendMessage(mess []byte) error {
	if p.channel == nil {
		return errors.New("Invalid channel")
	}
	if err := p.Exchange.Validate(); err != nil {
		return err
	}
	if len(mess) < 1 {
		return errors.New("Message cannot be empty")
	}
	return p.channel.Publish(
		p.Exchange.Name,
		p.Exchange.RoutingKey,
		false, //mandatory?
		false, //immediate?
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "application/json",
			ContentEncoding: "UTF-8",
			Body:            mess,
			DeliveryMode:    amqp.Transient,
			Priority:        0,
		},
	)
}
