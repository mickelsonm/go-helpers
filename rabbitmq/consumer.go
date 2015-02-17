package rabbitmq

import (
	"errors"
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	Name     string
	Exchange Exchange
	DoneChan chan error
}

func (c *Consumer) Close() error {
	var err error
	if c.channel != nil {
		if err = c.channel.Cancel(c.Name, true); err != nil {
			return err
		}
	}
	if c.conn != nil {
		if err = c.conn.Close(); err != nil {
			return err
		}
	}

	return nil
}

func NewConsumer(consumerName string, queueName string, exchange Exchange, config *Config) (consumer *Consumer, err error) {
	//validate our parameters
	if consumerName == "" {
		err = errors.New("Must give the consumer a name")
		return
	}
	if queueName == "" {
		err = errors.New("Must specify the queue name")
		return
	}
	if exchange.Type == "" {
		exchange.Type = "direct"
	}
	if err = exchange.Validate(); err != nil {
		return
	}

	//setup our connection etc
	if config == nil {
		config = NewConfig()
	}
	conn, err := amqp.Dial(config.GetConnectionString())
	if err != nil {
		return
	}

	//setup the channel
	var channel *amqp.Channel
	if channel, err = conn.Channel(); err != nil {
		return
	}

	//setup the exchange
	if err = channel.ExchangeDeclare(
		exchange.Name, //exchange name
		exchange.Type, //exchange type
		true,          //durable
		false,         //remove when complete
		false,         //internal
		false,         //noWait
		nil,           //arguments
	); err != nil {
		return
	}

	//setup the queue
	var queue amqp.Queue
	if queue, err = channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	); err != nil {
		return
	}
	if err = channel.QueueBind(
		queue.Name,          //queue name
		exchange.RoutingKey, //routing key ("binding key")
		exchange.Name,       //exchange (source)
		false,               //noWait
		nil,                 //arguments
	); err != nil {
		return
	}

	//setup the deliverables
	var deliveries <-chan amqp.Delivery

	if deliveries, err = channel.Consume(
		queue.Name,   //queue name
		consumerName, //consumer name
		false,        //auto acknowledge
		false,        //exclusive
		false,        //not local
		false,        //no wait
		nil,          //arguments
	); err != nil {
		return
	}

	consumer = new(Consumer)
	consumer.Name = consumerName
	consumer.conn = conn
	consumer.channel = channel
	consumer.Exchange = exchange

	go deliveryHandler(deliveries, consumer.DoneChan)

	return
}

func deliveryHandler(deliveries <-chan amqp.Delivery, doneChan chan error) {
	for d := range deliveries {
		log.Printf("Got message: %s\n", string(d.Body))
		d.Ack(false)
	}
	doneChan <- nil
}
