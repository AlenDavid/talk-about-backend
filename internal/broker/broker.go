package broker

import (
	"errors"

	"github.com/rabbitmq/amqp091-go"
)

type Broker struct {
	connection *amqp091.Connection
	Channel    *amqp091.Channel
}

func (b Broker) Close() error {
	return errors.Join(
		b.connection.Close(),
		b.Channel.Close(),
	)
}

func Run(url string) (*Broker, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, errors.Join(err, conn.Close())
	}

	return &Broker{conn, channel}, nil
}
