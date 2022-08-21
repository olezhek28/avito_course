package rabbit

import (
	"io"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/config"
	"github.com/streadway/amqp"
)

const contentType = "text/plain"

// Client ...
type Client interface {
	Publish(msg []byte) error
	Close() error
}

type client struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue

	queueName string

	closeFuncs []io.Closer
}

// NewClient ...
func NewClient(config *config.Rabbit) (Client, error) {
	var closeFuncs []io.Closer

	conn, err := amqp.Dial(config.DSN)
	if err != nil {
		return nil, err
	}
	closeFuncs = append(closeFuncs, conn)

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	closeFuncs = append(closeFuncs, ch)

	q, err := ch.QueueDeclare(
		config.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &client{
		connection: conn,
		channel:    ch,
		queue:      &q,
		queueName:  config.QueueName,
	}, nil
}

func (c *client) Publish(msg []byte) error {
	err := c.channel.Publish(
		"",
		c.queueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  contentType,
			Body:         msg,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) Close() error {
	for _, closeFunc := range c.closeFuncs {
		closeFunc.Close()
	}

	return nil
}
